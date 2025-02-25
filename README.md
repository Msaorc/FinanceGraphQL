# FinanceGraphQL

## Url para criação da estrutura GraphQL:

* https://gqlgen.com/getting-started/

## Técnologias utilizadas:

* **Banco de dados:** SQlite
* **Linguagem:** Golang

## Execução do servidor GraphQL:

- `go run server.go`

* Para fazer a consulta de lançamentos com seus relacionamentos:

```
  query Consultarlancamentos{
  lancamentos {
    id
    descricao
    valor
    observacao
  	recorrencia
  	categoria {
      id
      descricao
  	} 
    tipo {
      id
      descricao
    } 
    necessidade {
      id
      descricao
    } 
    formaPagamento {
      id
      descricao
    }
  }
}
```

## Links importantes:

-- https://github.com/mattn/go-sqlite3

-- https://ultahost.com/knowledge-base/install-sqlite-on-ubuntu/

## Execução via docker:

-- docker build -t "nome_imagem" .

-- docker run -dit --name "nome_container" -p8080:8080 "nome_imagem"

## Aplicação

-- http://localhost:8080/