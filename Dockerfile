FROM golang:1.21 AS builder

WORKDIR /app

COPY main.go .
COPY go.mod .

RUN go get github.com/gorilla/mux
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o user-service .

FROM alpine

COPY --from=builder /app/user-service /user-service

EXPOSE 8080

CMD ["/user-service"]
