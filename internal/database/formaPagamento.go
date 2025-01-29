package database

import (
	"database/sql"

	"github.com/google/uuid"
)

const sqlInsertFormaPagamento = "INSERT INTO nivel_necessidade (id, descricao) VALUES ($1,$2)"

type FormaPagamento struct {
	db        *sql.DB
	ID        string
	Descricao string
}

func NewPaymentMethod(db *sql.DB) *NivelNecessidade {
	return &NivelNecessidade{db: db}
}

func (f *FormaPagamento) Create(descricao string) (FormaPagamento, error) {
	id := uuid.New().String()
	_, err := f.db.Exec(sqlInsertFormaPagamento, id, descricao)

	if err != nil {
		return FormaPagamento{}, nil
	}

	return FormaPagamento{
		ID:        id,
		Descricao: descricao,
	}, nil
}
