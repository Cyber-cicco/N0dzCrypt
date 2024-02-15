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

type PageBackDir struct {
    RootDir string `json:"root"`
    Controller string `json:"controller"`
    Irrigator string `json:"irrigator"`
}

type UtilsDir struct {
    RootDir string `json:"root"`
    HX string`json:"hx"`
}

type Domain struct {
    Name string
    SubDomains []Domain
}

type SecurityDir struct {
    RootDir string `json:"root"`
    Service string`json:"service"`
    Config string`json:"config"`
}

type EntityDir struct {
    RootDir string `json:"root"`
    Enum string`json:"enum"`
}

type JsonDir struct {
    RootDir string `json:"root"`
    Dto string`json:"dto"`
    Mapper string`json:"mapper"`
}

type JavaBack struct {
    RootDir string `json:"root"`
    BasePackage string`json:"basepackage"`
    PagesDir PageBackDir`json:"page"` 
    Service string `json:"service"`
    Security SecurityDir `json:"security"`
    Repository string`json:"repository"`
    Utils UtilsDir `json:"util"`
    Entities EntityDir`json:"entities"`
    Validators string`json:"validators"`
    Json JsonDir`json:"json"`
}

type FrontTemplates struct {
    RootDir string `json:"root"`
    PagesFront string`json:"pages"`
    Components string`json:"components"`
    Layouts string`json:"layouts"`
    SVG string`json:"svg"`
    JS string`json:"javascript"`
    StyleTemplates string`json:"styletemplates"`
}


type Resources struct {
    RootDir string `json:"root"`
    Templates FrontTemplates`json:"templates"`
    Static Static`json:"static"`
}

type Static struct {
    RootDir string `json:"root"`
    Style string`json:"style"`
    JS string`json:"jsstatic"`
    Img string`json:"img"`
}

type JavaTest struct {
    RootDir string `json:"root"`
    TestRessources string `json:"resources"`
    TestMain string `json:"base-dir"`
}


func (f FileTree) GetJavaDir() string{
    return "./" + f.JavaBack.RootDir 
}

func (f FileTree) GetPageBackDir() string{
    return "./" + f.JavaBack.RootDir + f.JavaBack.PagesDir.RootDir
}

func (f FileTree) GetJsonDir() string{
    return "./" + f.JavaBack.RootDir + f.JavaBack.Json.RootDir
}

func (f FileTree) GetEntityDir() string{
    return "./" + f.JavaBack.RootDir + f.JavaBack.Entities.RootDir
}
func (f FileTree) GetSecurityDir() string{
    return "./" + f.JavaBack.RootDir + f.JavaBack.Security.RootDir
}
func (f FileTree) GetUtilsDir() string{
    return "./" + f.JavaBack.RootDir + f.JavaBack.Utils.RootDir
}

func (f FileTree) GetResourcesDir() string{
    return "./" + f.Resources.RootDir 
}
func (f FileTree) GetTemplateDir() string{
    return "./" + f.Resources.RootDir + f.Resources.Templates.RootDir
}

func (f FileTree) GetPageFrontDir() string{
    return "./" + f.Resources.RootDir + f.Resources.Templates.RootDir + f.Resources.Templates.PagesFront
}

func (f FileTree) GetTestDir() string{
    return "./" + f.Resources.RootDir + f.Resources.Static.RootDir
}
func (f FileTree) GetStaticDir() string{
    return "./" + f.Resources.RootDir + f.Resources.Static.RootDir
}


type FileTree struct {
    JavaBack JavaBack `json:"java-back"`
    Resources Resources`json:"resources"`
    JavaTest JavaTest`json:"java-test"`
    CurrentDirectory string`json:"-"`
    PageDomains []Domain `json:"page-domains"`
    Routes []Route `json:"routes"`
}

type Route struct {
    RouteName string`json:"name"`
    CorrespondingRoute string`json:"route"`
}

func CreateConfig(mainPackage string) *FileTree {
    javaDir := "src/main/java/" + utils.GetDirNameFromPackage(mainPackage)

    fileTree := &FileTree{
        JavaBack: JavaBack{
            RootDir:javaDir,
            BasePackage:mainPackage,
            PagesDir:PageBackDir{},
            Service:"service/",
            Security:SecurityDir{
                RootDir:"security/",
                Service:"service/",
                Config:"config/",
            },
            Repository:"repository/",
            Utils:UtilsDir{
                RootDir:"util/",
                HX:"hx/",
            },
            Entities:EntityDir{
                RootDir:"entity/",
                Enum:"enums/",
            },
            Validators:"validator/",
            Json:JsonDir{
                RootDir:"json/",
                Dto:"dto/",
                Mapper:"mapper/",
            },
        },
        Resources: Resources{
            RootDir:"src/main/resources",
            Templates:FrontTemplates{
                RootDir:"templates/",
                PagesFront:PageFront{
                    RootDir: "pages/",
                    Fragments: "fragment/",
                },
                Components:"component/",
                Layouts:"layout/",
                SVG:"svg/",
                JS:"script/",
                StyleTemplates:"style/",
            },
            Static:Static{
                RootDir:"static/",
                Style:"style/",
                JS:"script/",
                Img:"img/",
            },
        },
        JavaTest:JavaTest{
            RootDir:"src/test/java",
        },
        Routes: []Route{
            {
                RouteName: "ADR_BASE_LAYOUT",
                CorrespondingRoute: "layout/base",
            },
            {
                RouteName: "ADR_HOME",
                CorrespondingRoute: "page/home/home",
            },
            {
                RouteName: "ADR_ABOUT",
                CorrespondingRoute: "page/about/about",
            },
            {
                RouteName: "ADR_LOGIN",
                CorrespondingRoute: "page/login/login",
            },
            {
                RouteName: "ADR_FORM_ERROR",
                CorrespondingRoute: "components/form-error",
            },
        },
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

