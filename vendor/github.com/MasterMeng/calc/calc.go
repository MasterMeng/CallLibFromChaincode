package calc

/*
#cgo LDFLAGS: -lcalc
#include <calc.h>
*/
import "C"

func Add(a, b int) int {
	c := C.add(C.int(a), C.int(b))
	return int(c)
}

func Sub(a, b int) int {
	c := C.sub(C.int(a), C.int(b))
	return int(c)
}

func Mul(a, b int) int {
	c := C.mul(C.int(a), C.int(b))
	return int(c)
}

func Div(a, b int) float32 {
	c := C.div(C.int(a), C.int(b))
	return float32(c)
}
