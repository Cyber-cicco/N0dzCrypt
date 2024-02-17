package cna

import (
	"fr/nzc/config"
	"fr/nzc/utils"
	"os"
)

func createDirectories(fileTree *config.FileTree) {

    for _, dirname := range []string{
        fileTree.GetPageBackDir(),
        fileTree.GetIrrigatorDir(),
        fileTree.GetEntityDir(),
        fileTree.GetEntityEnumDir(),
        fileTree.GetUtilsDir(), 
        fileTree.GetHXDir(),
        fileTree.GetRepositoryDir(),
        fileTree.GetServiceDir(),
        fileTree.GetSecurityDir(),
        fileTree.GetSecurityServiceDir(),
        fileTree.GetSecurityConfigDir(),
        fileTree.GetValidatorDir(),
        fileTree.GetJsonDir(),
        fileTree.GetJsonMapperDir(), 
        fileTree.GetJsonDtoDir(),
        fileTree.GetPageFrontDir(),
        fileTree.GetTemplateComponentsDir(),
        fileTree.GetTemplateLayoutsDir(),
        fileTree.GetTemplateJSDir(),
        fileTree.GetTemplateSVGDir(),
        fileTree.GetTemplateStyleDir(),
        fileTree.GetStaticStyleDir(),
        fileTree.GetStaticJsDir(),
        fileTree.GetStaticImgDir(),
        fileTree.GetTestDir(),
        fileTree.GetTestMainDir(),
        fileTree.GetTestRessourcesDir(),
        fileTree.GetPageFrontDir() + "home/",
        fileTree.GetPageFrontDir() + "about/",
    } {
        err := os.MkdirAll(dirname, os.ModePerm)
        utils.HandleTechnicalError(err, config.ERR_DIR_CREATION)
    }
}
