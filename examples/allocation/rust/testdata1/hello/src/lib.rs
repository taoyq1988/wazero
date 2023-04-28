// #[link(wasm_import_module = "env")]
// extern "C" {
//     /// WebAssembly import which prints a string (linear memory offset,
//     /// byteCount) to the console.
//     ///
//     /// Note: This is not an ownership transfer: Rust still owns the pointer
//     /// and ensures it isn't deallocated during this call.
//     #[link_name = "log"]
//     fn _log(ptr: u32, size: u32);
// }
//
// unsafe fn string_to_ptr(s: &str) -> (u32, u32) {
//     return (s.as_ptr() as u32, s.len() as u32);
// }

// struct Point {
//     x: bool,
//     y: f32,
// }

// #[no_mangle]
// unsafe fn test1(a: f32, b: f32) -> f32 {
//     // let s = "aaaaaaaaaa";
//     // let (ptr, len) = string_to_ptr(s);
//     // _log(ptr, len);
//     let n = a !=2.0;
//     let p1 = Point{x: n, y: b/(2 as f32)};
//     // let p1 = Point{x: n};
//     // let p1 = test2(a);
//     // let p2 = test2(a+1);
//     // if p1.x {
//         return p1.y;
//     // }
//     // return p2;
// }
//
// #[no_mangle]
// fn test2(n: f32) -> f32 {
//     return n;
// }

#[no_mangle]
fn add(a: usize) -> u32 {
    let arr:[u32;4] = [10,20,30,40];
    return arr[a];
    // return a + b;
    // return callAdd(a, b);
}

// fn callAdd(a: i32, b: i32) -> i32 {
//     let c = a + b;
//     return c;
// }

// fn test2(a: i32) -> Point {
//     let n = a !=2;
//     let p = Point{x: n, y: (a/2) as f32};
//     return p;
// }
