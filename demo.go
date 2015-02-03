package main

/*
    extern int start(const char *typeName);
    extern void qtDebug(const char *typeName);
*/
// #include <stdio.h>
// #include <stdlib.h>
// #cgo LDFLAGS: -L./ -lqlib
import "C"

import (
    "unsafe"
    "fmt"
)

func main() {
    cTypeName := C.CString("xxoo")
    C.qtDebug(cTypeName)
    C.start(cTypeName)
    C.free(unsafe.Pointer(cTypeName))
}

//export func_written_in_go
func func_written_in_go() {
    fmt.Println("I'm written in golang, but called by QT.")
}
