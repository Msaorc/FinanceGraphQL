package database

import (
	"database/sql"

	"github.com/google/uuid"
)

type TipoLacamento struct {
	db        *sql.DB
	ID        string
	Descricao string
}

const sqlInsert = "INSERT INTO tipo_lancamento (id, descricao) VALUES ($1,$2)"

func NewReleaseType(db *sql.DB) *TipoLacamento {
	return &TipoLacamento{db: db}
}

func (t *TipoLacamento) Create(descricao string) (TipoLacamento, error) {
	id := uuid.New().String()
	_, err := t.db.Exec(sqlInsert, id, descricao)

	if err != nil {
		return TipoLacamento{}, err
	}

	return TipoLacamento{ID: id, Descricao: descricao}, nil
}
