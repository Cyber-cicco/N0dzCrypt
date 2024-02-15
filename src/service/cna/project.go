package cna

import (
	"fr/nzc/config"
	"fr/nzc/utils"
	"os"
)

func createDirectories(fileTree *config.FileTree) {

    for _, dirname := range []string{
        fileTree.GetPageBackDir(),
        fileTree.GetPageBackDir() + fileTree.JavaBack.PagesDir.Irrigator,
        fileTree.GetEntityDir(),
        fileTree.GetEntityDir() + fileTree.JavaBack.Entities.Enum,
        fileTree.GetUtilsDir() + fileTree.JavaBack.Utils.HX,
        fileTree.GetJavaDir() + fileTree.JavaBack.Repository,
        fileTree.GetJavaDir() + fileTree.JavaBack.Service,
        fileTree.GetSecurityDir(),
        fileTree.GetSecurityDir() + fileTree.JavaBack.Security.Service,
        fileTree.GetSecurityDir() + fileTree.JavaBack.Security.Config,
        fileTree.GetJavaDir() + fileTree.JavaBack.Validators,
        fileTree.GetJsonDir(),
        fileTree.GetJsonDir() + fileTree.JavaBack.Json.Mapper,
        fileTree.GetJsonDir() + fileTree.JavaBack.Json.Dto,
        fileTree.GetPageFrontDir(),
        fileTree.GetTemplateDir() + fileTree.Resources.Templates.Components,
        fileTree.GetTemplateDir() + fileTree.Resources.Templates.Layouts,
        fileTree.GetStaticDir() + fileTree.Resources.Static.Style,
        fileTree.GetTemplateDir() + fileTree.Resources.Templates.StyleTemplates,
        fileTree.GetStaticDir() + fileTree.Resources.Static.JS,
        fileTree.GetTemplateDir() + fileTree.Resources.Templates.JS,
        fileTree.GetTemplateDir() + fileTree.Resources.Templates.SVG,
        fileTree.GetTestDir(),
        fileTree.GetTestDir() + fileTree.JavaTest.TestMain,
        fileTree.GetTestDir() + fileTree.JavaTest.TestRessources,
        fileTree.GetPageFrontDir() + "home/",
        fileTree.GetPageFrontDir() + "about/",
    } {
        err := os.MkdirAll(dirname, os.ModePerm)
        utils.HandleTechnicalError(err, config.ERR_DIR_CREATION)
    }
}
