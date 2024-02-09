package cna

import (
	"fmt"
	"fr/hijokaidan/config"
)

type FileToAdress struct {
    fileName string
    adress string
}

func createFiles(pom *Pom, fileTree config.FileTree) {
    files := []FileToAdress{
        {fileName: "AboutController.txt", adress: fileTree.PagesBack},
        {fileName: "AuthenticationInfos.txt", adress: fileTree.Security},
        {fileName: "AuthenticationService.txt", adress: fileTree.PagesBack},
        {fileName: "BaseUser.txt", adress: fileTree.Entities},
        {fileName: "BaseUserRepository.txt", adress: fileTree.Repository},
        {fileName: "CustomAuthenticationEntryPoint.txt", adress: fileTree.SecurityConfig},
        {fileName: "CustomLogoutHandler.txt", adress: fileTree.SecurityConfig},
        {fileName: "CustomLogoutSuccessHandler.txt", adress: fileTree.SecurityConfig},
        {fileName: "HX.txt", adress: fileTree.HX},
        {fileName: "HomeController.txt", adress: fileTree.PagesBack},
        {fileName: "JwtAuthenticationFilter.txt", adress: fileTree.SecurityConfig},
        {fileName: "JwtService.txt", adress: fileTree.SecurityService},
        {fileName: "LayoutIrrigator.txt", adress: fileTree.Irrigator},
        {fileName: "LoginController.txt", adress: fileTree.PagesBack},
        {fileName: "ProfileController.txt", adress: fileTree.PagesBack},
        {fileName: ".txt", adress: fileTree.Irrigator},
        {fileName: ".txt", adress: fileTree.Irrigator},
        {fileName: ".txt", adress: fileTree.Irrigator},
        {fileName: ".txt", adress: fileTree.Irrigator},
        {fileName: ".txt", adress: fileTree.Irrigator},
        {fileName: ".txt", adress: fileTree.Irrigator},
        {fileName: ".txt", adress: fileTree.Irrigator},
        {fileName: ".txt", adress: fileTree.Irrigator},
    }
    for _, fileToAdress := range files {
        fmt.Println(fileToAdress)
    }
}
