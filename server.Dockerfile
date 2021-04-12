FROM golang:1.16.2-buster

WORKDIR /src
COPY server server
ENTRYPOINT ["go", "run", "server/main.go"]
