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
	"time"
	"unsafe"
)

func main() {
	verify()
	minSize()
	hugeSize()
}

func hugeSize() {
	// Create a Go slice.
	goSlice := []uint8{}
	for i := 0; i < 1000000; i++ {
		goSlice = append(goSlice, (uint8)(i))
	}
	fmt.Println(len(goSlice))
	start := time.Now()
	for i := 0; i < 1000; i++ {
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
		var _ []byte = *(*[]uint8)(unsafe.Pointer(&rustSlice))
		// fmt.Printf("Before: %+v, after: %+v\n", goSlice, goSliceFromRust)
		// Free the memory allocated by Rust.
		C.free_rust_slice(rustSlice.data, rustSlice.len)
	}
	end := time.Now()
	fmt.Printf("1000 calls with 10000 elements: %+v ms\n", end.UnixMilli()-start.UnixMilli())
}

func minSize() {
	// Create a Go slice.
	goSlice := []uint8{1, 2, 3, 4, 5}
	start := time.Now()
	for i := 0; i < 1000; i++ {
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
		var _ []byte = *(*[]uint8)(unsafe.Pointer(&rustSlice))
		// fmt.Printf("Before: %+v, after: %+v\n", goSlice, goSliceFromRust)
		// Free the memory allocated by Rust.
		C.free_rust_slice(rustSlice.data, rustSlice.len)
	}
	end := time.Now()
	fmt.Printf("1000 calls with 5 elements: %+v ms\n", end.UnixMilli()-start.UnixMilli())
}

func verify() {
	// Create a Go slice.
	goSlice := []uint8{1, 2, 3, 4, 5}
	start := time.Now()
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
	var goSliceFromRust []byte = *(*[]uint8)(unsafe.Pointer(&rustSlice))
	fmt.Printf("Before: %+v, after: %+v\n", goSlice, goSliceFromRust)
	// Free the memory allocated by Rust.
	C.free_rust_slice(rustSlice.data, rustSlice.len)
	end := time.Now()
	fmt.Printf("1 calls with 5 elements: %+v micro-seconds\n", end.UnixMicro()-start.UnixMicro())
}
