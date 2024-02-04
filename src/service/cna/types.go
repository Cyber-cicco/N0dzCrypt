package cna

type SQLConnection struct {
    Url string
    Driver string
}

type Pom struct {
    SpringVersion string
    GroupId string
    ArtifactId string
    ProjectName string
    Description string
    JavaVersion int
    JwtVersion string
    AdditionalProperties []string
    MainClass string
    Profiles []string
}

type DBInfos struct {
    DBUrl string
    DBUser string
    DBPassword string
    DBDriver string
}

type Profile struct {
    ProfileName string
    ActiveByDefault bool
    DBInfos DBInfos
}

type ApplicationProperties struct {
    JWTSecret string
}

type N0dzCryptProject struct {
    //properties
    Pom string
    ApplicationProperties string

    //backend entities
    BaseUser string
    UserRole string
    RoleType string

    //backend repos
    BaseUserRepository string
    UserRoleRepository string

    //htmx constants
    HX string

    //routes
    Routes string

    //security config
    CustomAuthenticationEntryPoint string
    CustomLogoutHandler string
    CustomLogoutSuccessHandler string
    JwtAuthenticationFilter string
    WebSecurityConfig string
    AuthenticationService string
    AuthenticationInfos string
    JwtService string
    

    //backend controllers
    LoginController string
    ProfileController string
    AboutController string
    HomeController string

    //backend irrigators
    LayoutIrrigator string
    ProfileIrrigator string

    //backend main class
    MainClass string

    //front end templates
    BaseLayout string
    Headers string
    NavBar string
    HomePage string
    //LoginPage string

    //front end style
    InputCSS string

}

type JavaFile struct {
    BasePackage string
}
