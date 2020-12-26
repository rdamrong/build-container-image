FROM golang:alpine as buildx
WORKDIR /myapp
COPY server.go .
RUN go build -o myserver

FROM alpine
WORKDIR /myprod
COPY --from=buildx /myapp/myserver .
ENTRYPOINT ["/myprod/myserver"]
