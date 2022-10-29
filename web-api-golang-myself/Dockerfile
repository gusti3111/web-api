FROM golang:alpine AS builder

RUN apk update && apk add --no-cache git

WORKDIR /app
COPY . .
COPY go.mod /app/ 
COPY go.sum /app/
RUN go mod download

RUN go build -o web



CMD ["./web"]
