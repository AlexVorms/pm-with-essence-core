package userDTO

type RegisterDTO struct {
	UserName string `json:"user_name"`
	Password string `json:"password"`
	Email    string `json:"email"` //валидация
	Phone    string `json:"phone"` //валидация
}
