# GCC_BIN ?= $(shell which gcc)
#cargo install --force cbindgen

GCC_BIN ?= $(shell which clang)
CARGO_BIN ?= $(shell which cargo)
CBINDGEN_BIN ?=$(shell which cbindgen)
run: clean build
	../go/hello
clean:
	$(CARGO_BIN) clean
	rm -f ../go/hello
build:
    # $(CBINDGEN_BIN) --output c/cinterop.h --crate cinterop --config cbindgen.toml
	$(CARGO_BIN) build --release
	$(GCC_BIN) -o ../go/hello ../go/main.c -Isrc  -lcinterop -L../target/release/
