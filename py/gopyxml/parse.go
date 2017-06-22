package main

/*
#cgo pkg-config: python
#include <Python.h>

int PyArg_ParseTuple_LL(PyObject * args, long long * a, long long * b) {
    return PyArg_ParseTuple(args, "LL", a, b);
}

int PyArg_ParseTuple_S(PyObject * args, char** str) {
	return PyArg_ParseTuple(args, "s", str);
}

*/
import "C"
