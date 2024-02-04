package service

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
