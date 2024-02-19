package java

import (
	"context"
	"fmt"
	"fr/nzc/config"
	"fr/nzc/utils"
	"os"
	"strings"

	sitter "github.com/smacker/go-tree-sitter"
	"github.com/smacker/go-tree-sitter/java"
)

var javaParser *sitter.Parser
var javaLang *sitter.Language

func init(){
    javaLang = java.GetLanguage()
    javaParser = sitter.NewParser()
    javaParser.SetLanguage(javaLang)
}

func changeRouteInJavaFile(sourceCodeAsString, oldName, newName string) string {
    newCode := []string{}
    sourceCode := []byte(sourceCodeAsString)
	n, _ := sitter.ParseCtx(context.Background(), sourceCode, javaLang)
	q, _ := sitter.NewQuery([]byte(Q_JAVA_FINAL), javaLang)
	qc := sitter.NewQueryCursor()
	qc.Exec(q, n)
    lastIndex := 0
    for {
		m, ok := qc.NextMatch()
		if !ok {
			break
		}
        m = qc.FilterPredicates(m, sourceCode)
        for _, c := range m.Captures {
            stringLitteral := c.Node.ChildByFieldName("value").Child(1)
            contentAsString := stringLitteral.Content(sourceCode)
            fmt.Println(contentAsString)
            if(strings.Contains(contentAsString, oldName)) {
                newAdress := strings.ReplaceAll(contentAsString, oldName, newName)
                idxStart := stringLitteral.StartByte()
                idxEnd := stringLitteral.EndByte()
                newCode = append(newCode, sourceCodeAsString[lastIndex:idxStart] + newAdress)
                lastIndex = int(idxEnd)
            }
        }
    }
    newCode = append(newCode, sourceCodeAsString[lastIndex:])
    fmt.Println(strings.Join(newCode, ""))
    return strings.Join(newCode, "")
}

func RenameRoute(oldname, newName string, fileTree *config.FileTree) {
    pathOfRoutes := fileTree.ProjectAbsolutePath + fileTree.GetPageBackDir() + "Routes.java"
    file, err := os.ReadFile(pathOfRoutes)
    utils.HandleTechnicalError(err, config.ERR_TEMPLATE_FILE_READ)
    changeRouteInJavaFile(oldname, newName, string(file))
    //daos.WriteToFile([]byte(newCode), pathOfRoutes)
}

