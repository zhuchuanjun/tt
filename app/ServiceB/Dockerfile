FROM golang:1.22-alpine

WORKDIR /app

COPY . .

RUN go mod tidy

RUN go build -o serviceB .

EXPOSE 50052

CMD ["./serviceB"]