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