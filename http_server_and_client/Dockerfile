FROM golang:1.9
COPY http_server.go /
RUN go build -o /http_server /http_server.go
EXPOSE 8000
ENTRYPOINT ["/http_server"]
