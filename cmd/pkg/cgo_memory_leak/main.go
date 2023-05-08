package main

//#include "printer.h"
import "C"
import "encoding/json"

//export EncodeF64
func EncodeF64(key *C.char, value float64) *C.char {
	m := map[string]float64{
		C.GoString(key): value,
	}
	b, _ := json.Marshal(&m)
	return C.CString(string(b))
}

//
/*
	Warning: this is a memory leak
	typically we would use defer
	var cstar *C.Char = C.String("My-string")
	and
	defer C.free(unsafe.Pointer(cstr))
	and
	C.free requires #include <stdlib.h>
*/

func main() {
	C.printPiJSON()
}
