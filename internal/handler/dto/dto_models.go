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
	Theme   string `json:"theme"`
	Reason  int    `json:"reason"`
	Message string `json:"message"`
}

type EditEmailOrLogin struct {
	NewValue string `json:"new_value"`
}

type ResetPasswordRequest struct {
	Email string `json:"email"`
}

type VerifyOtp struct {
	Email   string `json:"email"`
	OtpCode int64  `json:"otp_code"`
}

type NewPassword struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}
