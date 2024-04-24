#[no_mangle]
pub extern "C" fn process_slice(input: *const u8, len: usize) -> *const u8 {
let input_slice: &[u8] = unsafe { std::slice::from_raw_parts(input, len) };
// Process the slice. For demonstration, we'll just return it as is.
// In a real-world scenario, you might modify the data or create a new slice.
println!("new_pslice len {}, {:?}, {:?}", len, input_slice, input_slice.as_ptr());
input_slice.as_ptr()
}
#[no_mangle]
pub extern "C" fn free_rust_slice(_ptr: *mut i32) {
// Since we're just returning the same slice we got from Go, we don't actually
// allocate anything in Rust, so we don't need to free anything here.
// If you were to allocate new memory in Rust to return to Go, you would use
// Box::from_raw here to convert the raw pointer back into a Box and drop it.   
}