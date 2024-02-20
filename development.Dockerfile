FROM golang:latest

WORKDIR /app

COPY . .

RUN go mod download

CMD ["bash", "-c", "go run ."]
