Google Remote Procedure Client

Tipos:
Unary: Cliente manda uma request e server retorna uma response
Server Streaming: Cliente manda uma request e server retorna varias responses
Client Streaming: Cliente manda varias requests e server retorna uma response
Bi directional Streaming: Cliente manda várias requests e server retorna várias responses


go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.28
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.2

protoc --proto_path=proto --go_out=pb --go-grpc_out=. --go_opt=paths=source_relative proto/*.proto
go get google.golang.org/grpc 



Após implementar basta ir até o insomnia, importasr o arquivo e enviar a requisição




https://www.youtube.com/watch?v=UVdY-EK3c3o
https://www.youtube.com/watch?v=ES_GI-lmhEU