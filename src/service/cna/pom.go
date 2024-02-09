package cna

import (
	"bytes"
	"fr/hijokaidan/config"
	"fr/hijokaidan/utils"
	"os"
	"text/template"
)

func createPom(pom *Pom, fileTree *config.FileTree) {
    var tplBytes bytes.Buffer
    fileContent, err := os.ReadFile(config.RESOURCE_FOLDER + "pom.xml");
    utils.HandleTechnicalError(err, config.ERR_TEMPLATE_FILE_READ)
    tmpl, err := template.New("test").Parse(string(fileContent))
    if err != nil { panic(err) }
    err = tmpl.Execute(&tplBytes, pom)
    writePom(&tplBytes, fileTree)
}

func writePom(content *bytes.Buffer, fileTree *config.FileTree) {
		f, err := os.OpenFile(fileTree.Root + "pom.xml", os.O_TRUNC|os.O_CREATE|os.O_WRONLY, 0644)
		defer f.Close()
		if err != nil {
			panic(err)
		}
		f.Write(content.Bytes())
}

