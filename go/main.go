package main

/*
#include <stdint.h>

#cgo LDFLAGS: -L../target/release/ -lcinterop
size_t *new_pslice(size_t len, const char *v_ptr);
// Declare the Rust functions with C linkage
extern const int32_t* process_slice(const uint8_t* input, uintptr_t len);
extern void free_rust_slice(int32_t* ptr);
*/
import "C"
import (
	"fmt"
	"unsafe"
)

// type GPSlice struct {
// 	Values *C.char
// 	Len    uint64
// 	Cap    uint64
// }

// func main() {
// 	C.print_hello_from_rust()
// 	a := C.char(8)
// 	b := C.print_hello_from_rust2(a)
// 	fmt.Println(b)
// 	d := make([]byte, 0, 6)
// 	d = append(d, 'a', 'b')
// 	f := *(*GPSlice)(unsafe.Pointer(&d))
// 	// fmt.Println(f.Values)
// 	var e *C.ulong = C.new_pslice(C.size_t(f.Len), (*C.char)(f.Values))
// 	v := unsafe.Slice((*uint64)(e), 2)
// 	fmt.Println(v)
// 	// fmt.Println(C.pslice_len(e))

// 	// fmt.Pritnln(e)
// }

// // go build -o hello && ./hello

func main() {
	input := []uint8{1, 2, 3, 4, 5}
	// Call the Rust function
	output := C.process_slice((*C.uint8_t)(unsafe.Pointer(&input[0])), C.ulong(len(input)))
	defer C.free_rust_slice(output)
	outputSlice := (*[1 << 30]uint8)(unsafe.Pointer(output))[:len(input):len(input)]
	fmt.Println("Output slice:", outputSlice)
}
