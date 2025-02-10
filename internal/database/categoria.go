package database

import (
	"database/sql"

	"github.com/google/uuid"
)

const sqlInsertCategoria = "INSERT INTO categorias (id, descricao) VALUES ($1,$2)"
const sqlFindCategorias = "SELECT id, descricao FROM categorias"
const sqlFindByLacamentoID = "SELECT C.id, C.descricao FROM categorias C JOIN lancamentos L ON (C.id = L.categoria_id) WHERE L.id = $1"

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
	rows, err := c.db.Query(sqlFindCategorias)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	categorias, err := processaRetornoCategoria(rows)
	if err != nil {
		return nil, err
	}
	return categorias, nil
}

func (c *Categoria) FindCategoriaByLancamentoID(lancamentoID string) (Categoria, error) {
	var id, descricao string
	err := c.db.QueryRow(sqlFindByLacamentoID, lancamentoID).Scan(&id, &descricao)
	if err != nil {
		return Categoria{}, err
	}
	return Categoria{ID: id, Descricao: descricao}, nil
}

func processaRetornoCategoria(rows *sql.Rows) ([]Categoria, error) {
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
