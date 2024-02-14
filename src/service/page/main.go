package page

import (
	"fr/nzc/config"
	"fr/nzc/daos"
	"fr/nzc/utils"
	"os"
	"path/filepath"
)

func CreatePage(args []string) {
    currDir, err := filepath.Abs(".")
    utils.HandleTechnicalError(err, config.ERR_CURR_DIR_OPEN)
    conf, err := daos.GetConfigFile(currDir)
    utils.HandleUsageError(err, config.ERR_COULDNT_FIND_CONFIG)
    pageHTML, err := os.ReadFile(config.RESOURCE_FOLDER + "n0dzcrypt.page.html")
    for _, arg := range args {
        javaInfos := &JavaClassInfos{
            BasePackage: conf.BasePackage,
            ClassName: utils.GetCamelCaseFromKebab(arg),
            PageName: arg,
            UpperClassName: utils.GetUpperSnakeCaseFromKebab(arg),
        }
        writePageFront(arg, pageHTML, conf)
        writePageBack(arg, conf, javaInfos)
        writeIrrigator(arg, conf, javaInfos)
        appendRoute(arg, conf, javaInfos)
    }
}

func writePageFront(arg string, pageHTML []byte, conf *config.FileTree) {
    dirPageFront := conf.CurrentDirectory + conf.PagesFront + arg + "/"
    err := os.MkdirAll(dirPageFront, os.ModePerm)
    utils.HandleTechnicalError(err, config.ERR_DIR_CREATION)
    daos.WriteToFile(pageHTML, dirPageFront + "base.html")
}

func writePageBack(arg string, conf *config.FileTree, javaInfos *JavaClassInfos) {
    dirPageBack := conf.CurrentDirectory + conf.PagesBack 
    pageContent := daos.GetTemplBytes[JavaClassInfos](arg, config.BASE_PAGE_BACK, *javaInfos)
    daos.WriteToFile(pageContent, dirPageBack + javaInfos.ClassName + "Controller.java")
}

func writeIrrigator(arg string, conf *config.FileTree, javaInfos *JavaClassInfos) {
    dirPageBack := conf.CurrentDirectory + conf.Irrigator 
    pageContent := daos.GetTemplBytes[JavaClassInfos](arg, config.BASE_IRRIGATOR, *javaInfos)
    daos.WriteToFile(pageContent, dirPageBack + javaInfos.ClassName + "Irrigator.java")
}

func appendRoute(arg string, conf *config.FileTree, javaInfos *JavaClassInfos) {

}
