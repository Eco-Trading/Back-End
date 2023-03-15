# api

A pasta api, é onde deixamos os arquivos openapi yml

# cmd

A pasta cmd, é onde deixamos os arquivos principais da lambda aws

# config

A pasta config, é onde deixamos os arquivos de configuraçao

# internal

A pasta internal, é onde deixamos o core da aplicação e não queremos expor para outras aplicações

# test

A pasta test, é onde colocamos os teste E2E, arquivo mocados e arquivo para testes de eventos

# util

A pasta util, é onde fica struturas independentes que serão utilizadas para auxiliar

# Commands

## Invoke locality

```shell
GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -o main main.go
sam local invoke "HelloWorldFunction"
```

## Invoke locality once line

```shell
cd hello-world/ && GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -o main main.go && cd .. && sam local invoke "HelloWorldFunction"
```

## Deploy

```shell
sam deploy --guided
```

```shell
sam build
sam deploy --no-confirm-changeset --no-fail-on-empty-changeset
```
# Docker 
## criar uma rede local no docker

```shell
docker network create eco_tradind_net
docker compose -f docker-compose-dev.yml up -d
```

# Test
## validar cobertura de teste 
```shell
go test -cover -coverprofile=coverage.out ./... && go tool cover -html=coverage.out -o coverage.html

```

# Command Proto3

```shell
go get google.golang.org/protobuf/cmd/protoc-gen-go
protoc --proto_path=proto proto/*.proto --go_out=plugins=grpc:./internal/infra/proto_buffer
```