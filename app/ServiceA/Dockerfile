FROM golang:1.22-alpine

WORKDIR /app

COPY . .

RUN go mod tidy

RUN go build -o serviceA .

EXPOSE 50051

CMD ["./serviceA"]