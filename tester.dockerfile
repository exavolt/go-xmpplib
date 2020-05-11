FROM golang:1.14

WORKDIR /github.com/rez-go/xmpplib/

# Get the dependencies so it can be cached into a layer
COPY go.mod go.sum ./
RUN go mod download

ENV GO111MODULE=on

ENTRYPOINT [ "go" ]
