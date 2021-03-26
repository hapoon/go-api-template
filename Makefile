SERVER=main.go
IMAGE=hapoon/go-api-template

.PHONY: run test docker-build
run:
	go run ${SERVER}

test:
	go test -cover ./...

docker-build:
	docker build -t ${IMAGE} .

docker-run:
	docker run -d -p 8080:8080 ${IMAGE}