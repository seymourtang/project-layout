ROOT := github.com/seymourtang/project-layout
PREFIX =
VERSION =  $(shell date -u +v%y%m%d%H%M)-$(shell git rev-parse --short HEAD)
NAME = app
COMMIT = $(shell git rev-parse HEAD)
BRANCH = $(shell git branch | grep \* | cut -d ' ' -f2)
BUILD_DATE = $(shell date -u +'%Y-%m-%dT%H:%M:%SZ')
# It's necessary to set the errexit flags for the bash shell.
export SHELLOPTS := errexit

# This will force go to use the vendor files instead of using the `$GOPATH/pkg/mod`. (vendor mode)
# more info: https://github.com/golang/go/wiki/Modules#how-do-i-use-vendoring-with-modules-is-vendoring-going-away
export GOFLAGS := -mod=vendor

# this is not a public registry; change it to your own
REGISTRY ?= hub.agoralab.co/adc/
BASE_REGISTRY ?= hub.agoralab.co/adc/

ARCH ?=
GO_VERSION ?= 1.20

CPUS ?= $(shell /bin/bash hack/read_cpus_available.sh)

# Track code version with Docker Label.
DOCKER_LABELS ?= git-describe="$(shell date -u +v%Y%m%d)-$(shell git rev-parse --short HEAD)"

GOPATH ?= $(shell go env GOPATH)
BIN_DIR := $(GOPATH)/bin
CMD_DIR := ./cmd
OUTPUT_DIR := ./bin
BUILD_DIR := ./build

build: build-linux

build-local: clean
	@echo ">> building binaries"``
	@GOOS=$(shell uname -s | tr A-Z a-z) GOARCH=$(ARCH) CGO_ENABLED=0	 \
	go build  -v -o $(OUTPUT_DIR)/$(PREFIX)$(NAME) -p $(CPUS) 		     \
		 -ldflags "-s -w						\
		  -X $(ROOT)/pkg/version.Version=$(VERSION)							 \
          -X $(ROOT)/pkg/version.GitBranch=$(BRANCH)							 \
          -X $(ROOT)/pkg/version.GitCommit=$(COMMIT)							 \
          -X $(ROOT)/pkg/version.BuildDate=$(BUILD_DATE)"					 \
	 $(CMD_DIR)/$(NAME)

build-linux:
	@docker run --rm -t                                                                \
	  -v $(PWD):/go/src/$(ROOT)                                                        \
	  -w /go/src/$(ROOT)                                                               \
	  -e GOOS=linux                                                                    \
	  -e GOARCH=amd64                                                                  \
	  -e GOPATH=/go                                                                    \
	  -e CGO_ENABLED=0																    \
	  -e GOFLAGS=$(GOFLAGS)   	                                                       \
	  -e SHELLOPTS=$(SHELLOPTS)                                                        \
	  $(BASE_REGISTRY)golang:$(GO_VERSION)                                            \
	    /bin/bash -c '                                    								\
	      	go build  -v -o $(OUTPUT_DIR)/$(PREFIX)$(NAME) -p $(CPUS)        			\
          		 -ldflags "-s -w													   \
          		 -X $(ROOT)/internal/version.Version=$(VERSION)							 \
                 -X $(ROOT)/internal/version.GitBranch=$(BRANCH)					    \
			     -X $(ROOT)/internal/version.GitCommit=$(COMMIT)					    \
                 -X $(ROOT)/internal/version.BuildDate=$(BUILD_DATE)"					 \
          	 $(CMD_DIR)/$(NAME)'                                                    			\

container: build-linux
	@echo ">> building image"
	@docker build -t $(REGISTRY)$(PREFIX)$(NAME):$(VERSION) --label $(DOCKER_LABELS)  -f $(BUILD_DIR)/$(NAME)/Dockerfile .

push:
	@echo ">> pushing image"
	@make container VERSION=$(VERSION_TAG)
	@docker push $(REGISTRY)$(PREFIX)$(NAME):$(VERSION_TAG)

clean:
	@echo ">> cleaning up"
	@-rm -vrf $(OUTPUT_DIR)
	@rm -f coverage.out