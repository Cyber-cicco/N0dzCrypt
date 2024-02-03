package service

import (
	"bufio"
	"fmt"
	"fr/hijokaidan/config"
	"fr/hijokaidan/utils"
	"os"
	"slices"
	"strconv"
	"strings"
	"text/template"

	"github.com/IMQS/options"
)

func CreateNodzCryptApp() {

    scanner := bufio.NewScanner(os.Stdin)
    projectProps := config.ProjectProps{}
    profile := Profile{}
    pom := Pom{}
    // sqlConnection := SQLConnection{}

    pom.SpringVersion = "3.2.2"
    askOrganisationId(&pom, scanner)
    askProjectName(&pom, scanner)
    projectProps.MainPackage = utils.GetPackageName(pom.ArtifactId, pom.GroupId)
    askProjectDescription(&pom, scanner)
    askJavaVersion(&pom, scanner)
    pom.JwtVersion = "0.11.5"
    askConnectionInfos(scanner, &profile)

    dependencies := askDependencies(scanner)



    createApplicationProperties(dependencies)

    if wantsProfile(scanner) {
    }

    fileContent, err := os.ReadFile("../resources/pom.xml");
    utils.HandleTechnicalError(err, config.ERR_TEMPLATE_FILE_READ)
    tmpl, err := template.New("test").Parse(string(fileContent))
    if err != nil { panic(err) }
    err = tmpl.Execute(os.Stdout, pom)
    if err != nil { panic(err) }
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

func wantsProfile(scanner *bufio.Scanner) bool {
    answer := ""
    possibleAnswers := []string{
        "y", "n","Y", "N", "yes", "no", "Yes", "No",
    }
    fmt.Print("Do you want to set up a new profile ? (y/n) : ")
    for slices.Contains(possibleAnswers, answer) {
        if scanner.Scan() {
            answer = scanner.Text()
        }
    }
    return answer == "y" || answer == "Y" || answer == "yes" || answer == "Yes"
}

func askDependencies(scanner *bufio.Scanner) *SQLConnection {
    con := options.NewConsole()
	defer con.Close()
    sqlConnection := SQLConnection{}

	boxes := []string {
		"mysql",
		"mariadb",
		"postgresql",
		"h2",
		"oracle",
	}
    con.Radio("Chose your database", "", -1, boxes)
    return &sqlConnection;
    
}

func askConnectionInfos(scanner *bufio.Scanner, profile *Profile, ) {
    fmt.Println("Enter the name of your database : ")
    if scanner.Scan() {
    }
}

func createApplicationProperties(dependencies *config.Dependencies) {

}
