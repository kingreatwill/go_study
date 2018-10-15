package main

//#include <studio.h>
import "C"

func main() {
	C.puts(C.CString("nihao"))
}
