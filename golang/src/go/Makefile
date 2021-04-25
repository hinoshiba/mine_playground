PRJ=src/goplay

GOENV=GOPATH=$(CURDIR)
GOCMD=$(GOENV) go
GOBUILD=$(GOCMD) install
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test
GOMOD=$(GOENV) go mod

#BUILD_FLGS=-a -tags netgo -installsuffix netgo -ldflags='-extldflags="static" -X main.Version=$(BUILD_VERSION)'
BUILD_FLGS=-a -tags netgo -installsuffix netgo -ldflags='-X main.Version=$(BUILD_VERSION)'
WIN_OPT=CGO_ENABLED=1 GOOS=windows GOARCH=amd64 CXX=x86_64-w64-mingw32-g++ CC=x86_64-w64-mingw32-gcc
MAC_OPT=CGO_ENABLED=1 GOOS=darwin GOARCH=amd64 CC=gcc CXX=g++

SRCS := $(shell find . -name '*.go' -type f)
BINS := $(shell test -d ./bin && find ./bin/ -type f)

all: test build ## test & build

build: $(SRCS) ## build to linux binary
	cd $(CURDIR)/$(PRJ); CGO_ENABLED=1 $(GOBUILD) $(BUILD_FLGS) goplay/exec/...

.PHONY: test
test: ## run test
	$(GOTEST) -count=1 ./src/goplay/...

.PHONY: clean
clean: $(BINS) ## cleanup
	$(GOCLEAN)
	rm -f $(BINS)

build-windows: ## build to windows binary
	cd $(CURDIR)/$(PRJ); $(WIN_OPT) $(GOBUILD) $(BUILD_FLGS) goplay/exec/...
build-mac: ## build to mac binary
	cd $(CURDIR)/$(PRJ); $(MAC_OPT) $(GOBUILD) $(BUILD_FLGS) goplay/exec/...

mod: $(CURDIR)/$(PRJ)/go.mod ## mod ensure
	cd $(CURDIR)/$(PRJ); $(GOMOD) tidy
	cd $(CURDIR)/$(PRJ); $(GOMOD) vendor
modinit: ## mod init
	cd $(CURDIR)/$(PRJ); $(GOMOD) init

.PHONY: help
	all: help
help: ## help
	@awk -F ':|##' '/^[^\t].+?:.*?##/ {\
		printf "\033[36m%-30s\033[0m %s\n", $$1, $$NF \
	}' $(MAKEFILE_LIST)
