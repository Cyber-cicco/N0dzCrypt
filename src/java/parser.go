package java

import (
	"context"
	"fmt"
	"fr/nzc/config"
	"fr/nzc/daos"
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

func changeRouteInRoutesFile(oldName, newName, sourceCodeAsString string) string {
    newCode := []string{}
    fmt.Println(sourceCodeAsString)
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
            name := c.Node.ChildByFieldName("name")
            contentAsString := stringLitteral.Content(sourceCode)
            if contentAsString == oldName {
                newVarName := utils.GetNewAdressVarFromOldName(newName)
                idxStart := name.StartByte()
                idxEnd := name.EndByte()
                newCode = append(newCode, sourceCodeAsString[lastIndex:idxStart] + newVarName)
                lastIndex = int(idxEnd)
                newAdress := strings.ReplaceAll(contentAsString, oldName, newName)
                idxStart = stringLitteral.StartByte()
                idxEnd = stringLitteral.EndByte()
                newCode = append(newCode, sourceCodeAsString[lastIndex:idxStart] + newAdress)
                lastIndex = int(idxEnd)
            }
        }
    }
    newCode = append(newCode, sourceCodeAsString[lastIndex:])
    return strings.Join(newCode, "")
}

func RenameRoute(oldname, newName string, fileTree *config.FileTree) {
    pathOfRoutes := fileTree.ProjectAbsolutePath + fileTree.GetPageBackDir() + "Routes.java"
    fmt.Println(pathOfRoutes)
    file, err := os.ReadFile(pathOfRoutes)
    fmt.Println(string(file))
    utils.HandleTechnicalError(err, config.ERR_TEMPLATE_FILE_READ)
    newCode := changeRouteInRoutesFile(oldname, newName, string(file))
    fmt.Println(newCode)
    daos.WriteToFile([]byte(newCode), pathOfRoutes)
}

