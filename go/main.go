package main

/*
#cgo LDFLAGS: -L../target/release/ -lcinterop
#include <stdint.h>
#include <stdlib.h>
typedef struct {
void* data;
size_t len;
size_t cap;
} GoSlice;
// Declare the Rust function to process the slice.
extern GoSlice process_slice(GoSlice input);
// Declare the Rust function to free a slice.
extern void free_rust_slice(void* ptr, size_t len);
*/
import "C"
import (
	"fmt"
	"unsafe"
)

func main() {
	// Create a Go slice.
	goSlice := []uint8{1, 2, 3, 4, 5}
	// Convert the Go slice to the C structure.
	cSlice := C.GoSlice{
		data: unsafe.Pointer(&goSlice[0]),
		len:  C.size_t(len(goSlice)),
		cap:  C.size_t(cap(goSlice)),
	}
	// Pass the Go slice to Rust.
	rustSlice := C.process_slice(cSlice)
	// Convert the C pointer back to a Go slice.
	// The length may have changed, so we use the length returned by Rust.
	goSliceFromRust := (*[1 << 30]uint8)(unsafe.Pointer(rustSlice.data))[:rustSlice.len:rustSlice.len]
	fmt.Printf("Before: %+v, after: %+v\n", goSlice, goSliceFromRust)
	// Free the memory allocated by Rust.
	C.free_rust_slice(rustSlice.data, rustSlice.len)
}
