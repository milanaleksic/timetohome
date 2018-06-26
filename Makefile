PACKAGE := $(shell go list -e)
APP_NAME = $(lastword $(subst /, ,$(PACKAGE)))

include gomakefiles/common.mk
include gomakefiles/metalinter.mk
include gomakefiles/upx.mk

SOURCES := iconunix.go iconwin.go $(shell find $(SOURCEDIR) -name '*.go' \
		-not -path './vendor/*')

$(APP_NAME): $(MAIN_APP_DIR)/$(APP_NAME)

$(MAIN_APP_DIR)/$(APP_NAME): $(SOURCES) $(BINDATA_DEBUG_FILE)
		cd $(MAIN_APP_DIR)/ && go build -ldflags '-X main.Version=${VERSION}' -o ${APP_NAME}

include gomakefiles/semaphore.mk

iconunix.go:
	@echo "//+build linux darwin" > iconunix.go
	@echo >> iconunix.go
	@cat "home.png" | 2goarray Data main >> iconunix.go

iconwin.go:
	@echo "//+build windows" > iconwin.go
	@echo >> iconwin.go
	@cat "home.ico" | 2goarray Data main >> iconwin.go

.PHONY: prepare
prepare: 
	@sudo apt-get install libgtk-3-dev libappindicator3-dev -y
	@go get github.com/cratonica/2goarray

.PHONY: clean
clean: clean_common
