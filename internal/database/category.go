package database

import (
	"database/sql"
	"github.com/google/uuid"
)

type Category struct {
	db          *sql.DB //Estamos obtendo as interfaces para trabalharmos com banco de dados.
	ID          string
	Name        string
	Description string
}

// Esse é o método construtor
func NewCategory(db *sql.DB) *Category {

	return &Category{db: db}
}

func (c *Category) Create(name string, description string) (Category, error) {

	id := uuid.New().String()
	_, err := c.db.Exec("INSERT INTO categories (id, name, description) VALUES ($1, $2, $3)",
		id, name, description)

	if err != nil {
		return Category{}, err
	}

	return Category{ID: id, Name: name, Description: description}, nil
}
