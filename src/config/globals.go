package config

import "fr/hijokaidan/utils"

var RESOURCE_FOLDER = "../resources/"

type ProjectProps struct {
    MainPackage string
    HasMavenProfil bool
}


type FileTree struct {
    Root string
    JavaRoot string
    PagesBack string
    Irrigator string
    Service string
    Security string
    SecurityService string
    SecurityConfig string
    Repository string
    Utils string
    HX string
    Entities string
    Templates string
    Validators string
    Json string
    Static string
    StyleStatic string
    JSStatic string
    Img string
    Layouts string
    Components string
    JSTemplates string
    PagesFront string
    StyleTemplates string
    Test string
    TestResources string
}

func InitConfig(mainPackage string) *FileTree {
    root := "../../TestNSC/"
    javaDir := "src/main/java/" + utils.GetDirNameFromPackage(mainPackage)
    javaRoot := root + javaDir
    resources := root + "src/main/resources/"
    templates := resources + "templates/"
    static := resources + "static/"

    return &FileTree{
        Root: root,
        JavaRoot: javaRoot,
        PagesBack: javaRoot + "pages/",
        Irrigator: javaRoot + "pages/irrigator/",
        Service: javaRoot + "service/",
        Security: javaRoot + "security/",
        SecurityService: javaDir + "security/service/",
        SecurityConfig: javaDir + "security/config/",
        Repository: javaRoot + "repository/",
        Utils: javaRoot + "util/",
        HX: javaRoot + "util/hx/",
        Entities: javaRoot + "entity/",
        Validators: javaRoot + "validator/",
        Json: javaRoot + "json/",
        Templates: templates,
        Static: static,
        JSStatic: static + "script/",
        Img: static + "script/",
        JSTemplates: templates + "script/",
        Layouts: templates + "layout/",
        Components: templates + "components/",
        PagesFront: templates + "pages/",
        StyleStatic: static + "style/",
        StyleTemplates: templates + "style/",
        Test: root + "src/test/java/",
        TestResources: root + "src/test/resources/",
    }
}
