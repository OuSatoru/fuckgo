package main

/*
#cgo pkg-config: python
#include <Python.h>

int PyArg_ParseTuple_LL(PyObject *, long long *, long long *);
int PyArg_ParseTuple_S(PyObject *, char **);
PyObject * out_of_disk(PyObject *, PyObject *);
PyObject * sum(PyObject *, PyObject *);
*/
import "C"
import "fmt"
import "encoding/xml"
import "unsafe"

const xmlStr = `<?xml version="1.0" encoding="UTF-8"?>
<Persons>
    <Person name="polaris" age="28">
        <Career>无业游民</Career>
        <Interests>
            <Interest>编程</Interest>
            <Interest>下棋</Interest>
        </Interests>
    </Person>
    <Person name="studygolang" age="27">
        <Career>码农</Career>
        <Interests>
            <Interest>编程</Interest>
            <Interest>下棋</Interest>
        </Interests>
    </Person>
</Persons>`

type OutOfDisk struct {
	Data DataOut `xml:"DATA"`
}

type BackOfDisk struct {
	Data DataBack `xml:"DATA"`
}

type DataOut struct {
	Head
	RowSet []RowOut `xml:"ROWSET"`
}

type DataBack struct {
	Head
	SuccCount  int       `xml:"CGBS"`
	SuccAmount float64   `xml:"CGJE"`
	RowSet     []RowBack `xml:"ROWSET"`
}

type Head struct {
	InsurType   string  `xml:"BXZL"`
	DrCr        string  `xml:"JDFX"`
	FileDate    string  `xml:"WJSCRQ"`
	TotalCount  int     `xml:"ZBS"`
	TotalAmount float64 `xml:"ZJE"`
}

type RowOut struct {
	SerialNum   string  `xml:"YWLSB"`
	ServiceCata string  `xml:"YWLB"`
	StartYM     string  `xml:"KSNY"`
	EndYM       string  `xml:"ZZNY"`
	InsuredType string  `xml:"CBDXLX"`
	InsuredNo   string  `xml:"CBDXBH"`
	Name        string  `xml:"MC"`
	CertType    string  `xml:"ZJLX"`
	CertNo      string  `xml:"ZJHM"`
	Amount      float64 `xml:"YJJE"`
	TipNo       string  `xml:"JFDJH"`
	BankNo      string  `xml:"YHBH"`
	BankAcc     string  `xml:"DYFFZH"`
	MobileNo    string  `xml:"LXDH"`
}

type RowBack struct {
	RowOut
	RetCode string `xml:"RETCODE"`
	RetInfo string `xml:"RETINFO"`
}

//export out_of_disk
func out_of_disk(self, args *C.PyObject) *C.PyObject {
	var xm *C.char
	defer C.free(unsafe.Pointer(xm))
	if C.PyArg_ParseTuple_S(args, &xm) == 0 {
		return nil
	}
	fmt.Println("Go string", C.GoString(xm))
	return C.PyString_FromString(C.CString("ingo" + C.GoString(xm)))
}

//export sum
func sum(self, args *C.PyObject) *C.PyObject {
	var a, b C.longlong
	if C.PyArg_ParseTuple_LL(args, &a, &b) == 0 {
		return nil
	}
	return C.PyLong_FromLongLong(a + b)
}

var pyMethods = []C.PyMethodDef{
	{
		C.CString("out_of_disk"),
		C.PyCFunction(C.out_of_disk),
		C.METH_VARARGS, [4]byte{},
		C.CString("Now just return."),
	},
	{
		C.CString("sum"),
		C.PyCFunction(C.sum),
		C.METH_VARARGS, [4]byte{},
		C.CString("Add two numbers."),
	},
	{nil, nil, 0, [4]byte{}, nil},
}

//export initpyxml
func initpyxml() {
	C.Py_InitModule4(C.CString("pyxml"), &pyMethods[0], C.CString("doc"),
		nil, C.PYTHON_API_VERSION)
}

func unpack(ori string) {
	var outOfDisk OutOfDisk
	xml.Unmarshal([]byte(ori), &outOfDisk)
}

func main() {

}
