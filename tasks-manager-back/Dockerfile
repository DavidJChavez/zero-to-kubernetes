FROM golang:1.21.1-alpine3.18

WORKDIR /app
COPY . .
RUN go get -d -v ./...

CMD ["go", "run", "./cmd/main.go"]

