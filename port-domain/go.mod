module github.com/artzor/tech-test/port-domain

go 1.14

require (
	github.com/artzor/tech-test/port-domain/api v0.0.1
	github.com/stretchr/testify v1.3.0
	go.mongodb.org/mongo-driver v1.3.4
	google.golang.org/grpc v1.32.0
)

replace (
	github.com/artzor/tech-test/port-domain/api => ./api
)
