package entity

type LaptopParams struct {
    UserInput  string
    Brand      string
    RAM        string
    CPU        string
    ScreenSize string
    OpenAIKey  string
}

type LaptopUsecase interface {
    RecommendLaptop(params LaptopParams) (string, error)
}