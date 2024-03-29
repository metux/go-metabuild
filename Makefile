
my_path := $(dir $(abspath $(lastword $(MAKEFILE_LIST))))

test_zlib_path    := ../zlib
test_lincity_path := ../lincity
test_xfwm4_path   := ../xfwm4

metabuild := $(my_path)/bin/metabuild
conf := $(my_path)/examples/conf/settings.yaml

all: test

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

define RUNTEST
	cd $(1) && $(metabuild) -conf $(my_path)/examples/pkg/$(strip $(2))/metabuild.yaml -global $(conf) build
endef

test-zlib: build
	$(call RUNTEST, $(test_zlib_path), zlib)

test-lincity: build
	$(call RUNTEST, $(test_lincity_path), lincity)

test-xfwm4: build
	$(call RUNTEST, $(test_xfwm4_path), xfwm4)

build:
	@rm -Rf bin
	@mkdir -p bin
	@go build -o bin/ ./cmd/...
