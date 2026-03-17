FROM golang:1.22-alpine AS build
WORKDIR /app
COPY services/go-api ./services/go-api
WORKDIR /app/services/go-api
RUN go build -o /bin/go-api ./cmd/api

FROM alpine:3.20
WORKDIR /app
COPY --from=build /bin/go-api /usr/local/bin/go-api
EXPOSE 8080
CMD ["/usr/local/bin/go-api"]
