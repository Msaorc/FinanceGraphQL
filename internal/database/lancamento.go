package database

import (
	"database/sql"

	"github.com/google/uuid"
)

const SqlInsertLancamento = "INSERT INTO lacamentos (id, descricao, valor, observacao, recorrencia, tipo_lancamento_id, categoria_id, forma_pagamento_id, nivel_necessidade_id) VALUES ($1,$2,$3,$4,$5,$6,$7,$8,$9)"

type Lancamento struct {
	db                 *sql.DB
	ID                 string
	Descricao          string
	Valor              float64
	Observacao         string
	Recorrencia        int32
	TipoLancamentoID   string
	CategoriaID        string
	FormaPagamentoID   string
	NivelNecessidadeID string
}

func NewLaunch(db *sql.DB) *Lancamento {
	return &Lancamento{db: db}
}

func (c *Lancamento) Create(descricao string, valor float64, observacao string, recorrencia int32, tipoLancamentoID string, categoriaID string, formaPagamentoID string, nivelNecessidadeID string) (Lancamento, error) {
	id := uuid.New().String()
	_, err := c.db.Exec(SqlInsertLancamento, id, descricao, valor, observacao, recorrencia, tipoLancamentoID, categoriaID, formaPagamentoID, nivelNecessidadeID)

	if err != nil {
		return Lancamento{}, err
	}

	return Lancamento{ID: id, Descricao: descricao, Valor: valor, Observacao: observacao, Recorrencia: recorrencia, TipoLancamentoID: tipoLancamentoID,
		CategoriaID: categoriaID, FormaPagamentoID: formaPagamentoID, NivelNecessidadeID: nivelNecessidadeID}, nil
}
