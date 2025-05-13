FROM golang:1.20-alpine

WORKDIR /app

COPY app.go .

RUN go build -o main .

CMD ["./main"]

