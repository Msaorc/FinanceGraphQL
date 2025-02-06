package database

import (
	"database/sql"

	"github.com/google/uuid"
)

const sqlInsertLancamento = "INSERT INTO lancamentos (id, descricao, valor, observacao, recorrencia, tipo_lancamento_id, categoria_id, forma_pagamento_id, nivel_necessidade_id) VALUES ($1,$2,$3,$4,$5,$6,$7,$8,$9)"
const sqlFindLancamento = "SELECT id, descricao, valor, observacao, recorrencia, tipo_lancamento_id, categoria_id, forma_pagamento_id, nivel_necessidade_id FROM lancamentos"
const sqlFindByCategoriaID = "SELECT id, descricao, valor, observacao, recorrencia, tipo_lancamento_id, categoria_id, forma_pagamento_id, nivel_necessidade_id FROM lancamentos WHERE categoria_id = $1"
const sqlFindByFormaPagamentoID = "SELECT id, descricao, valor, observacao, recorrencia, tipo_lancamento_id, categoria_id, forma_pagamento_id, nivel_necessidade_id FROM lancamentos WHERE forma_pagamento_id = $1"
const sqlFindByNivelNecessidadeID = "SELECT id, descricao, valor, observacao, recorrencia, tipo_lancamento_id, categoria_id, forma_pagamento_id, nivel_necessidade_id FROM lancamentos WHERE nivel_necessidade_id = $1"
const sqlFindByTipoLancamentoID = "SELECT id, descricao, valor, observacao, recorrencia, tipo_lancamento_id, categoria_id, forma_pagamento_id, nivel_necessidade_id FROM lancamentos WHERE tipo_lancamento_id = $1"

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
	_, err := c.db.Exec(sqlInsertLancamento, id, descricao, valor, observacao, recorrencia, tipoLancamentoID, categoriaID, formaPagamentoID, nivelNecessidadeID)

	if err != nil {
		return Lancamento{}, err
	}

	return Lancamento{ID: id, Descricao: descricao, Valor: valor, Observacao: observacao, Recorrencia: recorrencia, TipoLancamentoID: tipoLancamentoID,
		CategoriaID: categoriaID, FormaPagamentoID: formaPagamentoID, NivelNecessidadeID: nivelNecessidadeID}, nil
}

func (c *Lancamento) FindAll() ([]Lancamento, error) {
	rows, err := c.db.Query(sqlFindLancamento)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	lancamentos, err := processaRetornoLancamento(rows)
	if err != nil {
		return nil, err
	}
	return lancamentos, nil
}

func (c *Lancamento) FindByCategoriaID(id string) ([]Lancamento, error) {
	rows, err := c.db.Query(sqlFindByCategoriaID, id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	lancamentos, err := processaRetornoLancamento(rows)
	if err != nil {
		return nil, err
	}
	return lancamentos, nil
}

func (c *Lancamento) FindByFormaPagamentoID(id string) ([]Lancamento, error) {
	rows, err := c.db.Query(sqlFindByFormaPagamentoID, id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	lancamentos, err := processaRetornoLancamento(rows)
	if err != nil {
		return nil, err
	}
	return lancamentos, nil
}

func (c *Lancamento) FindByNivelNecessidadeID(id string) ([]Lancamento, error) {
	rows, err := c.db.Query(sqlFindByNivelNecessidadeID, id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	lancamentos, err := processaRetornoLancamento(rows)
	if err != nil {
		return nil, err
	}
	return lancamentos, nil
}

func (c *Lancamento) FindByTipoLancamentoID(id string) ([]Lancamento, error) {
	rows, err := c.db.Query(sqlFindByTipoLancamentoID, id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	lancamentos, err := processaRetornoLancamento(rows)
	if err != nil {
		return nil, err
	}
	return lancamentos, nil
}

func processaRetornoLancamento(rows *sql.Rows) ([]Lancamento, error) {
	lancamentos := []Lancamento{}
	for rows.Next() {
		var id, descricao, observacao, tipo_lacamento_id, categoria_id, forma_pagamento_id, nivel_necessidade_id string
		var valor float64
		var recorrencia int32
		if err := rows.Scan(&id, &descricao, &valor, &observacao, &recorrencia, &tipo_lacamento_id, &categoria_id, &forma_pagamento_id, &nivel_necessidade_id); err != nil {
			return nil, err
		}
		lancamentos = append(lancamentos, Lancamento{
			ID:                 id,
			Descricao:          descricao,
			Valor:              valor,
			Observacao:         observacao,
			Recorrencia:        recorrencia,
			TipoLancamentoID:   tipo_lacamento_id,
			CategoriaID:        categoria_id,
			FormaPagamentoID:   forma_pagamento_id,
			NivelNecessidadeID: nivel_necessidade_id,
		})
	}
	return lancamentos, nil
}
