package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.63

import (
	"context"
	"fmt"

	"github.com/msaorc/FinanceGraphQL/graph/model"
)

// Lancamento is the resolver for the lancamento field.
func (r *categoriaResolver) Lancamento(ctx context.Context, obj *model.Categoria) ([]*model.Lancamento, error) {
	lancamentos, err := r.LancamentoDB.FindByCategoriaID(obj.ID)
	if err != nil {
		return nil, err
	}
	var lacamentosModel []*model.Lancamento
	for _, lancamento := range lancamentos {
		lacamentosModel = append(lacamentosModel, &model.Lancamento{
			ID:          lancamento.ID,
			Descricao:   lancamento.Descricao,
			Valor:       lancamento.Valor,
			Observacao:  &lancamento.Observacao,
			Recorrencia: &lancamento.Recorrencia,
		})
	}
	return lacamentosModel, nil
}

// Lancamento is the resolver for the lancamento field.
func (r *formaPagamentoResolver) Lancamento(ctx context.Context, obj *model.FormaPagamento) ([]*model.Lancamento, error) {
	panic(fmt.Errorf("not implemented: Lancamento - lancamento"))
}

// CriarLancamento is the resolver for the criarLancamento field.
func (r *mutationResolver) CriarLancamento(ctx context.Context, input model.NovoLancamento) (*model.Lancamento, error) {
	lancamento, err := r.LancamentoDB.Create(input.Descricao, input.Valor, *input.Observacao, *input.Recorrencia, input.TipoID, input.CategoriaID, input.FormaPagamentoID, input.NecessidadeID)
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

// CriarTipoLancamento is the resolver for the criarTipoLancamento field.
func (r *mutationResolver) CriarTipoLancamento(ctx context.Context, input model.NovoTipoLancamento) (*model.TipoLancamento, error) {
	tipoLancamento, err := r.TipoLancamentoDB.Create(input.Descricao)
	if err != nil {
		return nil, err
	}

	return &model.TipoLancamento{
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

// Lancamento is the resolver for the lancamento field.
func (r *nivelNecessidadeResolver) Lancamento(ctx context.Context, obj *model.NivelNecessidade) ([]*model.Lancamento, error) {
	panic(fmt.Errorf("not implemented: Lancamento - lancamento"))
}

// Lacamentos is the resolver for the lacamentos field.
func (r *queryResolver) Lancamentos(ctx context.Context) ([]*model.Lancamento, error) {
	lancamentos, err := r.LancamentoDB.FindAll()
	if err != nil {
		return nil, err
	}
	var lacamentosModel []*model.Lancamento
	for _, lancamento := range lancamentos {
		lacamentosModel = append(lacamentosModel, &model.Lancamento{
			ID:          lancamento.ID,
			Descricao:   lancamento.Descricao,
			Valor:       lancamento.Valor,
			Observacao:  &lancamento.Observacao,
			Recorrencia: &lancamento.Recorrencia,
		})
	}
	return lacamentosModel, nil
}

// Categorias is the resolver for the categorias field.
func (r *queryResolver) Categorias(ctx context.Context) ([]*model.Categoria, error) {
	categorias, err := r.CategoriaDB.FindAll()
	if err != nil {
		return nil, err
	}
	var categoriasModel []*model.Categoria
	for _, categoria := range categorias {
		categoriasModel = append(categoriasModel, &model.Categoria{
			ID:        categoria.ID,
			Descricao: categoria.Descricao,
		})
	}
	return categoriasModel, nil
}

// TiposLancamento is the resolver for the tiposLancamento field.
func (r *queryResolver) TiposLancamento(ctx context.Context) ([]*model.TipoLancamento, error) {
	tiposLancamento, err := r.TipoLancamentoDB.FindAll()
	if err != nil {
		return nil, err
	}
	var tiposLancamentoModel []*model.TipoLancamento
	for _, tipoLancamento := range tiposLancamento {
		tiposLancamentoModel = append(tiposLancamentoModel, &model.TipoLancamento{
			ID:        tipoLancamento.ID,
			Descricao: tipoLancamento.Descricao,
		})
	}
	return tiposLancamentoModel, nil
}

// FormasPagamento is the resolver for the formasPagamento field.
func (r *queryResolver) FormasPagamento(ctx context.Context) ([]*model.FormaPagamento, error) {
	formasPagamento, err := r.FormaPagamentoDB.FindAll()
	if err != nil {
		return nil, err
	}
	var formasPagamentoModel []*model.FormaPagamento
	for _, formaPagamento := range formasPagamento {
		formasPagamentoModel = append(formasPagamentoModel, &model.FormaPagamento{
			ID:        formaPagamento.ID,
			Descricao: formaPagamento.Descricao,
		})
	}
	return formasPagamentoModel, nil
}

// NiveisNecessidade is the resolver for the niveisNecessidade field.
func (r *queryResolver) NiveisNecessidade(ctx context.Context) ([]*model.NivelNecessidade, error) {
	niveisNecessidade, err := r.NivelNecessidadeDB.FindAll()
	if err != nil {
		return nil, err
	}
	var niveisNecessidadeModel []*model.NivelNecessidade
	for _, nivelNecessidade := range niveisNecessidade {
		niveisNecessidadeModel = append(niveisNecessidadeModel, &model.NivelNecessidade{
			ID:        nivelNecessidade.ID,
			Descricao: nivelNecessidade.Descricao,
		})
	}
	return niveisNecessidadeModel, nil
}

// Lancamento is the resolver for the lancamento field.
func (r *tipoLancamentoResolver) Lancamento(ctx context.Context, obj *model.TipoLancamento) ([]*model.Lancamento, error) {
	panic(fmt.Errorf("not implemented: Lancamento - lancamento"))
}

// Categoria returns CategoriaResolver implementation.
func (r *Resolver) Categoria() CategoriaResolver { return &categoriaResolver{r} }

// FormaPagamento returns FormaPagamentoResolver implementation.
func (r *Resolver) FormaPagamento() FormaPagamentoResolver { return &formaPagamentoResolver{r} }

// Mutation returns MutationResolver implementation.
func (r *Resolver) Mutation() MutationResolver { return &mutationResolver{r} }

// NivelNecessidade returns NivelNecessidadeResolver implementation.
func (r *Resolver) NivelNecessidade() NivelNecessidadeResolver { return &nivelNecessidadeResolver{r} }

// Query returns QueryResolver implementation.
func (r *Resolver) Query() QueryResolver { return &queryResolver{r} }

// TipoLancamento returns TipoLancamentoResolver implementation.
func (r *Resolver) TipoLancamento() TipoLancamentoResolver { return &tipoLancamentoResolver{r} }

type categoriaResolver struct{ *Resolver }
type formaPagamentoResolver struct{ *Resolver }
type mutationResolver struct{ *Resolver }
type nivelNecessidadeResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
type tipoLancamentoResolver struct{ *Resolver }
