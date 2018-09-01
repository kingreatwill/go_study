package main

type I interface {
	M1()
}
type T int64

func (T) M1() {}
func (T) M2() {}
func main() {
	var i I = T(10)
	i.M1()
	i.M2() // i.M2 undefined (type I has no field or method M2) 不能通过编译
}
