package userDTO

type RegisterDTO struct {
	UserName string `json:"user_name" binding:"required,min=3,max=32"`
	Password string `json:"password" binding:"required,min=8,max=72"`
	Email    string `json:"email" binding:"required,email"` //валидация
}
