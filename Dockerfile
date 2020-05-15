FROM golang:alpine AS builder
WORKDIR /app/
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build giphyneitor/main.go

FROM gcr.io/distroless/base
WORKDIR /app/
COPY --from=builder /app/main .
COPY --from=builder /app/giphyneitor/templates templates/.
CMD ["./main"]