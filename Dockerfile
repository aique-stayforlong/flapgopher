FROM golang:1.15-buster AS builder

RUN apt-get update && apt-get install -y \
    gcc \
    pkg-config \
    libsdl2-dev \
    libsdl2-image-dev \
    libsdl2-ttf-dev \
    && rm -rf /var/lib/apt/lists/*

WORKDIR /app
COPY app/go.mod app/go.sum ./
RUN go mod download

COPY app/. .
RUN go build .

FROM debian:stable-slim

RUN apt-get update && apt-get install -y \
    libsdl2-2.0.0 \
    libsdl2-image-2.0.0 \
    libsdl2-ttf-2.0.0 \
    && rm -rf /var/lib/apt/lists/*

WORKDIR /app/app
COPY --from=builder /app/flapgopher.asuarez.net .

CMD ["./flapgopher.asuarez.net"]