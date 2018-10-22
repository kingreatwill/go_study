package main

import (
	"fmt"
	"reflect"
)

type I interface {
	M1()
}
type T int64

func (T) M1() {}
func (T) M2() {}
func main() {
	var i I = T(10)
	i.M1()
	//i.M2() // i.M2 undefined (type I has no field or method M2) 不能通过编译


	type T struct {

		A int

		B string

	}

	t := T{23, "skidoo"}

	s := reflect.ValueOf(&t).Elem()

	typeOfT := s.Type()

	for i := 0; i < s.NumField(); i++ {

		f := s.Field(i)

		fmt.Printf("%d: %s %s = %v\n", i,

			typeOfT.Field(i).Name, f.Type(), f.Interface())

	}
}
