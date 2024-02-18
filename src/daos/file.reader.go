package daos

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"fr/nzc/config"
	"fr/nzc/utils"
	"io/fs"
	"io/ioutil"
	"os"
	"strings"
	"text/template"
)

/*
* fileFunc is a function that takes a string as parameter
*/
type fileFunc func(fileContent, filePath string)


func GetConfigFile(path string) *config.FileTree {

    path, fileName, err := GetConfigFilePath(path)
    utils.HandleUsageError(err, config.ERR_COULDNT_FIND_CONFIG)

    file, err := os.ReadFile(path + "/" + fileName)
    utils.HandleTechnicalError(err, config.ERR_OPEN_CONFIG)

    fileTree := &config.FileTree{
        ProjectAbsolutePath: path + "/",
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

func FileExists(path string) bool {
	_, err := os.Stat(path)
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

/* parseFolders 
* recursively processes files and directories within a specified path.
* 
* 
* This function recursively explores the directory structure starting from the provided 'path'.
* For each file encountered, it checks whether it is a directory or a file with the specified suffix.
* If the file is a directory, it continues to recursively process its contents.
* If the file matches the specified suffix, it reads its content and applies the 'executable' function to it.
* 
* Parameters 'path' and 'suffix' determine the scope of the search, and 'executable' defines the action to be performed on matching files.
* 
* @param files ([]fs.FileInfo): A list of FileInfo objects representing files and directories.
* @param path (string): The current directory path being processed.
* @param suffix (string): A file extension suffix used to filter files (e.g., ".txt", ".java").
* @param executable (fileFunc): A function that accepts a string as input and performs a specific operation.
*/
func parseFolders(files []fs.FileInfo, path, suffix string, executable fileFunc){
    for _, file := range files {
        if file.IsDir() {
            files, err := ioutil.ReadDir(path+"/"+file.Name())
            utils.HandleTechnicalError(err, config.ERR_CURR_DIR_OPEN)
            parseFolders(files, path+"/"+file.Name(), suffix, executable)
        } else if strings.HasSuffix(file.Name(), suffix) {
            filePath := path+"/"+file.Name()
            content, err := os.ReadFile(filePath)
            utils.HandleTechnicalError(err, config.ERR_CURR_DIR_OPEN)
            executable(string(content), filePath)
        }
    }
}
func ParseFolders(suffix, path string, executable fileFunc) {
	files, err := ioutil.ReadDir(path)
    utils.HandleTechnicalError(err, config.ERR_CURR_DIR_OPEN)
    parseFolders(files, path, suffix, executable)
}
