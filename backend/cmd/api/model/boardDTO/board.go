package boardDTO

type BoardDTO struct {
	Name        string `json:"name" gorm:"not null"`
	Description string `json:"description"`
	IsPublic    bool   `json:"isPublic"`
}
