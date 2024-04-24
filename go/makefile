# GCC_BIN ?= $(shell which gcc)
#cargo install --force cbindgen

GO_BIN ?= $(shell which go)
CARGO_BIN ?= $(shell which cargo)
CBINDGEN_BIN ?=$(shell which cbindgen)
run: clean build
	../go/hello
clean:
	rm -f ../go/hello
build:
	$(GCC_BIN) go build -o ../go/hello && ../go/hello
