package cna

import (
	"bytes"
	"fr/nzc/config"
	"fr/nzc/utils"
	"os"
	"text/template"
)

type FileToAdress struct {
    fileName string
    adress string
}

func createBaseFiles(pom *Pom, fileTree *config.FileTree, mainPackage string) {
    javaFile := JavaFile{BasePackage: mainPackage, MainClass: pom.MainClass}
    files := []FileToAdress{
        {fileName: "AboutController.java", adress: fileTree.PagesBack},
        {fileName: "AuthenticationInfos.java", adress: fileTree.Security},
        {fileName: "AuthenticationService.java", adress: fileTree.SecurityService},
        {fileName: "BaseUser.java", adress: fileTree.Entities},
        {fileName: "BaseUserRepository.java", adress: fileTree.Repository},
        {fileName: "CustomAuthenticationEntryPoint.java", adress: fileTree.SecurityConfig},
        {fileName: "CustomLogoutHandler.java", adress: fileTree.SecurityConfig},
        {fileName: "CustomLogoutSuccessHandler.java", adress: fileTree.SecurityConfig},
        {fileName: "HX.java", adress: fileTree.HX},
        {fileName: "HomeController.java", adress: fileTree.PagesBack},
        {fileName: "JwtAuthenticationFilter.java", adress: fileTree.SecurityConfig},
        {fileName: "JwtService.java", adress: fileTree.SecurityService},
        {fileName: "LayoutIrrigator.java", adress: fileTree.Irrigator},
        {fileName: "LoginController.java", adress: fileTree.PagesBack},
        {fileName: "RoleType.java", adress: fileTree.EntityEnum},
        {fileName: "UserRole.java", adress: fileTree.Entities},
        {fileName: "UserRoleRepository.java", adress: fileTree.Repository},
        {fileName: "Routes.java", adress: fileTree.PagesBack},
        {fileName: "WebSecurityConfig.java", adress: fileTree.SecurityConfig},
        {fileName: "base.html", adress: fileTree.Layouts},
        {fileName: "headers.html", adress: fileTree.Layouts},
        {fileName: "navbar.html", adress: fileTree.Components},
        {fileName: "tailwind.config.js", adress: fileTree.Templates},
        {fileName: "home.html", adress: fileTree.PagesFront + "home/"},
        {fileName: "about.html", adress: fileTree.PagesFront + "about/"},
        {fileName: "input.css", adress: fileTree.StyleTemplates},
        {fileName: "output.css", adress: fileTree.StyleStatic},
        {fileName: "n0dzCrypt.html", adress: fileTree.SVG},
        {fileName: "htmx.min.js", adress: fileTree.JSTemplates},
    }
    for _, fileToAdress := range files {
        var tplBytes bytes.Buffer
        fileContent, err := os.ReadFile(config.RESOURCE_FOLDER + fileToAdress.fileName);
        utils.HandleTechnicalError(err, config.ERR_TEMPLATE_FILE_READ)
        tmpl, err := template.New(fileToAdress.fileName).Parse(string(fileContent))
        if err != nil { panic(err) }
        err = tmpl.Execute(&tplBytes, javaFile)
        createFile(&tplBytes, &fileToAdress)
    }
    var tplBytes bytes.Buffer
    fileContent, err := os.ReadFile(config.RESOURCE_FOLDER + "MainClass.java");
    utils.HandleTechnicalError(err, config.ERR_TEMPLATE_FILE_READ)
    mainClass := javaFile.MainClass + ".java"
    tmpl, err := template.New(mainClass).Parse(string(fileContent))
    if err != nil { panic(err) }
    err = tmpl.Execute(&tplBytes, javaFile)
    f, err := os.OpenFile(fileTree.JavaRoot + mainClass, os.O_TRUNC|os.O_CREATE|os.O_WRONLY, 0644)
    defer f.Close()
    if err != nil {
        panic(err)
    }
    f.Write(tplBytes.Bytes())
}

func createFile(content *bytes.Buffer, fileToAdress *FileToAdress) {
		f, err := os.OpenFile(fileToAdress.adress + fileToAdress.fileName, os.O_TRUNC|os.O_CREATE|os.O_WRONLY, 0644)
		defer f.Close()
		if err != nil {
			panic(err)
		}
		f.Write(content.Bytes())
}

