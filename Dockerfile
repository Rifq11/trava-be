# --- Build Stage ---
FROM golang:1.25-alpine AS builder

WORKDIR /app

# Copy go mod files
COPY go.mod go.sum ./
RUN go mod download

# Copy seluruh folder project
COPY . .

# Build binary
RUN go build -o trava-be ./server/main.go

# --- Runtime Stage ---
FROM alpine:latest

WORKDIR /app

# Copy binary dari builder
COPY --from=builder /app/trava-be .

# Copy public directory
COPY --from=builder /app/public ./public

EXPOSE 8080

CMD ["./trava-be"]

# if you wanna run using docker use this code
# docker run -d -p 8080:8080 --add-host host.docker.internal:host-gateway -v $(pwd)/public/uploads:/root/public/uploads --name trava_be_container trava-be