package handlers

type RegisterRequest struct {
	FirstName     string  `json:"firstName" validate:"required,max=100"`
	MiddleInitial *string `json:"middleInitial" validate:"max=10"`
	Surname       string  `json:"surname" validate:"required,max=100"`
	NameExtension *string `json:"nameExtension" validate:"max=10"`
	Username      string  `json:"username" validate:"required,max=255"`
	Password      string  `json:"password" validate:"required,min=6,max=25"`
}

type RegisterStatus string

const (
	RegisterSuccess RegisterStatus = "success"
	RegisterFail    RegisterStatus = "fail"
)

type RegisterResponse struct {
	Status      string  `json:"status"`
	FailMessage *string `json:"FailMessage"`
}
