package cna

import (
	"fmt"
	"fr/hijokaidan/config"
	"fr/hijokaidan/utils"
	"os"
)

func createDirectories(fileTree *config.FileTree) {

    for _, dirname := range []string{
        fileTree.PagesBack,
        fileTree.Irrigator,
        fileTree.Entities,
        fileTree.Repository,
        fileTree.Service,
        fileTree.Utils,
        fileTree.Validators,
        fileTree.Json,
        fileTree.Components,
        fileTree.PagesFront,
        fileTree.StyleStatic,
        fileTree.StyleTemplates,
        fileTree.Test,
        fileTree.TestResources,
        fileTree.Layouts,
    } {
        err := os.MkdirAll(dirname, os.ModePerm)
        fmt.Printf("dirname: %v\n", dirname)
        utils.HandleTechnicalError(err, config.ERR_DIR_CREATION)
    }
}
