package thymeleaf

import (
	"context"
	"fmt"
	"fr/nzc/config"

	sitter "github.com/smacker/go-tree-sitter"
	"github.com/smacker/go-tree-sitter/html"
)

func TestParser() {
    parser := sitter.NewParser()
    lang := html.GetLanguage()
    parser.SetLanguage(lang)
    sourceCodeAsString := `
<div th:text="${cours.getNom()}">Placeholder</div>
    `
    var newCode string
    sourceCode := []byte(sourceCodeAsString)
	n, _ := sitter.ParseCtx(context.Background(), sourceCode, lang)
	q, _ := sitter.NewQuery([]byte(Q_TH_ATTRIBUTE), lang)
	qc := sitter.NewQueryCursor()
	qc.Exec(q, n)
    for {
		m, ok := qc.NextMatch()
		if !ok {
			break
		}
        m = qc.FilterPredicates(m, sourceCode)
        for _, c := range m.Captures {
            targetedNode := c.Node.Parent().Child(2).Child(1)
            idxStart := targetedNode.StartByte()
            idxEnd := targetedNode.EndByte()
            fmt.Println(string(sourceCode[idxStart:idxEnd]))
            newCode = sourceCodeAsString[:idxStart] + "test" + sourceCodeAsString[idxEnd:]
        }
    }
    fmt.Println(newCode)
}

func RenameProjectFiles(oldname, newName string, fileTree *config.FileTree) {
    
}

func parsePage() {

}

func renameRoute() {

}

func replaceInsertsAndReplaces() {

}

