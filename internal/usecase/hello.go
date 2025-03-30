package usecase

type HelloOutputDTO struct {
	Message string `json:"message"`
}

type HelloUseCase struct{}

func NewHelloUseCase() *HelloUseCase {
	return &HelloUseCase{}
}

func (h *HelloUseCase) Execute() HelloOutputDTO {
	return HelloOutputDTO{
		Message: "Hello, World!",
	}
}
