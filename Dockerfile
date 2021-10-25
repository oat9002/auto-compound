# Builder
FROM golang:1.17-alpine AS builder

RUN apk add --no-cache gcc musl-dev linux-headers git

WORKDIR /app

COPY go.mod ./
COPY go.sum ./
COPY config ./config
COPY contracts ./contracts
COPY services ./services
COPY utils ./utils
COPY *.go ./

RUN go mod download
RUN go build -o /auto-compound

# Runner
FROM alpine:latest  

# RUN apk add --no-cache gcc

WORKDIR /app

COPY --from=builder /auto-compound ./auto-compound
CMD ["/app/auto-compound"]

