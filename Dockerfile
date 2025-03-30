FROM golang:latest AS builder
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN go build -o server ./cmd/main.go

FROM scratch
COPY --from=builder /app/cmd/server/server .
CMD ["./server"]