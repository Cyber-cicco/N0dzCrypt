package config

import (
	"encoding/json"
	"fr/nzc/utils"
	"os"
)

var RESOURCE_FOLDER = os.Getenv("GOPATH") + "/bin/resources/"
var CONFIG_FILE = "n0dzcrypt.json"
var BASE_PAGE_BACK = RESOURCE_FOLDER + "BaseController.java"
var BASE_IRRIGATOR = RESOURCE_FOLDER + "BaseIrrigator.java"
var SINGLE_ROUTE = RESOURCE_FOLDER + "SingleRoute.java"
var ROUTES = RESOURCE_FOLDER + "Routes.java"

type ProjectProps struct {
    MainPackage string
    HasMavenProfil bool
}


type FileTree struct {
    Root string `json:"root"`
    JavaRoot string `json:"javaroot"`
    PagesBack string `json:"pagesback"`
    Irrigator string`json:"irrigator"`
    SVG string`json:"svg"`
    Service string`json:"service"`
    Security string`json:"security"`
    SecurityService string`json:"securityservice"`
    SecurityConfig string`json:"securityconfig"`
    Repository string`json:"repository"`
    Utils string`json:"utils"`
    HX string`json:"hx"`
    Entities string`json:"entities"`
    EntityEnum string`json:"entityenum"`
    Templates string`json:"templates"`
    Validators string`json:"validators"`
    Json string`json:"json"`
    Static string`json:"static"`
    StyleStatic string`json:"stylestatic"`
    JSStatic string`json:"jsstatic"`
    Img string`json:"img"`
    Layouts string`json:"layouts"`
    Resources string`json:"resources"`
    Components string`json:"components"`
    JSTemplates string`json:"jstemplates"`
    PagesFront string`json:"pagesfront"`
    StyleTemplates string`json:"styletemplates"`
    Test string`json:"test"`
    TestResources string`json:"testresources"`
    BasePackage string`json:"basepackage"`
    CurrentDirectory string
}

func CreateConfig(mainPackage string) *FileTree {
    root := "./"
    javaDir := "src/main/java/" + utils.GetDirNameFromPackage(mainPackage)
    javaRoot := root + javaDir
    resources := root + "src/main/resources/"
    templates := resources + "templates/"
    static := resources + "static/"

    fileTree := &FileTree{
        Root: root ,
        JavaRoot: javaRoot,
        PagesBack: javaRoot + "page/",
        Irrigator: javaRoot + "page/irrigator/",
        Service: javaRoot + "service/",
        Security: javaRoot + "security/",
        SecurityService: javaRoot + "security/service/",
        SecurityConfig: javaRoot + "security/config/",
        Repository: javaRoot + "repository/",
        Utils: javaRoot + "util/",
        HX: javaRoot + "util/hx/",
        Entities: javaRoot + "entity/",
        EntityEnum: javaRoot + "entity/enums/",
        Validators: javaRoot + "validator/",
        Json: javaRoot + "json/",
        Templates: templates,
        Resources: resources,
        Static: static,
        JSStatic: static + "script/",
        Img: static + "script/",
        SVG: templates + "svg/",
        JSTemplates: templates + "script/",
        Layouts: templates + "layout/",
        Components: templates + "components/",
        PagesFront: templates + "page/",
        StyleStatic: static + "style/",
        StyleTemplates: templates + "style/",
        Test: root + "src/test/java/",
        BasePackage: mainPackage,
        TestResources: root + "src/test/resources/",
    }
    configTree, err := json.MarshalIndent(fileTree, "", "  ")
    utils.HandleTechnicalError(err, ERR_MARSHARL)
    f, err := os.OpenFile(CONFIG_FILE, os.O_TRUNC|os.O_CREATE|os.O_WRONLY, 0644)
    defer f.Close()
    if err != nil {
        panic(err)
    }
    f.Write(configTree)
    return fileTree
}

