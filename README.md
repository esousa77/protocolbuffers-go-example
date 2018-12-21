# protocolbuffers-go-example
Exemplo simples de Protocol Buffers em GOLANG

# PASSOS

## Baixar e instalar pacote golang

go get -u github.com/golang/protobuf/proto
c-gen-go

## Baixar compilador protoc e descompactar

wget https://github.com/protocolbuffers/protobuf/releases/download/v3.6.1/protoc-3.6.1-linux-x86_64.zip
unzip protoc-3.6.1-linux-x86_64.zip

## Uso do compilador

protoc --go_out=. *.proto
