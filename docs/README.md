# Como executar o projeto

Com Docker compose

```json
docker-compose up
```

Com Makefile:

```json
make run
```

Sem Makefile:

```json
go run ./cmd/server/main.go
```

# Exemplo de requisição

```json
  http://localhost:8000/exchange/10/BRL/USD/4.50
```

## Variáveis de ambiente

`DB_HOST=localhost`
`DB_USER=root`
`DB_NAME=desafio`
`DB_PASSWORD=root`
`DB_PORT=3308`
`DB_TIME_ZONE=America%2FSao_Paulo`
`PORT=8080`
`GIN_MODE=debug`
