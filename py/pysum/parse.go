package main

/*

#include <Python.h>

// Workaround missing variadic function support
// https://github.com/golang/go/issues/975

int PyArg_ParseTuple_LL(PyObject * args, long long * a, long long * b) {
    return PyArg_ParseTuple(args, "LL", a, b);
}

*/
import "C"
