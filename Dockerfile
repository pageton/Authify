FROM golang:1.23-alpine

RUN apk update && apk add gcc musl-dev

ENV CGO_ENABLED=1

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod tidy

COPY . .

RUN go build -o authfiy ./cmd/main.go

RUN go run ./db/migrations/db_migraions/setup.go

EXPOSE 3000

CMD ["./authfiy"]
