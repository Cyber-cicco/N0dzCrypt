package cna

import (
	"bytes"
	"fr/hijokaidan/config"
	"fr/hijokaidan/utils"
	"os"
	"text/template"
)

func createPom(pom *Pom) string {
    var tplBytes bytes.Buffer
    fileContent, err := os.ReadFile("../resources/pom.xml");
    utils.HandleTechnicalError(err, config.ERR_TEMPLATE_FILE_READ)
    tmpl, err := template.New("test").Parse(string(fileContent))
    if err != nil { panic(err) }
    err = tmpl.Execute(&tplBytes, pom)
    return tplBytes.String()
}

func writePom(fileTree *config.FileTree) {
    
}

