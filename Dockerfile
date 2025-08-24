# ---- Build Stage ----
FROM golang:1.23-alpine AS build

# Add git & ssl
RUN apk add --no-cache ca-certificates git

# Set dir
WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

# Build app
RUN CGO_ENABLED=0 GOOS=linux go build \
    -ldflags="-w -s" \
    -v -o /app/main ./cmd/main.go

FROM alpine:latest

RUN apk add --no-cache ca-certificates tzdata wget

RUN addgroup -S appgroup && adduser -S appuser -G appgroup

WORKDIR /app

COPY --from=build /app/main .

RUN chown appuser:appgroup /app/main

USER appuser

EXPOSE 8080

ENTRYPOINT [ "/app/main", "serve" ]
