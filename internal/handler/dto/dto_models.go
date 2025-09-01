package dto

type LoginBody struct {
	LoginOrEmail string `json:"login_or_email"`
	Password     string `json:"password"`
}

type RegisterBody struct {
	Login    string `json:"login"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type AppealBody struct {
	Name    string `json:"name"`
	Email   string `json:"email"`
	Theme   int    `json:"theme"`
	Message string `json:"message"`
}
