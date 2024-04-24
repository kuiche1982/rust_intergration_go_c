#include <cstdarg>
#include <cstdint>
#include <cstdlib>
#include <ostream>
#include <new>

struct GoSlice {
  void *data;
  uintptr_t len;
  uintptr_t cap;
};

extern "C" {

GoSlice process_slice(GoSlice input);

void free_rust_slice(void *ptr, uintptr_t len);

} // extern "C"
