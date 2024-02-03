package service

import (
	"fr/hijokaidan/config"
	"fr/hijokaidan/utils"
	"os"
	"text/template"
)

type Pom struct {
    SpringVersion string
    GroupId string
    ArtifactId string
    ProjectName string
    JavaVersion string
    JwtVersion string
    AdditionalProperties []string
    MainClass string
    Profiles []string
}

func CreateNodzCryptApp() {
    pom := Pom{}
    pom.SpringVersion = "3.1.5"
    fileContent, err := os.ReadFile("../resources/pom.xml");
    utils.HandleTechnicalError(err, config.ERR_TEMPLATE_FILE_READ)
    tmpl, err := template.New("test").Parse(string(fileContent))
    if err != nil { panic(err) }
    err = tmpl.Execute(os.Stdout, pom)
    if err != nil { panic(err) }
}
