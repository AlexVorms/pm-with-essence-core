package responseDTO

type ErrorResponse struct {
	Message string `json:"message" example:"Invalid project data"`
}

type LoginResponse struct {
	AccessToken string `json:"accessToken" example:"eyJhbGciOiJIUzI1NiIs..."`
	TokenType   string `json:"tokenType" example:"Bearer"`
}
type MessageResponse struct {
	Message string `json:"message"`
}
