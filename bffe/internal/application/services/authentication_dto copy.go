package appservices

type RegisterRequest struct {
	FirstName     string  `json:"firstName" validate:"required,max=100"`
	MiddleInitial *string `json:"middleInitial" validate:"max=10"`
	Surname       string  `json:"surname" validate:"required,max=100"`
	NameExtension *string `json:"nameExtension" validate:"max=10"`
	Username      string  `json:"username" validate:"required,min=6,max=255"`
	Password      string  `json:"password" validate:"required,min=6,max=25"`
}

type RegisterStatus string

const (
	RegisterSuccess RegisterStatus = "success"
	RegisterFail    RegisterStatus = "fail"
)

type RegisterFailure int

const (
	UsernameExistFailure        RegisterFailure = 1
	NotAllowedCharactersFailure RegisterFailure = 2
)

type RegisterResponse struct {
	Status      RegisterStatus    `json:"status"`
	FailMessage []RegisterFailure `json:"failMessage"`
}
