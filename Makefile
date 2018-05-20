.PHONY: fmt test update-dependencies

PKG_PATH = github.com/exavolt/go-xmpplib
DEP_IMAGE ?= go-xmpplib-dep
TESTER_IMAGE ?= go-xmpplib-tester
GOLANG_IMAGE ?= golang:1.10

fmt:
	docker run --rm \
		-v $(CURDIR):/go \
		--entrypoint gofmt \
		$(GOLANG_IMAGE) -w -l -s .

test:
	docker build -t $(TESTER_IMAGE) -f tester.dockerfile .
	docker run --rm \
		-v $(CURDIR):/go/src/$(PKG_PATH) \
		--workdir /go/src/$(PKG_PATH) \
		$(TESTER_IMAGE) test -v ./...

update-dependencies:
	docker build -t $(DEP_IMAGE) -f dep.dockerfile .
	docker run --rm \
		-v $(CURDIR):/go/src/$(PKG_PATH) \
		--workdir /go/src/$(PKG_PATH) \
		$(DEP_IMAGE) ensure -update -v
