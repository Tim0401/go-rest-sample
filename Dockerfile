FROM golang:latest AS build-env

WORKDIR /app/go
ADD ./app /app/go
RUN go build -o main main.go

FROM alpine:latest
COPY --from=build-env /app/go/main /usr/local/bin/main
ENTRYPOINT ["/usr/local/bin/main"]