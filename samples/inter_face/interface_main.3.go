package main

import (
	"fmt"
	"regexp"
)

type I interface {
	Find(b []byte) []byte
}

func f(i I) {
	fmt.Printf("%s\n", i.Find([]byte("abc")))
}
func main() {
	var re = regexp.MustCompile(`b`)
	// re 实现了Find
	f(re)
}
