FROM golang:alpine

RUN apk add --no-cache git bash yq curl sed

WORKDIR /app

# COPY go.mod go.sum ./

COPY . .

RUN go mod download

RUN go build -o /go/bin/api

SHELL ["/bin/bash", "-c"]
RUN ./seed/characters/seed-characters.sh

EXPOSE 3030

ENTRYPOINT ["/go/bin/api", "--seed"]
