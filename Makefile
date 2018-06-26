PACKAGE := $(shell go list -e)
APP_NAME = $(lastword $(subst /, ,$(PACKAGE)))

include gomakefiles/common.mk
include gomakefiles/metalinter.mk
include gomakefiles/upx.mk

SOURCES := $(shell find $(SOURCEDIR) -name '*.go' \
		-not -path './vendor/*')

$(APP_NAME): $(MAIN_APP_DIR)/$(APP_NAME)

$(MAIN_APP_DIR)/$(APP_NAME): $(SOURCES) $(BINDATA_DEBUG_FILE)
		cd $(MAIN_APP_DIR)/ && go build -ldflags '-X main.Version=${VERSION}' -o ${APP_NAME}

include gomakefiles/semaphore.mk

.PHONY: prepare
prepare:
	@sudo apt-get install libgtk-3-dev libappindicator3-dev

.PHONY: clean
clean: clean_common
