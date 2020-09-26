all: start

start:
	docker-compose up --build

test:
	cd port-domain && go test ./... -v
	cd client-api && go test ./... -v

gen-proto:
	protoc -I port-domain/api port-domain/api/api.proto --go_out=plugins=grpc:port-domain/api
