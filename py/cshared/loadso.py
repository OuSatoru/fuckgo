from ctypes import cdll

lib = cdll.LoadLibrary("godll.dll")

lib.GoCall('aa')
