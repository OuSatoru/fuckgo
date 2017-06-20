package main

// #cgo pkg-config: python
// #include <Python.h>
import "C"

//export Sum
func Sum(self, args *C.PyObject) *C.PyObject {
	// var a, b C.longlong
	// if C.PyArg_ParseTuple_LL(args, &a, &b) == 0 {
	// 	return nil
	// }
	return C.PyLong_FromLongLong(0)
}

func main() {

}
