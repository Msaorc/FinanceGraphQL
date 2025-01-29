package database

import (
	"database/sql"

	"github.com/google/uuid"
)

const sqlInsertNivelNecessidade = "INSERT INTO nivel_necessidade (id, descricao, cor) VALUES ($1,$2,$3)"

type NivelNecessidade struct {
	db        *sql.DB
	ID        string
	Descricao string
	Cor       string
}

func NewNeedLevel(db *sql.DB) *NivelNecessidade {
	return &NivelNecessidade{db: db}
}

func (n *NivelNecessidade) Create(descricao string, cor string) (NivelNecessidade, error) {
	id := uuid.New().String()
	_, err := n.db.Exec(sqlInsertNivelNecessidade, id, descricao, cor)
	if err != nil {
		return NivelNecessidade{}, nil
	}

	return NivelNecessidade{
		ID:        id,
		Descricao: descricao,
		Cor:       cor,
	}, nil
}
