FROM golang:1.22

WORKDIR /app

COPY . .

RUN go get -d -v ./...
RUN go build -o api ./cmd/api

EXPOSE 8000

CMD ["./api"]