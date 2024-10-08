FROM golang:alpine AS builder

RUN apk add --no-cache git

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN go build -o /go/bin/api

FROM alpine:latest

COPY --from=builder /go/bin/api /go/bin/api

EXPOSE 3030

CMD ["/go/bin/api"]
