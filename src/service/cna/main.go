package cna

import (
	"bufio"
	"bytes"
	"fmt"
	"fr/nzc/config"
	"fr/nzc/utils"
	"os"
	"strconv"
	"strings"
	"text/template"

	"github.com/IMQS/options"
	"github.com/google/uuid"
	"golang.org/x/crypto/ssh/terminal"
)

func CreateNodzCryptApp() {

    scanner := bufio.NewScanner(os.Stdin)
    projectProps := config.ProjectProps{}
    pom := Pom{}

    pom.SpringVersion = "3.2.2"
    askOrganisationId(&pom, scanner)
    askProjectName(&pom, scanner)
    projectProps.MainPackage = utils.GetPackageName(pom.ArtifactId, pom.GroupId)
    askProjectDescription(&pom, scanner)
    askJavaVersion(&pom, scanner)
    pom.JwtVersion = "0.11.5"
    sqlConnection := askDependencies(scanner)
    pom.DBDependencie = sqlConnection.MavenDependencie
    dbInfos := askConnectionInfos(scanner, sqlConnection)
    pom.MainClass = utils.GetApplicationName(pom.ArtifactId)

    pom.Profiles = getProfile(&pom, dbInfos)
    fileTree := config.CreateConfig(projectProps.MainPackage)
    createDirectories(fileTree)
    createPom(&pom, fileTree)
    createBaseFiles(&pom, fileTree, projectProps.MainPackage)

    createApplicationProperties(fileTree)

}

func getProfile(pom *Pom, dbInfos *DBInfos, ) []string {
    profile := Profile{}
    profile.ActiveByDefault = true
    profile.DBInfos = *dbInfos
    profile.ProfileName = "dev"
    var tplBytes bytes.Buffer
    profileFileTemplate, err := os.ReadFile(config.RESOURCE_FOLDER + "profile.xml") 
    utils.HandleTechnicalError(err, config.ERR_TEMPLATE_FILE_READ)
    profileTmpl, err := template.New("test").Parse(string(profileFileTemplate))
    err = profileTmpl.Execute(&tplBytes, profile)
    return []string{string(tplBytes.String())}
}

func askOrganisationId(pom *Pom, scanner *bufio.Scanner) {
    valid := false
    var input string
    for !valid {
        valid = true
        input = ""
        fmt.Println("Enter the name of your organisation (eg. org.example) : ")
        if scanner.Scan() {
            input = scanner.Text()
            if len(input) > 255 || !utils.IsAcceptedCharacters(input) {
                fmt.Println("There was an error reading your input : you should have a name without any whitespaces whith only alphanumeric characters (along with '_' and '-') whose length shouldn't be beyond 255 characters")
                valid = false
            }
            continue
        }
        valid = false;
    }
    pom.GroupId = input
}

func askProjectName(pom *Pom, scanner *bufio.Scanner) {
    valid := false 
    var input string
    for !valid {
        valid = true
        fmt.Println("Enter the name of your project")
        if scanner.Scan() {
            input = scanner.Text()
            if len(input) > 255 || strings.Contains(input, " ") {
                fmt.Println("There was an error reading your input : you should have a name without any whitespaces whith only alphanumeric characters (along with '_' and '-') whose length shouldn't be beyond 255 characters")
                valid = false
            }
            continue
        }
        valid = false;
    }
    pom.ArtifactId = input
}

func askProjectDescription(pom *Pom, scanner *bufio.Scanner) {
    valid := false 
    var input string
    for !valid {
        valid = true
        fmt.Println("Enter the description of your project")
        if scanner.Scan() {
            input = scanner.Text()
            if len(input) > 510  {
                fmt.Println("There was an error reading your input : you should have a description whose length shouldn't be beyond 510 characters")
                valid = false
            }
            continue
        }
        valid = false;
    }
    pom.Description = input
}

func askJavaVersion(pom *Pom, scanner *bufio.Scanner) {
    valid := false 
    var input int
    var strInput string
    for !valid {
        valid = true
        fmt.Println("Enter the java version for your project (between 8 and 21)")
        if scanner.Scan() {
            strInput = scanner.Text()
            i, err := strconv.Atoi(strInput) 
            if err != nil {
                fmt.Println("You need to give a version between 8 and 21")
                valid = false
                continue
            }
            if (i < 8 || i > 21) {
                fmt.Println("You need to give a version between 8 and 21")
                valid = false
                continue
            }
            input = i
            continue
        }
        valid = false;
    }
    pom.JavaVersion = input
}

func askDependencies(scanner *bufio.Scanner) *SQLConnection {
    con := options.NewConsole()
	defer con.Close()

	boxes := []string {
		"mysql",
		"mariadb",
		"postgresql",
	}
    infos := con.Radio("Chose your database", "", -1, boxes)

    switch infos {
    case 0 : {
        return &SQLConnection{
            Url: "jdbc:mysql://localhost:3306/",
            Driver: "com.mysql.cj.jdbc.Driver",
            MavenDependencie: `

        <!-- https://mvnrepository.com/artifact/mysql/mysql-connector-java -->
        <dependency>
            <groupId>mysql</groupId>
            <artifactId>mysql-connector-java</artifactId>
            <version>8.0.33</version>
        </dependency>

            `,
        }
    }
    case 1 : {
        return &SQLConnection{
            Url: "jdbc:mariadb://localhost:3306/",
            Driver: "org.mariadb.jdbc.Driver",
            MavenDependencie: `
        <!-- https://mvnrepository.com/artifact/org.mariadb.jdbc/mariadb-java-client -->
        <dependency>
            <groupId>org.mariadb.jdbc</groupId>
            <artifactId>mariadb-java-client</artifactId>
            <version>3.3.2</version>
        </dependency>

            `,
        }
    }
    default : {
        return &SQLConnection{
            Url: "jdbc:postgresql://localhost:5432/",
            Driver: "org.postgresql.Driver",
            MavenDependencie: `
        <!-- https://mvnrepository.com/artifact/org.postgresql/postgresql -->
        <dependency>
            <groupId>org.postgresql</groupId>
            <artifactId>postgresql</artifactId>
            <version>42.7.1</version>
        </dependency>

            `,
        }
    }
    }
}

func askConnectionInfos(scanner *bufio.Scanner, sqlConnection *SQLConnection) *DBInfos {
    dbInfos := DBInfos{}
    dbInfos.DBDriver = sqlConnection.Driver
    fmt.Println("Enter the name of your database : ")
    if scanner.Scan() {
        sqlConnection.Url += scanner.Text()
        dbInfos.DBUrl = sqlConnection.Url
    } 
    fmt.Println("Enter the username of the account that will access the database : ")
    if scanner.Scan() {
        dbInfos.DBUser = scanner.Text()
    }
    fmt.Println("Enter the password of the account that will access the database : ")
    password, err := terminal.ReadPassword(0)
    utils.HandleTechnicalError(err, "Can't read password")
    dbInfos.DBPassword = string(password)
    return &dbInfos
}

func getJwtSecret() string {
	newUUID := uuid.New()
	return  newUUID.String()
}

func createApplicationProperties(fileTree *config.FileTree){
    properties := ApplicationProperties{
        JWTSecret: utils.RandStringBytes(20),
    }
    var tplBytes bytes.Buffer
    fileContent, err := os.ReadFile(config.RESOURCE_FOLDER + "application.yml");
    utils.HandleTechnicalError(err, config.ERR_TEMPLATE_FILE_READ)
    tmpl, err := template.New("application.yml").Parse(string(fileContent))
    if err != nil { panic(err) }
    err = tmpl.Execute(&tplBytes, properties)
    f, err := os.OpenFile(fileTree.GetResourcesDir() + "application.yml", os.O_TRUNC|os.O_CREATE|os.O_WRONLY, 0644)
    defer f.Close()
    if err != nil {
        panic(err)
    }
    f.Write(tplBytes.Bytes())
}
