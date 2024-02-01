package database

import (
	"database/sql"

	"github.com/maite-marques/fullcycle-ecommerce/goapi/internal/entity"
)

type CategoryDB struct {
	db *sql.DB // conexão com um banco de dados
}

func NewCategoryDB(db *sql.DB) *CategoryDB {
	return &CategoryDB{db: db}
}

func(cd *CategoryDB) GetCategories() ([]*entity.Category, error) {
	rows, err := cd.db.Query("SELECT id, name FROM categories")
	if err != nil {
		return nil, err
	}
	defer rows.Close() // o defer garante que essa linha só vai ser executada no final, depois de tudo executado na função.

	var categories []*entity.Category
	for rows.Next() {
		var category entity.Category
		if err := rows.Scan(&category.ID); err != nil { // O & indica que a variável vai ser alterada na memória direto
			return nil, err
		} 
		categories = append(categories, &category)
	}
	return categories, nil
}

func (cd *CategoryDB) GetCategory(id string) (*entity.Category, error) {
	var category entity.Category
	err := cd.db.QueryRow("SELECT id, name FROM categories WHERE id = ?").Scan(&category.ID, &category.Name)
	if err != nil {
		return nil, err
	}
	return &category, nil
}

func (cd *CategoryDB) CreateCategory(category *entity.Category) (string, error) {
	_, err := cd.db.Exec("INSERT INTO categories (id, name) VALUES (?, ?)", category.ID, category.Name) // as ? serão substituidas pelas variaveis id e name
	if err != nil {
		return "", err
	}
	return category.ID, nil
}
