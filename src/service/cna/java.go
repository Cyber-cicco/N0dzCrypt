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
        {fileName: "AboutController.java", adress: fileTree.GetPageBackDir()},
        {fileName: "AuthenticationInfos.java", adress: fileTree.GetSecurityDir()},
        {fileName: "AuthenticationService.java", adress: fileTree.GetSecurityConfigDir()},
        {fileName: "BaseUser.java", adress: fileTree.GetEntityDir()},
        {fileName: "BaseUserRepository.java", adress: fileTree.GetRepositoryDir()},
        {fileName: "CustomAuthenticationEntryPoint.java", adress: fileTree.GetSecurityConfigDir()},
        {fileName: "CustomLogoutHandler.java", adress: fileTree.GetSecurityConfigDir()},
        {fileName: "CustomLogoutSuccessHandler.java", adress: fileTree.GetSecurityConfigDir()},
        {fileName: "HX.java", adress: fileTree.GetHXDir()},
        {fileName: "HomeController.java", adress: fileTree.GetPageBackDir()},
        {fileName: "JwtAuthenticationFilter.java", adress: fileTree.GetSecurityConfigDir()},
        {fileName: "JwtService.java", adress: fileTree.GetSecurityServiceDir()},
        {fileName: "LayoutIrrigator.java", adress: fileTree.GetIrrigatorDir()},
        {fileName: "LoginController.java", adress: fileTree.GetPageBackDir()},
        {fileName: "RoleType.java", adress: fileTree.GetEntityEnumDir()},
        {fileName: "UserRole.java", adress: fileTree.GetEntityDir()},
        {fileName: "UserRoleRepository.java", adress: fileTree.GetRepositoryDir()},
        {fileName: "Routes.java", adress: fileTree.GetPageBackDir()},
        {fileName: "WebSecurityConfig.java", adress: fileTree.GetSecurityConfigDir()},
        {fileName: "base.html", adress: fileTree.GetTemplateLayoutsDir()},
        {fileName: "headers.html", adress: fileTree.GetTemplateLayoutsDir()},
        {fileName: "navbar.html", adress: fileTree.GetTemplateComponentsDir()},
        {fileName: "tailwind.config.js", adress: fileTree.GetTemplateDir()},
        {fileName: "home.html", adress: fileTree.GetPageFrontDir() + "home/"},
        {fileName: "about.html", adress: fileTree.GetPageFrontDir() + "about/"},
        {fileName: "input.css", adress: fileTree.GetTemplateStyleDir()},
        {fileName: "output.css", adress: fileTree.GetStaticStyleDir()},
        {fileName: "n0dzCrypt.html", adress: fileTree.GetTemplateSVGDir()},
        {fileName: "htmx.min.js", adress: fileTree.GetTemplateJSDir()},
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
    f, err := os.OpenFile(fileTree.GetJavaDir() + mainClass, os.O_TRUNC|os.O_CREATE|os.O_WRONLY, 0644)
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

