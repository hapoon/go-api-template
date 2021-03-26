FROM golang:1.16-alpine3.13 as builder
WORKDIR /go/src
COPY . .
RUN go mod download

ARG CGO_ENABLED=0
ARG GOOS=linux
ARG GOARCH=amd64
RUN go build -o /go/bin/api -ldflags '-s -w'

FROM scratch as runner
COPY --from=builder /go/bin/api /app/api
ENTRYPOINT ["/app/api"]