SERVER=main.go

.PHONY: run
run:
	go run ${SERVER}

.PHONY: test
test:
	go test -cover ./...
