package models

type Snack struct {
	ID             int     `json:"id"`
	Name           string  `json:"name"`
	Category       string  `json:"category"`
	Price          float64 `json:"price"`
	ImageURL       string  `json:"image_url"`
	HealthBenefits string  `json:"health_benefits"`
	Description    string  `json:"description"`
	Discount       float64 `json:"discount"`
	Ingredients    string  `json:"ingredients"`
}
