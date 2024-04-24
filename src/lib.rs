#[repr(C)]
pub struct GoSlice {
    data: *const u8,
    len: usize,
    cap: usize,
}
#[no_mangle]
pub extern "C" fn process_slice(input: GoSlice) -> GoSlice {
    let input_slice = unsafe { std::slice::from_raw_parts(input.data as *const u8, input.len) };
    // Process the slice. For demonstration, we'll just return it as is.
    // In a real-world scenario, you might modify the data or create a new slice.
    // If you create a new slice, you must ensure the memory is managed correctly to prevent leaks.
    // For this example, let's just double the elements and return a new Vec.
    let mut output_vec: Vec<u8> = input_slice.iter().map(|&x| x * 2).collect();
    output_vec.extend_from_slice(input_slice);
    let output_slice = output_vec.into_boxed_slice();
    let data_ptr = output_slice.as_ptr();
    let len = output_slice.len();
    std::mem::forget(output_slice); // Prevent Rust from freeing the memory.
    GoSlice {
        data: data_ptr,
        len,
        cap: len, // Capacity is the same as length for simplicity.
    }
}
#[no_mangle]
pub extern "C" fn free_rust_slice(ptr: *mut u8, len: usize) {
    unsafe {
        let _vec = Vec::from_raw_parts(ptr, len, len);
        // Vec is dropped here, and memory is freed.
    }
}
