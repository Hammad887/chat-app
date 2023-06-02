FROM golang:1.16-alpine

WORKDIR /app

COPY . .

RUN go mod download

RUN go build -o chat-app

EXPOSE 8000

CMD ["./chat-app"]
