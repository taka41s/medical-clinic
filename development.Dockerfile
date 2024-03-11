FROM golang:1.21.6

ARG GOOSE_DRIVER
ARG GOOSE_DBSTRING
ARG GOOSE_MIGRATION_DIR

ENV GOOSE_DRIVER=postgres \
    GOOSE_DBSTRING="user=postgres password=postgres dbname=medical_clinic host=db port=5432 sslmode=disable" \
    GOOSE_MIGRATION_DIR="./config/migrations"

WORKDIR /app

COPY . .

RUN go mod tidy
RUN go mod download

RUN go install github.com/pressly/goose/v3/cmd/goose@latest
RUN go install github.com/volatiletech/sqlboiler/v4@latest
RUN go install github.com/volatiletech/sqlboiler/v4/drivers/sqlboiler-psql@latest
RUN go get -u -t github.com/volatiletech/sqlboiler

CMD ["bash", "-c", "go run ."]
