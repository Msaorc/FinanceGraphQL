package database

import (
	"database/sql"

	"github.com/google/uuid"
)

const SqlInsertLancamento = "INSERT INTO lacamentos (id, descricao, valor, observacao, recorrencia) VALUES ($1,$2,$3,$4,$5)"

type Lancamento struct {
	db          *sql.DB
	ID          string
	Descricao   string
	Valor       float64
	Observacao  string
	Recorrencia int32
}

func NewLaunch(db *sql.DB) *Lancamento {
	return &Lancamento{db: db}
}

func (c *Lancamento) Create(descricao string, valor float64, observacao string, recorrencia int32) (Lancamento, error) {
	id := uuid.New().String()
	_, err := c.db.Exec(SqlInsertLancamento, id, descricao, valor, observacao, recorrencia)

	if err != nil {
		return Lancamento{}, err
	}

	return Lancamento{ID: id, Descricao: descricao, Valor: valor, Observacao: observacao, Recorrencia: recorrencia}, nil
}
