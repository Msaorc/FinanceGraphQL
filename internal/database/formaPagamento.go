package database

import (
	"database/sql"

	"github.com/google/uuid"
)

const sqlInsert = "INSERT INTO forma_pagamento (id, descricao) VALUES ($1,$2)"

type FormaPagamento struct {
	db        *sql.DB
	ID        string
	Descricao string
}

func (f *FormaPagamento) Create(descricao string) (FormaPagamento, error) {
	id := uuid.New().String()
	_, err := f.db.Exec(sqlInsert, id, descricao)

	if err != nil {
		return FormaPagamento{}, nil
	}

	return FormaPagamento{
		ID:        id,
		Descricao: descricao,
	}, nil
}
