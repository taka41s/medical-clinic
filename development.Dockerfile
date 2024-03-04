FROM golang:1.21.6

ARG GOOSE_DRIVER
ARG GOOSE_DBSTRING
ARG GOOSE_MIGRATION_DIR

ENV GOOSE_DRIVER=postgres \
    GOOSE_DBSTRING="user=postgres password=postgres dbname=medical_clinic host=db port=5432 sslmode=disable" \
    GOOSE_MIGRATION_DIR="./config/migrations"

WORKDIR /app

COPY . .

RUN go mod download

RUN go install github.com/pressly/goose/v3/cmd/goose@latest

CMD ["bash", "-c", "go run ."]
