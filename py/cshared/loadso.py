from ctypes import cdll

lib = cdll.LoadLibrary("godll.dll")

# print(lib.ForPy(4, 5))
lib.GoCall(b"1234556")
lib.Hello()
