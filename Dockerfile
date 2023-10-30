FROM golang:latest

WORKDIR /app

COPY . .

RUN go mod download
RUN go install -tags 'postgres' github.com/golang-migrate/migrate/v4/cmd/migrate@latest

RUN go build -o main .

EXPOSE 8080

CMD ["./main"]
