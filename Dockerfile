# Builder
FROM golang:1.19-bullseye AS builder

WORKDIR /app

COPY go.mod ./
COPY go.sum ./
COPY config ./config
COPY contracts ./contracts
COPY services ./services
COPY utils ./utils
COPY *.go ./
COPY healthcheck.sh ./

RUN go mod download
RUN go build -o /auto-compound

# Runner
FROM debian:bullseye-slim 

WORKDIR /app

COPY --from=builder /auto-compound ./auto-compound
COPY --from=builder /app/healthcheck.sh ./healthcheck.sh

CMD ["/app/auto-compound"]

