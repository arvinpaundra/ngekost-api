FROM golang:alpine AS builder

WORKDIR /app

COPY go.mod .
COPY go.sum .

RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o /app/dist/main .

FROM alpine:latest

WORKDIR /

RUN apk add --no-cache ca-certificates

COPY --from=builder /app/dist/main .
COPY --from=builder /app/version.txt .

EXPOSE 9012

CMD ["./main"]