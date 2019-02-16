.PHONY: fmt test update-dependencies

PKG_PATH = github.com/exavolt/go-xmpplib
TESTER_IMAGE ?= go-xmpplib-tester
GOLANG_IMAGE ?= golang:1.11

fmt:
	@echo "Formatting files..."
	@docker run --rm \
		-v $(CURDIR):/go \
		--entrypoint gofmt \
		$(GOLANG_IMAGE) -w -l -s .

test:
	@echo "Preparing test runner..."
	@docker build -t $(TESTER_IMAGE) -f tester.dockerfile . > /dev/null
	@echo "Executing unit tests..."
	@docker run --rm \
		-v $(CURDIR):/go/src/$(PKG_PATH) \
		--workdir /go/src/$(PKG_PATH) \
		$(TESTER_IMAGE) test ./...

update-dependencies:
	@echo "Updating all dependencies..."
	@docker run --rm \
		-v $(CURDIR):/$(PKG_PATH) \
		--workdir /$(PKG_PATH) \
		$(GOLANG_IMAGE) /bin/sh -c "go get -u all && go mod tidy"
