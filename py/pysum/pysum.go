package main

/*
#cgo pkg-config: python
#include <Python.h>

int PyArg_ParseTuple_LL(PyObject *, long long *, long long *);
PyObject * sum(PyObject *, PyObject *);
*/
import "C"

//export sum
func sum(self, args *C.PyObject) *C.PyObject {
	var a, b C.longlong
	if C.PyArg_ParseTuple_LL(args, &a, &b) == 0 {
		return nil
	}
	return C.PyLong_FromLongLong(a + b)
}

//export initpys
func initpys() {
	C.Py_InitModule4(C.CString("pys"), &sumMethods[0], C.CString("doc"),
		nil, C.PYTHON_API_VERSION)
}

var sumMethods = []C.PyMethodDef{
	{
		C.CString("sum"),
		C.PyCFunction(C.sum),
		C.METH_VARARGS, [4]byte{},
		C.CString("Add two numbers."),
	},
	{nil, nil, 0, [4]byte{}, nil},
}

func main() {

}
