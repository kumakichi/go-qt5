package main

/*
#include <stdio.h>
#include <stdlib.h>

extern int start(const char *typeName);
extern void qtDebug(const char *typeName);

extern void go_print_info();
extern int go_add_fun(int, int);
extern void bind_go_export_funcs_for_fucking_ms(int, void *);

static void fucking_ms_init()
{
#define FUNC_NULL   0
#define FUNC_TWO    2
	bind_go_export_funcs_for_fucking_ms(FUNC_NULL, &go_print_info);
	bind_go_export_funcs_for_fucking_ms(FUNC_TWO, &go_add_fun);
}
*/
// #cgo LDFLAGS: -L./ -lqlib
import "C"

import (
	"fmt"
	"unsafe"
)

func main() {
	C.fucking_ms_init()
	cTypeName := C.CString("call qDebug wrapper in golang")
	C.qtDebug(cTypeName)
	C.start(cTypeName)
	C.free(unsafe.Pointer(cTypeName))
}

//export go_print_info
func go_print_info() {
	fmt.Println("I'm written in golang, No args, No return values, called by QT.")
}

//export go_add_fun
func go_add_fun(a, b C.int) C.int {
	fmt.Println("I'm written in golang, 2 int args, 1 int return value, called by QT.")
	return a + b
}
