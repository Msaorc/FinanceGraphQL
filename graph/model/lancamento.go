package model

type Lancamento struct {
	ID          string  `json:"id"`
	Descricao   string  `json:"descricao"`
	Valor       float64 `json:"valor"`
	Observacao  *string `json:"observacao,omitempty"`
	Recorrencia *int32  `json:"recorrencia,omitempty"`
}
