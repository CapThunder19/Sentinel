package service

type RegisterRequest struct {
	Username string
	Email    string
	Password string
}

type RegisterResponse struct {
	ID       string
	Username string
	Email    string
}

type LoginRequest struct {
	Email    string
	Password string
}

type LoginResponse struct {
	AccessToken string
	TokenType   string
}
