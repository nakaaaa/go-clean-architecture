FROM golang:1.18.3-alpine

WORKDIR /go/src

RUN apk upgrade --update && \
  apk --no-cache add git gcc musl-dev

COPY ./go.mod .
COPY ./go.sum .
RUN go mod download

COPY ./cmd/server/main.go .
COPY ./internal ./internal/
COPY ./pkg ./pkg/
RUN GOOS=linux GOARCH=amd64 go build -ldflags "-w -s" -o /go/bin/go-server

EXPOSE 4000

CMD ["/go/bin/go-server"]