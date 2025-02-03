package database

import (
	"database/sql"

	"github.com/google/uuid"
)

const sqlInsertNivelNecessidade = "INSERT INTO nivel_necessidade (id, descricao, cor) VALUES ($1,$2,$3)"
const sqlFindNivelNecessidade = "SELECT id, descricao FROM categorias"

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

func (f *NivelNecessidade) FindAll() ([]NivelNecessidade, error) {
	rows, err := f.db.Query(sqlFindNivelNecessidade)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	niveisNecessidade := []NivelNecessidade{}
	for rows.Next() {
		var id, descricao string
		if err := rows.Scan(&id, &descricao); err != nil {
			return nil, err
		}
		niveisNecessidade = append(niveisNecessidade, NivelNecessidade{ID: id, Descricao: descricao})
	}
	return niveisNecessidade, nil
}
