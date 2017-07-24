from ctypes import cdll

lib = cdll.LoadLibrary("./csh.so")

print(lib.ForPy(4, 5))
