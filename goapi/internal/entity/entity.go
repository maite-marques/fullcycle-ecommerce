package entity

import "github.com/google/uuid"

// obrigatório que a 1a linha seja o nome do pacote
// vamos add as classes = structs

type Category struct {
	ID		string `json:"id"`
	Name	string `json:"name"`
}

func NewCategory(name string) *Category {
	return &Category{
		ID: 	uuid.New().String(),
		Name: name,
	}
}

type Product struct {
	ID	string `json:"id"`
	Name	string `json:"name"`
	Description	string `json:"description"`
	Price	float64 `json:"price"`
	CategoryID	string `json:"category_id"`
	ImageURL	string `json:"image_url"`
}

func NewProduct(name string, description string, categoryID, imageURL string, price float64) *Product {
	return &Product {
		ID: uuid.New().String(),
		Name: name,
		Description: description,
		Price: price,
		CategoryID: categoryID,
		ImageURL: imageURL,
	}
}