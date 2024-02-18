package thymeleaf

import (
	"context"
	"fmt"
	"fr/nzc/config"
	"strings"

	sitter "github.com/smacker/go-tree-sitter"
	"github.com/smacker/go-tree-sitter/html"
)

var htmlParser *sitter.Parser
var lang *sitter.Language

func init(){
    lang = html.GetLanguage()
    htmlParser = sitter.NewParser()
    htmlParser.SetLanguage(lang)
}

func mvReplacesAndInserts(sourceCodeAsString, oldName, newName string) string {
    newCode := sourceCodeAsString
    sourceCode := []byte(sourceCodeAsString)
	n, _ := sitter.ParseCtx(context.Background(), sourceCode, lang)
	q, _ := sitter.NewQuery([]byte(Q_TH_REPLACE_INSERT), lang)
	qc := sitter.NewQueryCursor()
	qc.Exec(q, n)
    for {
		m, ok := qc.NextMatch()
		if !ok {
			break
		}
        m = qc.FilterPredicates(m, sourceCode)
        fmt.Println(m.Captures)
        for _, c := range m.Captures {
            fmt.Println(c.Node.Content(sourceCode))
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
                newName = strings.ReplaceAll(contentAsString, oldName, newName)
                idxStart := attributeContent.StartByte()
                idxEnd := attributeContent.EndByte()
                newCode = sourceCodeAsString[:idxStart] + newName + sourceCodeAsString[idxEnd:]
            }
        }
    }
    return newCode
}

func RenameProjectFiles(oldname, newName string, fileTree *config.FileTree) {
    
}

func parsePage() {

}

func renameRoute() {

}

func replaceInsertsAndReplaces() {

}

