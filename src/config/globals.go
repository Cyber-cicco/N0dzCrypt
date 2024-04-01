package config

import (
	"encoding/json"
	"fr/nzc/utils"
	"os"
	"path/filepath"
	"strings"
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


func (f FileTree) GetRepositoryDir() string{
    return "./" + f.GetRootOfJavaProject()  + f.JavaBack.Repository
}
func (f FileTree) GetServiceDir() string{
    return "./" + f.GetRootOfJavaProject() + f.JavaBack.Service
}
func (f FileTree) GetValidatorDir() string{
    return "./" + f.GetRootOfJavaProject() + f.JavaBack.Validators
}
func (f FileTree) GetJavaDir() string{
    return "./" + f.GetRootOfJavaProject() 
}
func (f FileTree) GetPageBackDir() string{
    return "./" + f.GetRootOfJavaProject() + f.JavaBack.PagesDir.RootDir
}
func (f *FileTree) GetRootOfJavaProject() string {
    return f.JavaBack.RootDir + strings.ReplaceAll(f.JavaBack.BasePackage, ".", "/") + "/" 
}
func (f FileTree) GetIrrigatorDir() string {
    return "./" + f.GetRootOfJavaProject() + f.JavaBack.PagesDir.RootDir + f.JavaBack.PagesDir.Irrigator
}
func (f FileTree) GetJsonMapperDir() string{
    return "./" + f.GetRootOfJavaProject() + f.JavaBack.Json.RootDir + f.JavaBack.Json.Mapper
}
func (f FileTree) GetJsonDtoDir() string{
    return "./" + f.GetRootOfJavaProject() + f.JavaBack.Json.RootDir + f.JavaBack.Json.Dto
}
func (f FileTree) GetJsonDir() string{
    return "./" + f.GetRootOfJavaProject() + f.JavaBack.Json.RootDir
}
func (f FileTree) GetEntityDir() string{
    return "./" + f.GetRootOfJavaProject() + f.JavaBack.Entities.RootDir
}
func (f FileTree) GetEntityEnumDir() string{
    return "./" + f.GetRootOfJavaProject() + f.JavaBack.Entities.RootDir + f.JavaBack.Entities.Enum
}
func (f FileTree) GetSecurityDir() string{
    return "./" + f.GetRootOfJavaProject() + f.JavaBack.Security.RootDir
}
func (f FileTree) GetSecurityConfigDir() string{
    return "./" + f.GetRootOfJavaProject() + f.JavaBack.Security.RootDir + f.JavaBack.Security.Config
}
func (f FileTree) GetSecurityServiceDir() string{
    return "./" + f.GetRootOfJavaProject() + f.JavaBack.Security.RootDir + f.JavaBack.Security.Service
}
func (f FileTree) GetUtilsDir() string{
    return "./" + f.GetRootOfJavaProject() + f.JavaBack.Utils.RootDir
}
func (f FileTree) GetHXDir() string{
    return "./" + f.GetRootOfJavaProject() + f.JavaBack.Utils.RootDir + f.JavaBack.Utils.HX
}
func (f FileTree) GetResourcesDir() string{
    return "./" + f.Resources.RootDir 
}
func (f FileTree) GetTemplateSVGDir() string{
    return "./" + f.Resources.RootDir + f.Resources.Templates.RootDir + f.Resources.Templates.SVG
}
func (f FileTree) GetTemplateStyleDir() string{
    return "./" + f.Resources.RootDir + f.Resources.Templates.RootDir + f.Resources.Templates.StyleTemplates
}
func (f FileTree) GetTemplateDir() string{
    return "./" + f.Resources.RootDir + f.Resources.Templates.RootDir
}
func (f FileTree) GetTemplateComponentsDir() string{
    return "./" + f.Resources.RootDir + f.Resources.Templates.RootDir + f.Resources.Templates.Components
}
func (f FileTree) GetTemplateLayoutsDir() string{
    return "./" + f.Resources.RootDir + f.Resources.Templates.RootDir + f.Resources.Templates.Layouts
}
func (f FileTree) GetTemplateJSDir() string{
    return "./" + f.Resources.RootDir + f.Resources.Templates.RootDir + f.Resources.Templates.JS
}
func (f FileTree) GetPageFrontDir() string{
    return "./" + f.Resources.RootDir + f.Resources.Templates.RootDir + f.Resources.Templates.PagesFront
}
func (f FileTree) GetTestDir() string{
    return "./" + f.Resources.RootDir + f.Resources.Static.RootDir
}
func (f FileTree) GetTestRessourcesDir() string{
    return "./" + f.Resources.RootDir + f.Resources.Static.RootDir + f.JavaTest.TestMain + f.JavaTest.TestRessources
}
func (f FileTree) GetTestMainDir() string{
    return "./" + f.Resources.RootDir + f.Resources.Static.RootDir + f.JavaTest.TestMain
}
func (f FileTree) GetStaticStyleDir() string{
    return "./" + f.Resources.RootDir + f.Resources.Static.RootDir + f.Resources.Static.Style
}
func (f FileTree) GetStaticImgDir() string{
    return "./" + f.Resources.RootDir + f.Resources.Static.RootDir + f.Resources.Static.Img
}
func (f FileTree) GetStaticJsDir() string{
    return "./" + f.Resources.RootDir + f.Resources.Static.RootDir + f.Resources.Static.JS
}
func (f FileTree) GetStaticDir() string{
    return "./" + f.Resources.RootDir + f.Resources.Static.RootDir
}
func (f FileTree) GetFragmentReference(fragmentPath string) string{
    var fragmentReference string
    absPath, err := filepath.Abs(fragmentPath)
    utils.HandleUsageError(err, "Couldn't find or open file")
    splitPath := strings.Split(absPath, f.Resources.Templates.RootDir)
    fragmentReference = splitPath[len(splitPath)-1]
    fragmentReference = utils.StripPathFromDirectoryReference(fragmentReference)
    if strings.HasSuffix(fragmentReference,  ".html") {
        fragmentReference = strings.Split(fragmentReference, ".")[0]
    }
    return fragmentReference
}

type FileTree struct {
    JavaBack JavaBack `json:"java-back"`
    Resources Resources`json:"resources"`
    JavaTest JavaTest`json:"java-test"`
    ProjectAbsolutePath string`json:"-"`
    PageDomains []Domain `json:"page-domains"`
    Routes []Route `json:"routes"`
}

type Route struct {
    RouteName string`json:"name"`
    CorrespondingRoute string`json:"route"`
}

func CreateConfig(mainPackage string) *FileTree {
    javaDir := "src/main/java/"

    fileTree := &FileTree{
        JavaBack: JavaBack{
            RootDir:javaDir,
            BasePackage:mainPackage,
            PagesDir:PageBackDir{
                RootDir:"page/",
                Controller:"",
                Irrigator:"irrigator/",
            },
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
            RootDir:"src/main/resources/",
            Templates:FrontTemplates{
                RootDir:"templates/",
                PagesFront:"page/",
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
            RootDir:"src/test/java/",
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
    return WriteFileTree(fileTree, CONFIG_FILE)
}


func WriteFileTree(fileTree *FileTree, route string) *FileTree {
    configTree, err := json.MarshalIndent(fileTree, "", "  ")
    utils.HandleTechnicalError(err, ERR_MARSHARL)
    f, err := os.OpenFile(route, os.O_TRUNC|os.O_CREATE|os.O_WRONLY, 0644)
    defer f.Close()
    if err != nil {
        panic(err)
    }
    f.Write(configTree)
    return fileTree
}
