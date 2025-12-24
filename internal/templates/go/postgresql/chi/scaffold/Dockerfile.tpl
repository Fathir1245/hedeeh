# =========================
# Build stage
# =========================
FROM golang:1.22-alpine AS builder

WORKDIR /app

# Install git (needed for go mod)
RUN apk add --no-cache git

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 \
    go build -o {{ .ProjectName }} ./cmd/server

# =========================
# Run stage
# =========================
FROM alpine:latest

WORKDIR /app

COPY --from=builder /app/{{ .ProjectName }} /app/{{ .ProjectName }}
COPY .env .env

EXPOSE {{ .Port }}

CMD ["./{{ .ProjectName }}"]
