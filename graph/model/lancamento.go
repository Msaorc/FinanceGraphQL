package model

type Lancamento struct {
	ID             string            `json:"id"`
	Descricao      string            `json:"descricao"`
	Tipo           *TipoLancamento   `json:"tipo"`
	Categoria      *Categoria        `json:"categoria"`
	Valor          float64           `json:"valor"`
	FormaPagamento *FormaPagamento   `json:"formaPagamento"`
	Observacao     *string           `json:"observacao,omitempty"`
	Necessidade    *NivelNecessidade `json:"necessidade"`
	Recorrencia    *int32            `json:"recorrencia,omitempty"`
}
