package java

import (
	"fmt"
	"testing"
)


func TestChangeRouteFile(t *testing.T) {
    var sourceCode = `
package fr.diginamic.digilearning.page;

/**
 * THIS CLASS CAN BE OVERWRITTEN BY N0DZCRYPT
 * IF YOU INTEND TO USE THE N0DZCRYPT CLI IN YOUT APP, TRY NOT TO CHANGE IT.
 *
 * Contains constants pointing to thymeleaf fragments
 */
public class Routes {

    public static final String ADR_BASE_LAYOUT = "layout/base";
    public static final String ADR_HOME = "page/home/home";
    public static final String ADR_ABOUT = "page/about/about";
    public static final String ADR_LOGIN = "page/about/about";
    public static final String ADR_FORM_ERROR = "components/form-error";
    public static final String ADR_TEST = "page/test/test";
    public static final Integer TEST_1 = 1

    public void caca() {
        String caca = "caca"
    }
}
`
    oldname := "page/about/about"
    newname := "page/about/infos"
    changeRouteInRoutesFile(oldname, newname, sourceCode)
}

func TestChangeNameJavaFile(t *testing.T) {
    sourceCode := `
package {{.BasePackage}}.page;

import org.springframework.stereotype.Controller;
import org.springframework.ui.Model;
import org.springframework.web.bind.annotation.GetMapping;
import org.springframework.web.bind.annotation.RequestMapping;
import lombok.RequiredArgsConstructor;

@Controller
@RequiredArgsConstructor
@RequestMapping("/home")
public class HomeController {

    @GetMapping
    public String getHome(Model model) {
        model.addAttribute("routerOutlet", Routes.ADR_HOME);
        return Routes.ADR_BASE_LAYOUT;
    }

    @GetMapping("/partial")
    public String getHomePart(Model model) {
        return Routes.ADR_ABOUT;
    }
}

    ` 
    oldname := "page/about/about"
    newname := "page/about/infos"
    newContent := changeNameInJavaFile(oldname, newname, sourceCode)
    fmt.Println(newContent)
}

func TestChangeControllerAdress(t *testing.T) {
    sourceCode := `
package test.page;

import org.springframework.stereotype.Controller;
import org.springframework.ui.Model;
import org.springframework.web.bind.annotation.GetMapping;
import org.springframework.web.bind.annotation.RequestMapping;
import lombok.RequiredArgsConstructor;

@Controller
@RequiredArgsConstructor
@RequestMapping(value = "/home", test="/caca")
public class HomeController {

    @GetMapping
    public String getHome(Model model) {
        model.addAttribute("routerOutlet", Routes.ADR_HOME);
        return Routes.ADR_BASE_LAYOUT;
    }

    @GetMapping("/partial")
    public String getHomePart(Model model) {
        return Routes.ADR_ABOUT;
    }

    @GetMapping(value = "/profile/{id}", testValue = "testString")
    public String getHome(Model model) {
        model.addAttribute("routerOutlet", Routes.ADR_HOME);
        return Routes.ADR_BASE_LAYOUT;
    }

    @GetMapping("/profile/partial/{id}")
    public String getHomePart(Model model) {
        return Routes.ADR_ABOUT;
    }

    @PostMapping("/profile")
    public String getHomePart(Model model) {
        return Routes.ADR_ABOUT;
    }
}
    ` 
    newCode := changeControllerAdress("page/home", "page/accueil", sourceCode)
    fmt.Printf("newCode: %v\n", newCode)
}
