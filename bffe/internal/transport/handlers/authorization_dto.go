package handlers

type RegisterRequest struct {
	FirstName     string  `json:"firstName"`
	MiddleInitial *string `json:"middleInitial"`
	Surname       string  `json:"surname"`
	NameExtension *string `json:"nameExtension"`
	Email         string  `json:"email"`
	Username      string  `json:"username"`
	Password      string  `json:"password"`
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
