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

func GetApplicationNameFromArtifactId(artifactId string) string {
    applicationName := ""
    posPrevSeparator := 0
    for i := 0; i < len(artifactId); i++ {
        nextIsUpper := false
        if artifactId[i] == '-' || artifactId[i] == '_' {
            applicationName += artifactId[posPrevSeparator : i]
            nextIsUpper = true
            for i < len(artifactId) && (artifactId[i] == '-' || artifactId[i] == '_') {
                i++
                posPrevSeparator = i+1
            }
        }
        if nextIsUpper && i < len(artifactId) {
            applicationName += string(unicode.ToUpper(rune(artifactId[i])))
        }
    }
    if posPrevSeparator < len(artifactId) - 1 {
            applicationName += artifactId[posPrevSeparator:]
    }
    return string(unicode.ToUpper(rune(applicationName[0]))) +  applicationName[1:] + "Application"
}

func GetDirNameFromPackage(p string) string {
    return strings.ReplaceAll(p, ".", "/") + "/"
}
