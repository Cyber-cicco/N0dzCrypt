package page

import (
	"fmt"
	"fr/nzc/config"
	"fr/nzc/daos"
	"fr/nzc/utils"
	"os"
	"path/filepath"
)

func CreatePage(args []string) {
    currDir, err := filepath.Abs(".")
    utils.HandleTechnicalError(err, config.ERR_CURR_DIR_OPEN)
    conf := daos.GetConfigFile(currDir)
    pageHTML, err := os.ReadFile(config.RESOURCE_FOLDER + "n0dzcrypt.page.html")
    utils.HandleTechnicalError(err, config.ERR_TEMPLATE_FILE_READ)
    for _, arg := range args {
        javaInfos := &JavaClassInfos {
            BasePackage: conf.JavaBack.BasePackage,
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
    dirPageFront := conf.CurrentDirectory + conf.GetPageFrontDir() + arg + "/"
    err := os.MkdirAll(dirPageFront, os.ModePerm)
    utils.HandleTechnicalError(err, config.ERR_DIR_CREATION)
    newFile := dirPageFront + arg + ".html"
    daos.WriteToFile(pageHTML, newFile)
    fmt.Println("written file : " + newFile)
}

func writePageBack(arg string, conf *config.FileTree, javaInfos *JavaClassInfos) {
    dirPageBack := conf.CurrentDirectory + conf.GetPageBackDir() 
    pageContent := daos.GetTemplBytes[JavaClassInfos](arg, config.BASE_PAGE_BACK, *javaInfos)
    newFile := dirPageBack + javaInfos.ClassName + "Controller.java"
    daos.WriteToFile(pageContent, newFile)
    fmt.Println("written file : " + newFile)
}

func writeIrrigator(arg string, conf *config.FileTree, javaInfos *JavaClassInfos) {
    dirPageBack := conf.CurrentDirectory + conf.GetIrrigatorDir() 
    pageContent := daos.GetTemplBytes[JavaClassInfos](arg, config.BASE_IRRIGATOR, *javaInfos)
    newFile := dirPageBack + javaInfos.ClassName + "Irrigator.java"
    daos.WriteToFile(pageContent, newFile)
    fmt.Println("written file : " + newFile)
}

func appendRoute(arg string, conf *config.FileTree, javaInfos *JavaClassInfos) {
    pageContent := daos.GetTemplBytes[JavaClassInfos](arg, config.SINGLE_ROUTE, *javaInfos)
    routesPath := conf.CurrentDirectory + conf.GetPageBackDir() + "Routes.java"
    routesBytes, err := os.ReadFile(routesPath)
    utils.HandleTechnicalError(err, config.ERR_TEMPLATE_FILE_READ)
    openBC := 0
    closeBC := 0
    newRoutesContent := ""
    for i, char := range routesBytes {
        if char == '{' {
            openBC++
        } else if char == '}' {
            closeBC++
        }
        if char == '}' && openBC == closeBC {
            newRoutesContent = string(routesBytes[:i]) + string(pageContent) + string(routesBytes[i:])
        }
    }
    daos.WriteToFile([]byte(newRoutesContent), routesPath)
    conf.Routes = append(conf.Routes, config.Route{
        RouteName: "ADR_"+ javaInfos.UpperClassName,
        CorrespondingRoute: "page/" + javaInfos.PageName + "/" + javaInfos.PageName,
    })
    path, fileName, err := daos.GetConfigFilePath(conf.CurrentDirectory)
    utils.HandleUsageError(err, config.ERR_COULDNT_FIND_CONFIG)
    config.WriteFileTree(conf, path + "/" + fileName)
    fmt.Println("Updated Routes.java and n0dzcrypt.json")
}
