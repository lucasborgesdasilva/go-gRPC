package database

import (
	"database/sql"

	"github.com/google/uuid"
)

// Category representa uma categoria no banco de dados.
type Category struct {
	db          *sql.DB
	ID          string
	Name        string
	Description string
}

// NewCategory cria uma nova instância de Category.
func NewCategory(db *sql.DB) *Category {
	return &Category{db: db}
}

// Create insere uma nova categoria no banco de dados.
func (c *Category) Create(name string, description string) (Category, error) {
	id := uuid.New().String()
	_, err := c.db.Exec("INSERT INTO categories (id, name, description) VALUES (?, ?, ?)",
		id, name, description)

	if err != nil {
		return Category{}, err
	}

	return Category{
		ID:          id,
		Name:        name,
		Description: description,
	}, nil
}

func (c *Category) FindAll() ([]Category, error) {
	rows, err := c.db.Query("SELECT id, name, description FROM categories")
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	categories := []Category{}
	for rows.Next() {
		var category Category
		if err := rows.Scan(&category.ID, &category.Name, &category.Description); err != nil {
			return nil, err
		}
		categories = append(categories, category)
	}
	return categories, nil
}

func (c *Category) FindByCourseID(courseID string) (*Category, error) {
	row := c.db.QueryRow(`
		SELECT cat.id, cat.name, cat.description
		FROM categories cat
		JOIN courses crs ON crs.category_id = cat.id
		WHERE crs.id = ?`, courseID)

	var category Category
	if err := row.Scan(&category.ID, &category.Name, &category.Description); err != nil {
		return nil, err
	}
	return &category, nil
}
