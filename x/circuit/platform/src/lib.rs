// Import just to include library to link against.
#[allow(unused_imports)]
use risc0_zkvm_platform;

// When std is not linked, register a panic handler here so the user does not
// have to. If std is linked, it will define the panic handler instead. This
// panic handler must not be included.
// #[cfg(all(target_os = "zkvm", not(feature = "std")))]
// #[panic_handler]
// pub fn panic_impl(panic_info: &core::panic::PanicInfo) -> ! {
//     risc0_zkvm_platform::rust_rt::panic_fault(panic_info);
// }