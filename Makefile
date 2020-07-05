all: start

start:
	docker-compose up --build

test:
	cd port-domain && go test ./... -v
	cd client-api && go test ./... -v

gen-proto:
	protoc -I port-domain/service port-domain/service/service.proto --go_out=plugins=grpc:port-domain/service
	protoc -I port-domain/service port-domain/service/service.proto --go_out=plugins=grpc:client-api/portdomain/service
