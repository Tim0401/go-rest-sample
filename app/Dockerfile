FROM golang:latest AS build-env

WORKDIR /app/go
ADD . /app/go
RUN go get google.golang.org/appengine
RUN go build -o main main.go

FROM alpine:latest
RUN apk add --no-cache ca-certificates
RUN mkdir /lib64 && ln -s /lib/libc.musl-x86_64.so.1 /lib64/ld-linux-x86-64.so.2
COPY --from=build-env /app/go/main /usr/local/bin/main
ENTRYPOINT ["/usr/local/bin/main"]