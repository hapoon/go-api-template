SERVER=main.go
IMAGE=hapoon/go-api-template

.PHONY: run test docker-build mockgen
run:
	go run ${SERVER}

test:
	go test -cover ./...

mockgen:
	mockgen -source=repository/sample.go -destination=repository/mock/sample.go

docker-build:
	docker build -t ${IMAGE} .

docker-run:
	docker run -d -p 8080:8080 ${IMAGE}