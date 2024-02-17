package daos

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"fr/nzc/config"
	"fr/nzc/utils"
	"os"
	"text/template"
)


func GetConfigFile(path string) *config.FileTree {

    path, fileName, err := GetConfigFilePath(path)
    utils.HandleUsageError(err, config.ERR_COULDNT_FIND_CONFIG)

    file, err := os.ReadFile(path + "/" + fileName)
    utils.HandleTechnicalError(err, config.ERR_OPEN_CONFIG)

    fileTree := &config.FileTree{
        CurrentDirectory: path + "/",
    }

    json.Unmarshal(file, &fileTree)
    return fileTree
}

func GetConfigFilePath(path string) (string, string, error) {

    if path == "" {
        return "", "", errors.New("Couldn't find n0dzCrypt.json")
    }

    if !FileExists(path) {
        return "","", errors.New("Directory does not exist")
    }

    dir, err := os.ReadDir(path)

    utils.HandleTechnicalError(err, config.ERR_CURR_DIR_OPEN)
    
    for _, file := range dir {
        if file.Name() == config.CONFIG_FILE {
            return path, file.Name(), nil
        }
    }

    return GetConfigFilePath(walkToParentDirectory(path))
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

func GetTemplBytes[V any](name, fileName string, linkedStruct V) []byte {
    var tplBytes bytes.Buffer
    fmt.Printf("linkedStruct: %v\n", linkedStruct)
    fileContent, err := os.ReadFile(fileName);
    utils.HandleTechnicalError(err, config.ERR_TEMPLATE_FILE_READ)
    tmpl, err := template.New(name).Parse(string(fileContent))
    utils.HandleTechnicalError(err, config.ERR_TEMPLATE_FILE_READ)
    err = tmpl.Execute(&tplBytes, linkedStruct)
    return tplBytes.Bytes()
}
