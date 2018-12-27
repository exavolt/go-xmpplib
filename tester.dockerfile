FROM golang:1.11

# Get dep
RUN curl https://raw.githubusercontent.com/golang/dep/master/install.sh | sh

WORKDIR /go/src/github.com/exavolt/go-xmpplib/

# Get the dependencies so it can be cached into a layer
COPY Gopkg.lock Gopkg.toml ./
RUN dep ensure -v -vendor-only

ENTRYPOINT [ "go" ]
