# Stage 1
FROM golang:1.24-bookworm AS builder

ENV GOOSE_VERSION=3.24.2

WORKDIR /go/src/app

RUN --mount=type=cache,target=/var/cache/apt/archives/ \
    --mount=type=cache,target=/var/cache/wget/ \
    rm -v /etc/apt/apt.conf.d/docker-clean && \
    apt update && \
    apt install --yes --no-install-recommends wget && \
    wget -c -O /var/cache/wget/goose-${GOOSE_VERSION} https://github.com/pressly/goose/releases/download/v${GOOSE_VERSION}/goose_linux_x86_64 && \
    cp /var/cache/wget/goose-${GOOSE_VERSION} ./goose && \
    chmod +x ./goose && \
    apt purge --yes wget  && \
    apt autoremove --purge --yes && \
    rm -rf /var/lib/apt/lists/*

COPY src/go.* ./

RUN --mount=type=cache,target=/go/pkg/mod/cache/download/ go mod download -x

COPY src ./

RUN --mount=type=cache,target=/go/pkg/mod/cache/download/ go build main.go

# Stage 2
FROM debian:bookworm-slim

WORKDIR /app

RUN --mount=type=cache,target=/var/cache/apt/archives/ \
    rm -v /etc/apt/apt.conf.d/docker-clean && \
    apt update && \
    apt install --yes --no-install-recommends ca-certificates postgresql-client && \
    rm -rf /var/lib/apt/lists/*

COPY --from=builder /go/src/app/main .
COPY --from=builder /go/src/app/goose .
COPY src/config.json .
COPY migrations migrations
COPY seeds seeds
COPY entrypoint.sh .

EXPOSE 50051

ENTRYPOINT ["bash", "/app/entrypoint.sh"]