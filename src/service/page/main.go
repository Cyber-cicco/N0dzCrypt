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
    conf, err := daos.GetConfigFile(currDir)
    utils.HandleUsageError(err, config.ERR_COULDNT_FIND_CONFIG)
    pageHTML, err := os.ReadFile(config.RESOURCE_FOLDER + "n0dzcrypt.page.html")
    fmt.Printf("string(pageHTML): %v\n", string(pageHTML))
    for _, arg := range args {
        dirName := conf.CurrentDirectory + conf.PagesFront + arg + "/"
        err := os.MkdirAll(dirName, os.ModePerm)
        utils.HandleTechnicalError(err, config.ERR_DIR_CREATION)
        daos.WriteToFile(pageHTML, dirName + "base.html")
    }
}
