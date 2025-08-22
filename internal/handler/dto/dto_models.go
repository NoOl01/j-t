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
