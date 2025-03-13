package main

/*
#include <stdio.h>
*/
import "C"
import "unsafe"

func main() {
	C.puts((*C.char)(unsafe.Pointer(unsafe.StringData("Hello World\x00"))))
}
