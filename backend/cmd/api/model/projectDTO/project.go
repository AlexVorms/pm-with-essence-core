package projectDTO

type ProjectDTO struct {
	Name        string `json:"name"`
	IsPublic    bool   `json:"isPublic"`
	Description string `json:"description"`
}
