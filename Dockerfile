FROM golang:1.23.9-alpine AS builder
RUN apk add --no-cache git

WORKDIR /app
COPY go.mod .
RUN go mod tidy && go mod download

COPY cmd ./cmd
COPY internal ./internal

RUN CGO_ENABLED=0 GOOS=linux go build -o task-manager ./cmd/app/main.go

FROM alpine:latest
RUN apk --no-cache add tzdata
RUN adduser -D -g '' appuser

COPY --from=builder /app/task-manager /app/task-manager

WORKDIR /app
RUN chown -R appuser:appuser /app
USER appuser
EXPOSE 8080
CMD ["./task-manager"]