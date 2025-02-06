FROM golang:1.23.5-alpine

WORKDIR /app

COPY . .
RUN go mod download

EXPOSE 8080

CMD ["go", "run", "cmd/main.go"]