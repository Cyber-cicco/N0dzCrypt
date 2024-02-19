package mv

import (
	"errors"
	"fmt"
	"fr/nzc/config"
	"fr/nzc/daos"
	"fr/nzc/java"
	"fr/nzc/thymeleaf"
	"fr/nzc/utils"
	"os"
	"path/filepath"
)

func MovePage(args []string, flag string) {
    currDir, err := filepath.Abs(".")
    utils.HandleTechnicalError(err, config.ERR_CURR_DIR_OPEN)
    conf := daos.GetConfigFile(currDir)
    if len(args) != 2 {
        utils.HandleUsageError(errors.New("Usage error"), "Error : mv command requires exactly two arguments")
    }
    oldname := args[0]
    newName := args[1]
    switch flag {
    case "" : {
        handleBaseCase(oldname, newName, conf)
    }
    case "page" : {

    }
    case "comp" : {

    }
    case "frag" : {

    }
    case "layout" : {

    }
    case "tstyle" : {

    }
    case "tscript" : {

    }
    case "svg" : {

    }
    case "style" : {

    }
    case "script" : {

    }
    case "img" : {

    }
    default : {
        utils.HandleUsageError(errors.New("Unsupported file type"), config.ERR_BAD_MV_ARGS)
    }
    }
}

func handleBaseCase(oldName, newName string, fileTree *config.FileTree) {
    file, err := os.ReadFile(oldName)
    utils.HandleUsageError(err, config.ERR_FILE_DOES_NOT_EXIST)
    if daos.FileExists(newName) {
        utils.HandleUsageError(errors.New("Can't move file"), config.ERR_MOVING_FILE)
    }
    daos.WriteToFile(file, newName)
    os.Remove(oldName)
    oldName = fileTree.GetFragmentReference(oldName)
    newName = fileTree.GetFragmentReference(newName)
    thymeleaf.RenameProjectFiles(oldName, newName, fileTree)
    java.RenameRoute(oldName, newName, fileTree)
}
