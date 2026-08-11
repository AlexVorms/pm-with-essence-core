FROM golang:1.25 AS builder

WORKDIR /backend

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go build -o server .

FROM debian:bookworm-slim

WORKDIR /backend

COPY --from=builder /backend/server .
COPY --from=builder /backend/config ./config
COPY --from=builder /backend/docs ./docs

EXPOSE 8080

CMD ["./server"]