# protocolbuffers-go-example
Exemplo simples de Protocol Buffers em GOLANG

# PASSOS

Os comandos a seguir devem ser exeutados a partir da base do workspace

## Baixar e instalar pacote golang

go get -u github.com/golang/protobuf/proto
c-gen-go

## Baixar compilador protoc e descompactar

wget https://github.com/protocolbuffers/protobuf/releases/download/v3.6.1/protoc-3.6.1-linux-x86_64.zip
unzip protoc-3.6.1-linux-x86_64.zip

## Uso do compilador

protoc --go_out=. *.proto

(O compilador protoc fica localizado na pasta bin na base do workspace)

## Execução do exemplo

go run *.go




