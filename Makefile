NAME := message

VERSION := $(shell cat VERSION)
COMMIT := $(shell git rev-parse --short HEAD)

GO_LDFLAGS :="
GO_LDFLAGS += -X main.version=$(VERSION)
GO_LDFLAGS += -X main.buildDate=$(shell date +'%Y-%m-%dT%H:%M:%SZ')
GO_LDFLAGS += -X main.gitCommit=$(COMMIT)
GO_LDFLAGS +="

.PHONY: build
build: ## Build
	@echo "+ $@"
	CGO_ENABLED=0 GOOS=linux go build \
	-installsuffix cgo \
	-ldflags $(GO_LDFLAGS) \
	-o $(NAME) .

.PHONY: clean
clean: ## Cleanu
	@echo "+ $@"
	$(RM) $(NAME)
