package service

type SQLConnection struct {
    Url string
    driver string
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

type Profile struct {
    DBUrl string
    DBUser string
    DBPassword string
}
