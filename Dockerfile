FROM golang:latest

ENV GOPROXY https://goproxy.cn,direct
WORKDIR $GOPATH/src/github.com/afsxt/simple-vote
COPY . $GOPATH/src/github.com/afsxt/simple-vote
RUN go build .

EXPOSE 8000
ENTRYPOINT ["./simple-vote"]