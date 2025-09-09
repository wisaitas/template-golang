FROM golang:1.25-alpine AS builder

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download && go mod verify

COPY cmd/service/ ./cmd/service
COPY internal/service/ ./internal/service
COPY pkg/ ./pkg

RUN go mod tidy

RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main cmd/service/main.go

FROM alpine:latest

RUN apk --no-cache add curl

WORKDIR /app

COPY --from=builder /app/main .

CMD ["./main"]