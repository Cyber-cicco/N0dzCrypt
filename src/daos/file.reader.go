package daos

import (
	"encoding/json"
	"errors"
	"fr/nzc/config"
	"fr/nzc/utils"
	"os"
)


func GetConfigFile(path string) (*config.FileTree, error) {

    if path == "" {
        return nil, errors.New("Couldn't find n0dzCrypt.json")
    }

    if !FileExists(path) {
        return nil, errors.New("Directory does not exist")
    }

    dir, err := os.ReadDir(path)

    utils.HandleTechnicalError(err, config.ERR_CURR_DIR_OPEN)
    
    for _, file := range dir {
        if file.Name() == config.CONFIG_FILE {
            file, err := os.ReadFile(path + "/" + file.Name())
            utils.HandleTechnicalError(err, config.ERR_OPEN_CONFIG)
            fileTree := &config.FileTree{
                CurrentDirectory: path + "/",
            }
            json.Unmarshal(file, &fileTree)
            return fileTree, nil
        }
    }

    return GetConfigFile(walkToParentDirectory(path))
}

func walkToParentDirectory(path string) string {
    indexOfLastSlash := 0;
    i := 0
    for i < len(path) {
        if path[i] == '\\' {
            i += 2
            continue
        }
        if path[i] == '/' {
            indexOfLastSlash = i
        }
        i += 1
    }
    return path[:indexOfLastSlash]
}

func FileExists(dirPath string) bool {
	_, err := os.Stat(dirPath)
	if err != nil {
		if os.IsNotExist(err) {
			return false
		}
        panic(err)
	}
	return true
}
