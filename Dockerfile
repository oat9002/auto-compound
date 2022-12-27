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

FROM busybox:glibc AS busybox

# Runner
FROM gcr.io/distroless/base-debian11:latest

WORKDIR /app

COPY --from=builder /auto-compound ./auto-compound
COPY --from=builder /app/healthcheck.sh ./healthcheck.sh

COPY --from=busybox /bin/sh /bin/sh
COPY --from=busybox /bin/cat /bin/cat
COPY --from=busybox /bin/grep /bin/grep

CMD ["/app/auto-compound"]

