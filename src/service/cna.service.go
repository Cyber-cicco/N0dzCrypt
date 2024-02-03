package service

import (
	"bufio"
	"fmt"
	"fr/hijokaidan/config"
	"fr/hijokaidan/utils"
	"os"
	"strconv"
	"strings"
	"text/template"
)

type Pom struct {
    SpringVersion string
    GroupId string
    ArtifactId string
    ProjectName string
    Description string
    JavaVersion int
    JwtVersion string
    AdditionalProperties []string
    MainClass string
    Profiles []string
}

func CreateNodzCryptApp() {
    scanner := bufio.NewScanner(os.Stdin)
    projectProps := config.ProjectProps{}
    pom := Pom{}

    pom.SpringVersion = "3.1.5"
    askOrganisationId(&pom, scanner)
    askProjectName(&pom, scanner)
    projectProps.MainPackage = pom.GroupId + "." + pom.ArtifactId
    askProjectDescription(&pom, scanner)
    askJavaVersion(&pom, scanner)
    pom.JwtVersion = "0.11.5"

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
