.PHONY: test

test:
	docker build -t go-xmpplib-dep -f dep.dockerfile .
	docker run --rm \
		-v $(CURDIR):/go/src/github.com/exavolt/go-xmpplib \
		--workdir /go/src/github.com/exavolt/go-xmpplib \
		go-xmpplib-dep ensure -v -update
	docker build -t go-xmpplib-tester -f tester.dockerfile .
	docker run --rm \
		-v $(CURDIR):/go/src/github.com/exavolt/go-xmpplib \
		--workdir /go/src/github.com/exavolt/go-xmpplib \
		go-xmpplib-tester test -v ./...
