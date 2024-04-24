#include <cstdarg>
#include <cstdint>
#include <cstdlib>
#include <ostream>
#include <new>

extern "C" {

void print_hello_from_rust();

uint8_t print_hello_from_rust2(uint8_t a);

size_t (new_pslice(size_t len, const uint8_t *v_ptr))[2];

} // extern "C"
