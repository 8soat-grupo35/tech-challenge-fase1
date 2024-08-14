# CHALLENGE SOAT8 - GRUPO 35

## Conteúdo

- [Sobre](#sobre)
- [DDD - Event Storming](#ddd---event-storming)
- [Como rodar a aplicação](#como-rodar-a-aplicação)
- [Contribuidores](#contribuidores)

## Sobre

Neste primeira fase desenvolvemos uma projeto para o curso de Pós Graduação em Software Architecture da FIAP com os requisitos solicitados no [Challenge](https://on.fiap.com.br/mod/conteudoshtml/view.php?id=407435&c=11255). Foi desenvolvido pelo **Grupo 35** da **SOAT8**.


## DDD - Event Storming

Disponibilizamos através deste [link](https://miro.com/app/board/uXjVK4xDf-w=/?share_link_id=428293472540) o Event Storming realizado pela equipe com todas etapas realizadas.

## Como rodar a aplicação

Para rodar a aplicação, deve-se executar o comando:

```
docker-compose up
```

Para visualizar o **Swagger**, devemos manter a aplicação rodando e acessar a URL abaixo:

`http://localhost:8000/swagger/index.html`

<!-- 
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
``` -->

## Contribuidores

- Egio Lima
- Gabriel Rosa Nunes
- Jhony Eduardo Senem