FROM golang:latest

ENV GOPROXY https://goproxy.cn,direct
WORKDIR $GOPATH/src/GoIM
COPY . $GOPATH/src/GoIM
RUN go build .

EXPOSE 8080
EXPOSE 8089
EXPOSE 9001
ENTRYPOINT [ "./gowebsocket" ]