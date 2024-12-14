FROM golang:1.23

WORKDIR /app
COPY . /app

RUN go install github.com/air-verse/air@latest
RUN go install github.com/swaggo/swag/cmd/swag@latest

COPY go.mod go.sum ./
RUN go mod download

EXPOSE 8000

CMD ["air", "-c", ".air.toml"]