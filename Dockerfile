FROM golang:1.21-alpine AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -o simpledns

FROM scratch

COPY --from=builder /app/simpledns /

EXPOSE 53/udp
EXPOSE 53/tcp


ENTRYPOINT ["/simpledns"]


