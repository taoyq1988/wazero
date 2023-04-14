package main

// main is required for TinyGo to compile to Wasm.
func main() {}

//go:wasm-module env
//export log
//func _log(ptr uint32, size uint32)
//
//func stringToPtr(s string) (uint32, uint32) {
//	buf := []byte(s)
//	ptr := &buf[0]
//	unsafePtr := uintptr(unsafe.Pointer(ptr))
//	return uint32(unsafePtr), uint32(len(buf))
//}

type Demo struct {
	A bool
}

//export func1
func _func1() bool {
	//s := strconv.FormatUint(uint64(n), 32)
	//ptr, size := stringToPtr("aaa" + s)
	//_log(ptr, size)
	//
	//m := AAA()
	//_ = m["keyA"]
	//ptr, size = stringToPtr(m["keyA"])
	//_log(ptr, size)

	a1 := &Demo{}
	//a1.A = true
	a2 := &Demo{}

	return a1 == a2
	//return a1.A
}

//func AAA() map[string]string {
//	m := make(map[string]string)
//	m["keyA"] = "valueA"
//	return m
//}
