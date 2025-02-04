package database

import (
	"database/sql"

	"github.com/google/uuid"
)

const sqlInsertFormaPagamento = "INSERT INTO forma_pagamento (id, descricao) VALUES ($1,$2)"
const sqlFindFormaPagamento = "SELECT id, descricao FROM forma_pagamento"

type FormaPagamento struct {
	db        *sql.DB
	ID        string
	Descricao string
}

func NewPaymentMethod(db *sql.DB) *FormaPagamento {
	return &FormaPagamento{db: db}
}

func (f *FormaPagamento) Create(descricao string) (FormaPagamento, error) {
	id := uuid.New().String()
	_, err := f.db.Exec(sqlInsertFormaPagamento, id, descricao)

	if err != nil {
		return FormaPagamento{}, err
	}

	return FormaPagamento{
		ID:        id,
		Descricao: descricao,
	}, nil
}

func (f *FormaPagamento) FindAll() ([]FormaPagamento, error) {
	rows, err := f.db.Query(sqlFindFormaPagamento)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	formasPagamento := []FormaPagamento{}
	for rows.Next() {
		var id, descricao string
		if err := rows.Scan(&id, &descricao); err != nil {
			return nil, err
		}
		formasPagamento = append(formasPagamento, FormaPagamento{ID: id, Descricao: descricao})
	}
	return formasPagamento, nil
}
