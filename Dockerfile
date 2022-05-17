FROM golang:1.18-buster as builder

RUN apt-get update && \
    apt-get install -y git ca-certificates tzdata && \
    update-ca-certificates

WORKDIR /app

ENV GO111MODULE=on
COPY go.mod go.sum ./
RUN export GOPROXY=https://proxy.golang.org && go mod download -x
RUN go mod verify

COPY . .

RUN CGO_ENABLED=0 \
    GOOS=linux GOARCH=amd64 \
    go build -ldflags="-w -s" \
    -o vidlearn \
    ./main.go

FROM debian:buster-20200908-slim

WORKDIR /app
ENV USER=appuser
ENV UID=10001
ENV TZ=Asia/Makassar

COPY ./config/config.toml /app/config.toml

RUN adduser \    
    --disabled-password \    
    --gecos "" \    
    --home "/nonexistent" \    
    --shell "/sbin/nologin" \    
    --no-create-home \    
    --uid "${UID}" \    
    "${USER}"

COPY --from=builder --chown=appuser:appuser /app/vidlearn .
COPY --from=builder --chown=appuser:appuser /app/config ./config
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
RUN apt-get update; \
    apt-get upgrade; \
    cp /usr/share/zoneinfo/${TZ} /etc/localtime; \
    date;

RUN echo Y || apt-get install curl

STOPSIGNAL SIGINT

EXPOSE 8000

ENTRYPOINT ["./vidlearn"]