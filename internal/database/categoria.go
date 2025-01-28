package database

import (
	"database/sql"

	"github.com/google/uuid"
)

const sqlInsert = "INSERT INTO lacamentos (id, descricao) VALUES ($1,$2)"

type Categoria struct {
	db        *sql.DB
	ID        string
	Descricao string
}

func NewCategory(db *sql.DB) *Categoria {
	return &Categoria{db: db}
}

func (c *Categoria) Create(descricao string) (Categoria, error) {
	id := uuid.New().String()
	_, err := c.db.Exec(SqlInsert, id, descricao)

	if err != nil {
		return Categoria{}, err
	}

	return Categoria{ID: id, Descricao: descricao}, nil
}
