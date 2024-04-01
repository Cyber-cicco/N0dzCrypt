package java

import (
	"context"
	"fr/nzc/config"
	"fr/nzc/daos"
	"fr/nzc/utils"
	"os"
	"strings"

	sitter "github.com/Cyber-cicco/go-tree-sitter"
	"github.com/Cyber-cicco/go-tree-sitter/java"
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
    sourceCode := []byte(sourceCodeAsString)
	n, _ := sitter.ParseCtx(context.Background(), sourceCode, javaLang)
	q, _ := sitter.NewQuery([]byte(Q_JAVA_STRING), javaLang)
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
                newVarName := utils.GetNewAdressVarFromName(newName)
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

func changeNameInJavaFile(oldName, newName, sourceCodeAsString string) string {
    newCode := []string{}
    sourceCode := []byte(sourceCodeAsString)
	n, _ := sitter.ParseCtx(context.Background(), sourceCode, javaLang)
	q, _ := sitter.NewQuery([]byte(Q_JAVA_ROUTE), javaLang)
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
            name := c.Node
            contentAsString := name.Content(sourceCode)
            if utils.GetNewAdressVarFromName(oldName) == contentAsString {
                newVarName := utils.GetNewAdressVarFromName(newName)
                idxStart := name.StartByte()
                idxEnd := name.EndByte()
                newCode = append(newCode, sourceCodeAsString[lastIndex:idxStart] + newVarName)
                lastIndex = int(idxEnd)
            }
        }
    }
    newCode = append(newCode, sourceCodeAsString[lastIndex:])
    return strings.Join(newCode, "")
}

func changeControllerAdress(oldname, newname, sourceCodeAsString string) string {
    oldname = "/" + utils.StripPage(oldname)
    newname = "/" + utils.StripPage(newname)
    newCode := []string{}
    sourceCode := []byte(changeNameInJavaFile(oldname, newname, sourceCodeAsString))
    sourceCodeAsString = string(sourceCode)
	n, _ := sitter.ParseCtx(context.Background(), sourceCode, javaLang)
	q, _ := sitter.NewQuery([]byte(Q_JAVA_HTTP_ANNOTATION), javaLang)
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
            annotation := c.Node.Parent()
            argument := annotation.ChildByFieldName("arguments")
            annotation_name := annotation.Child(1).Content(sourceCode)
            for i := 0; i <  int(argument.ChildCount()); i++ {
                currArg := argument.Child(i)
                argType := currArg.Type()
                if argType == "element_value_pair" {
                    key := currArg.ChildByFieldName("key").Content(sourceCode)
                    litteral := currArg.ChildByFieldName("value").Child(1)
                    if key == "value" {
                        if litteral.Content(sourceCode) ==  oldname {
                            idxStart := litteral.StartByte()
                            idxEnd := litteral.EndByte()
                            newCode = append(newCode, sourceCodeAsString[lastIndex:idxStart] + newname)
                            lastIndex = int(idxEnd)
                        }
                    }
                } else if argType == "string_literal" {
                    litteral := currArg.Child(1)
                    if litteral.Content(sourceCode) == oldname {
                        idxStart := litteral.StartByte()
                        idxEnd := litteral.EndByte()
                        newCode = append(newCode, sourceCodeAsString[lastIndex:idxStart] + newname)
                        lastIndex = int(idxEnd)
                    }
                }
            }
            if annotation_name == "RequestMapping" {
            }
        }
    }
    newCode = append(newCode, sourceCodeAsString[lastIndex:])
    return strings.Join(newCode, "")
}

func RenameRoute(oldname, newName string, fileTree *config.FileTree) {
    pathOfRoutes := fileTree.ProjectAbsolutePath + fileTree.GetPageBackDir() + "Routes.java"
    file, err := os.ReadFile(pathOfRoutes)
    utils.HandleTechnicalError(err, config.ERR_TEMPLATE_FILE_READ)
    newCode := changeRouteInRoutesFile(oldname, newName, string(file))
    daos.WriteToFile([]byte(newCode), pathOfRoutes)
}

func ParseJavaPageFiles(oldname, newname string, fileTree *config.FileTree) {
    pathOfPages := fileTree.ProjectAbsolutePath + fileTree.GetPageBackDir()
    daos.ParseFolders(".java", pathOfPages, func(content, filePath string){
        newCode := changeNameInJavaFile(oldname, newname, content)
        daos.WriteToFile([]byte(newCode), filePath)
    })
}

func ChangeBackendPageReferences(oldname, newname string, fileTree *config.FileTree) {
    pathOfPages := fileTree.ProjectAbsolutePath + fileTree.GetPageBackDir()
    daos.ParseFolders(".java", pathOfPages, func(content, filePath string){
        newCode := changeControllerAdress(oldname, newname, content)
        daos.WriteToFile([]byte(newCode), filePath)
    })
}
