
my_path := $(dir $(abspath $(lastword $(MAKEFILE_LIST))))

test_zlib_path := ../zlib

all: test test-zlib

vet:
	@go vet ./...

fmt:
	@go fmt ./...

test: build test-util test-specobj test-autoconf

test-util:
	@go test -timeout 10s -v ./util/...             || (echo "======= TEST FAILED =======" ; false)

test-specobj:
	@go test -timeout 10s -v ./util/specobj/...     || (echo "======= TEST FAILED =======" ; false)

test-autoconf:
	@go test -timeout 10s -v ./engine/autoconf/...  || (echo "======= TEST FAILED =======" ; false)

test-zlib: build
	cd $(test_zlib_path)    && $(my_path)/bin/metabuild -conf $(my_path)/examples/pkg/zlib.yaml    -global $(my_path)/examples/settings.yaml

build:
	@rm -Rf bin
	@mkdir -p bin
	@go build -o bin/ ./cmd/...
