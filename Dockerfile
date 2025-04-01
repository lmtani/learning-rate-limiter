FROM golang:latest AS builder
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN GOOS=linux CGO_ENABLED=0 go build -ldflags="-w -s" -o server ./cmd/main.go

FROM scratch
WORKDIR /app
COPY --from=builder /app/server /app
CMD ["./server"]
