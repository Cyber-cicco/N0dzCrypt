package utils

import (
	"math/rand"
	"regexp"
	"strings"
	"unicode"
)

const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

func RandStringBytes(n int) string {
    b := make([]byte, n)
    for i := range b {
        b[i] = letterBytes[rand.Intn(len(letterBytes))]
    }
    return string(b)
}

func IsAcceptedCharacters(name string) bool {
    return regexp.MustCompile(`^[a-zA-Z0-9,\.,\-,_]*$`).MatchString(name)
}

func transformIntoPackageName(name string) string {
    packageName := ""
    posPrevSeparator := 0
    for i := 0; i < len(name); i++ {
        if name[i] == '-' || name[i] == '_' {
            if posPrevSeparator <= i  {
                packageName += name[posPrevSeparator : i]
            }
            posPrevSeparator = i + 1
            continue
        }
        if unicode.IsUpper(rune(name[i])) {
            if posPrevSeparator <= i  {
                packageName += name[posPrevSeparator : i] + string(unicode.ToLower(rune(name[i])))
            }
            posPrevSeparator = i + 1
        }
    }
    if posPrevSeparator < len(name) - 1 {
            packageName += name[posPrevSeparator:]
    }
    return packageName
}

func GetPackageName(artifactId, groupId string) string {
    return transformIntoPackageName(groupId) + "." + transformIntoPackageName(artifactId)
}

func GetUpperSnakeCaseFromKebab(content string) string {
    applicationName := ""
    for i := 0; i < len(content); i++ {
        applicationName += string(unicode.ToUpper(rune(content[i])))
        if content[i] == '-' {
            applicationName += "_"
        }
    }
    return applicationName
}
func GetUpperSnakeCaseFromDir(content string) string {
    applicationName := ""
    for i := 0; i < len(content); i++ {
        applicationName += string(unicode.ToUpper(rune(content[i])))
        if content[i] == '/' {
            applicationName += "_"
        }
    }
    return applicationName
}

func GetCamelCaseFromKebab(content string) string {
    applicationName := ""
    posPrevSeparator := 0
    for i := 0; i < len(content); i++ {
        nextIsUpper := false
        if content[i] == '-' || content[i] == '_' {
            applicationName += content[posPrevSeparator : i]
            nextIsUpper = true
            for i < len(content) && (content[i] == '-' || content[i] == '_') {
                i++
                posPrevSeparator = i+1
            }
        }
        if nextIsUpper && i < len(content) {
            applicationName += string(unicode.ToUpper(rune(content[i])))
        }
    }
    if posPrevSeparator < len(content) - 1 {
            applicationName += content[posPrevSeparator:]
    }
    return string(unicode.ToUpper(rune(applicationName[0]))) +  applicationName[1:]
}

func GetApplicationName(artifactId string) string {
    return GetCamelCaseFromKebab(artifactId) + "Application"
}

func GetDirNameFromPackage(p string) string {
    return strings.ReplaceAll(p, ".", "/") + "/"
}

func GetNewAdressVarFromName(routeName string) string {
    pageName := strings.Split(routeName, "/")
    return "ADR_" + GetUpperSnakeCaseFromDir(pageName[len(pageName) -1])
}
