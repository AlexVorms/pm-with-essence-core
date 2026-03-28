package columnDTO

type ColumnDTO struct {
	Name    string `json:"name"`
	IsFinal bool   `json:"is_final"`
	Order   int    `json:"order"` //порядок начинается с 0
}

type ColumnUpdateDTO struct {
	Name    string `json:"name"`
	IsFinal bool   `json:"is_final"`
}
