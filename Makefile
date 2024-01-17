PROJECT=testbot
GO              := GO111MODULE=on CGO_ENABLED=1 go
GOBUILD         := $(GO) build $(BUILD_FLAG) -tags codes
PACKAGE_LIST  := go list ./...
PACKAGE_DIRECTORIES := $(PACKAGE_LIST) | sed 's|github.com/reddio-com/$(PROJECT)/||'



default: build

build:
	make -C cairoVM build
	$(GOBUILD) -o $(PROJECT) cairoVM/cmd/main.go