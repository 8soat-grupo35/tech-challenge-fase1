# Rodar a aplicação 

1. Subir o container de banco de dados:

```
docker-compose up
```

2. Rodar a aplicação:

```
go run cmd/api main.go
```

> Posteriomente colocaremos este passo no Docker.


3. Consumir endpoint


```
http://localhost:8000/items
```

# Rodar os testes

1. Instalação do gomock para execução dos testes localmente

```
go install go.uber.org/mock/mockgen@latest
```

2. Geração dos arquivos de mock preenchidos via go generate no projeto

```
go generate ./...
```

3. Execução dos testes do projeto

```
go test ./test/...
```