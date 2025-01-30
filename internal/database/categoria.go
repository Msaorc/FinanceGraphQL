package database

import (
	"database/sql"

	"github.com/google/uuid"
)

const sqlInsertCategoria = "INSERT INTO categorias (id, descricao) VALUES ($1,$2)"
const SqlBuscarCategorias = "SELECT id, descricao FROM categorias"

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
	_, err := c.db.Exec(sqlInsertCategoria, id, descricao)

	if err != nil {
		return Categoria{}, err
	}

	return Categoria{ID: id, Descricao: descricao}, nil
}

func (c *Categoria) FindAll() ([]Categoria, error) {
	rows, err := c.db.Query(SqlBuscarCategorias)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	categorias := []Categoria{}
	for rows.Next() {
		var id, descricao string
		if err := rows.Scan(&id, &descricao); err != nil {
			return nil, err
		}
		categorias = append(categorias, Categoria{ID: id, Descricao: descricao})
	}
	return categorias, nil
}
