FROM golang:1.21.6

ARG GOOSE_DRIVER \
    GOOSE_DBSTRING \
    GOOSE_MIGRATION_DIR

ENV GOOSE_DRIVER=postgres \
    GOOSE_DBSTRING=DBSTRING \
    GOOSE_DBSTRING="user=postgres password=postgres dbname=medical_clinic host=db port=5432 sslmode=disable"

WORKDIR /app

COPY . .

RUN go mod download
RUN chmod +x entrypoint.sh

CMD ["bash", "-c", "go run ."]
