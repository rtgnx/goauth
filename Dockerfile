FROM golang:latest

RUN apt update && apt install libpam0g-dev
RUN go get github.com/rtgnx/goauth
CMD ["/go/bin/goauth"]