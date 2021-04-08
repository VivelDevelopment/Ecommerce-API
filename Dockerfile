FROM golang:latest AS builder

ADD . /app
WORKDIR /app
RUN CGO_ENABLED=0 GOOS=linux go build -o main ./cmd/server/main.go

FROM alpine:latest AS production
COPY --from=builder /app .
CMD [ "./main" ]
EXPOSE 8080