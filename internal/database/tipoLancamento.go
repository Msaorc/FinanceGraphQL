package database

import (
	"database/sql"

	"github.com/google/uuid"
)

const sqlInsertTipoLancamento = "INSERT INTO tipo_lancamento (id, descricao) VALUES ($1,$2)"
const sqlFindTipoLancamento = "SELECT id, descricao FROM tipo_lancamento"

type TipoLancamento struct {
	db        *sql.DB
	ID        string
	Descricao string
}

func NewReleaseType(db *sql.DB) *TipoLancamento {
	return &TipoLancamento{db: db}
}

func (t *TipoLancamento) Create(descricao string) (TipoLancamento, error) {
	id := uuid.New().String()
	_, err := t.db.Exec(sqlInsertTipoLancamento, id, descricao)

	if err != nil {
		return TipoLancamento{}, err
	}

	return TipoLancamento{ID: id, Descricao: descricao}, nil
}

func (c *TipoLancamento) FindAll() ([]TipoLancamento, error) {
	rows, err := c.db.Query(sqlFindTipoLancamento)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	tiposLancamento := []TipoLancamento{}
	for rows.Next() {
		var id, descricao string
		if err := rows.Scan(&id, &descricao); err != nil {
			return nil, err
		}
		tiposLancamento = append(tiposLancamento, TipoLancamento{ID: id, Descricao: descricao})
	}
	return tiposLancamento, nil
}
