# application name
APP ?= sqlfmt
# sanitizing app variable to become a valid go module name
MODULE = $(subst -,,$(APP))

RUN_ARGS ?= serve
COMMON_FLAGS ?= -v -mod=vendor

TEST_FLAGS ?= -failfast
TEST_EXTRA_FLAGS ?=

OUTPUT_DIR ?= build

default: build

.PHONY: vendor
vendor:
	@go mod vendor

$(OUTPUT_DIR)/$(APP):
	go build $(COMMON_FLAGS) -o="$(OUTPUT_DIR)/$(APP)" ./cmd/$(MODULE)/

build: vendor $(OUTPUT_DIR)/$(APP)

test: vendor
	go test $(COMMON_FLAGS) $(TEST_FLAGS) $(TEST_EXTRA_FLAGS) ./pkg/$(APP)/...

clean:
	@rm -rfv ./$(OUTPUT_DIR)