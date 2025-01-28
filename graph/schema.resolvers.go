package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.63

import (
	"context"
	"fmt"

	"github.com/msaorc/FinanceGraphQL/graph/model"
)

// CriarLancamento is the resolver for the criarLancamento field.
func (r *mutationResolver) CriarLancamento(ctx context.Context, input model.NovoLancamento) (*model.Lancamento, error) {
	lancamento, err := r.LacamentoDB.Create(input.Descricao, input.Valor, *input.Observacao, *input.Recorrencia)
	if err != nil {
		return nil, err
	}

	return &model.Lancamento{
		ID:          lancamento.ID,
		Descricao:   lancamento.Descricao,
		Valor:       lancamento.Valor,
		Observacao:  &lancamento.Observacao,
		Recorrencia: &lancamento.Recorrencia,
	}, nil
}

// CriarCategoria is the resolver for the criarCategoria field.
func (r *mutationResolver) CriarCategoria(ctx context.Context, input model.NovaCategoria) (*model.Categoria, error) {
	categoria, err := r.CategoriaDB.Create(input.Descricao)
	if err != nil {
		return nil, err
	}

	return &model.Categoria{
		ID:        categoria.ID,
		Descricao: categoria.Descricao,
	}, nil
}

// CriarTipoLacamento is the resolver for the criarTipoLacamento field.
func (r *mutationResolver) CriarTipoLacamento(ctx context.Context, input model.NovoTipoLancamento) (*model.TipoLacamento, error) {
	tipoLancamento, err := r.TipoLancamentoDB.Create(input.Descricao)
	if err != nil {
		return nil, err
	}

	return &model.TipoLacamento{
		ID:        tipoLancamento.ID,
		Descricao: tipoLancamento.Descricao,
	}, nil
}

// CriarFormaPagamento is the resolver for the criarFormaPagamento field.
func (r *mutationResolver) CriarFormaPagamento(ctx context.Context, input model.NovaFormaPagamento) (*model.FormaPagamento, error) {
	formaPagamento, err := r.FormaPagamentoDB.Create(input.Descricao)
	if err != nil {
		return &model.FormaPagamento{}, err
	}

	return &model.FormaPagamento{
		ID:        formaPagamento.ID,
		Descricao: formaPagamento.Descricao,
	}, nil
}

// CriarNivelNecessidade is the resolver for the criarNivelNecessidade field.
func (r *mutationResolver) CriarNivelNecessidade(ctx context.Context, input model.NovoNivelNecessidade) (*model.NivelNecessidade, error) {
	nivelNecessidade, err := r.NivelNecessidadeDB.Create(input.Descricao, input.Cor)
	if err != nil {
		return &model.NivelNecessidade{}, nil
	}

	return &model.NivelNecessidade{
		ID:        nivelNecessidade.ID,
		Descricao: nivelNecessidade.Descricao,
		Cor:       nivelNecessidade.Cor,
	}, nil
}

// Lacamentos is the resolver for the lacamentos field.
func (r *queryResolver) Lacamentos(ctx context.Context) ([]*model.Lancamento, error) {
	panic(fmt.Errorf("not implemented: Lacamentos - lacamentos"))
}

// Categorias is the resolver for the categorias field.
func (r *queryResolver) Categorias(ctx context.Context) ([]*model.Categoria, error) {
	panic(fmt.Errorf("not implemented: Categorias - categorias"))
}

// TiposLacamento is the resolver for the tiposLacamento field.
func (r *queryResolver) TiposLacamento(ctx context.Context) ([]*model.TipoLacamento, error) {
	panic(fmt.Errorf("not implemented: TiposLacamento - tiposLacamento"))
}

// FormasPagamento is the resolver for the formasPagamento field.
func (r *queryResolver) FormasPagamento(ctx context.Context) ([]*model.FormaPagamento, error) {
	panic(fmt.Errorf("not implemented: FormasPagamento - formasPagamento"))
}

// NiveisNecessidade is the resolver for the niveisNecessidade field.
func (r *queryResolver) NiveisNecessidade(ctx context.Context) ([]*model.NivelNecessidade, error) {
	panic(fmt.Errorf("not implemented: NiveisNecessidade - niveisNecessidade"))
}

// Mutation returns MutationResolver implementation.
func (r *Resolver) Mutation() MutationResolver { return &mutationResolver{r} }

// Query returns QueryResolver implementation.
func (r *Resolver) Query() QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
