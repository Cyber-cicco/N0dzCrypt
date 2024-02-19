package java

import (
	"testing"
)

func TestChangeRouteFile(t *testing.T) {
    sourceCode := `
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
