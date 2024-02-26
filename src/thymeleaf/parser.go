package thymeleaf

import (
	"context"
	"fr/nzc/config"
	"fr/nzc/daos"
	"fr/nzc/utils"
	"slices"
	"strings"

	sitter "github.com/smacker/go-tree-sitter"
	"github.com/smacker/go-tree-sitter/html"
)

var htmlParser *sitter.Parser
var htmlLang *sitter.Language

func init(){
    htmlLang = html.GetLanguage()
    htmlParser = sitter.NewParser()
    htmlParser.SetLanguage(htmlLang)
}

func mvReplacesAndInserts(sourceCodeAsString, oldName, newName string) string {
    newCode := []string{}
    sourceCode := []byte(sourceCodeAsString)
	n, _ := sitter.ParseCtx(context.Background(), sourceCode, htmlLang)
	q, _ := sitter.NewQuery([]byte(Q_TH_REPLACE_INSERT), htmlLang)
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
            parentNode := c.Node.Parent()
            if parentNode.ChildCount() < 3 {
                continue
            }
            attributeNode := parentNode.Child(2)
            if attributeNode.ChildCount() < 2 {
                continue
            }
            attributeContent := attributeNode.Child(1)
            contentAsString := attributeContent.Content(sourceCode)
            if(strings.Contains(contentAsString, oldName)) {
                newAdress := strings.ReplaceAll(contentAsString, oldName, newName)
                idxStart := attributeContent.StartByte()
                idxEnd := attributeContent.EndByte()
                newCode = append(newCode, sourceCodeAsString[lastIndex:idxStart] + newAdress)
                lastIndex = int(idxEnd)
            }
        }
    }
    newCode = append(newCode, sourceCodeAsString[lastIndex:])
    return strings.Join(newCode, "")
}

func mvReferencesOfOldName(oldname, newname, sourceCodeAsString string) string {
    oldname = utils.StripPage(oldname)
    newname = utils.StripPage(newname)
    newCode := []string{}
    sourceCode := []byte(sourceCodeAsString)
	n, _ := sitter.ParseCtx(context.Background(), sourceCode, htmlLang)
	q, _ := sitter.NewQuery([]byte(Q_ATTRIBUTE_HX), htmlLang)
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
            attribute := c.Node.Content(sourceCode)
            if slices.Contains[[]string, string](pageTags, attribute) {
                valueNode := c.Node.Parent().Child(2).Child(1)
                value := valueNode.Content(sourceCode)
                if strings.Contains(value, oldname) {
                    value = strings.ReplaceAll(value, oldname, newname)
                    idxStart := valueNode.StartByte()
                    idxEnd := valueNode.EndByte()
                    newCode = append(newCode, sourceCodeAsString[lastIndex:idxStart] + value)
                    lastIndex = int(idxEnd)
                }
            }
        }
    }
    newCode = append(newCode, sourceCodeAsString[lastIndex:])
    return strings.Join(newCode, "")
}

func RenameProjectFiles(oldname, newName string, fileTree *config.FileTree) {
    pathOfTemplates := fileTree.ProjectAbsolutePath + fileTree.GetTemplateDir()
    daos.ParseFolders(".html", pathOfTemplates, func(content, filePath string){
        newCode := mvReplacesAndInserts(content, oldname, newName)
        daos.WriteToFile([]byte(newCode), filePath)
    })
}

func ReplacePageReferences(oldname, newname string, fileTree *config.FileTree) {
    pathOfTemplates := fileTree.ProjectAbsolutePath + fileTree.GetTemplateDir()
    daos.ParseFolders(".html", pathOfTemplates, func(content, filePath string){
        newCode := mvReferencesOfOldName(oldname, newname, content)
        daos.WriteToFile([]byte(newCode), filePath)
    })
}



