# Copyright 2023 Wikimedia Foundation
#
# Licensed under the Apache License, Version 2.0 (the "License");
# you may not use this file except in compliance with the License.
# You may obtain a copy of the License at
#
#   http://www.apache.org/licenses/LICENSE-2.0
#
# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS,
# WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
# See the License for the specific language governing permissions and
# limitations under the License.

APPNAME     = aqs-test-proxy
VERSION     = $(shell /usr/bin/git describe --always)
BUILD_DATE  = $(shell date -u +%Y-%m-%dT%T:%Z)
HOSTNAME    = $(shell hostname)

GO_LDFLAGS  = -X main.version=$(if $(VERSION),$(VERSION),unknown)
GO_LDFLAGS += -X main.buildDate=$(if $(BUILD_DATE),$(BUILD_DATE),unknown)
GO_LDFLAGS += -X main.buildHost=$(if $(HOSTNAME),$(HOSTNAME),unknown)

build: ## build the service binary
	@echo "~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~"
	@echo "VERSION ......: $(VERSION)"
	@echo "BUILD HOST ...: $(HOSTNAME)"
	@echo "BUILD DATE ...: $(BUILD_DATE)"
	@echo "GO VERSION ...: $(word 3, $(shell go version))"
	@echo "~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~"
	go build -ldflags "$(GO_LDFLAGS)" -o $(APPNAME) .

run: ## execute the service (requires mock dataset to be running)
	go run -ldflags "$(GO_LDFLAGS)"

check: ## check for suspicious constructs, formatting errors, etc.
	@if [ -n "`goimports -l *.go`" ]; then \
	    echo "goimports: format errors detected" >&2; \
	    false; \
	fi
	@if [ -n "`gofmt -l *.go`" ]; then \
	    echo "gofmt: format errors detected" >&2; \
	    false; \
	fi
	go vet ./...

clean:  ## removes any existing binary.
	rm -f $(APPNAME)

help:  ## displays help for this makefile.
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-16s\033[0m %s\n", $$1, $$2}'

.PHONY: build run clean help
