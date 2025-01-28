package graph

import "github.com/msaorc/FinanceGraphQL/internal/database"

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	LacamentoDB      *database.Lancamento
	CategoriaDB      *database.Categoria
	TipoLancamentoDB *database.TipoLacamento
}
