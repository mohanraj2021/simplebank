FROM golang:1.22-alpine AS builder

WORKDIR /app
COPY . .
RUN go build -o simplebank main.go

FROM alpine
COPY .env .env
COPY --from=builder /app/simplebank .
EXPOSE 2207
CMD ["./simplebank"]    