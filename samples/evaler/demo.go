package main

import (
	"fmt"
	"github.com/soniah/evaler"
	"reflect"
)

func main(){

	result, _ := evaler.Eval("1+2")
	r,ok:=result.Float64()
	fmt.Println("r:",ok,r)

	user := User{5, 6.3, 3}
	data := Struct2Map(user)
	fmt.Printf("struct2map得到的map内容为:%v", data)

	ret, _ := evaler.EvalWithVariables("x + 1",  map[string]string{"x": "5"})
	fmt.Println("EvalWithVariables:",ret.Num())

	ret2, _ := evaler.EvalWithVariables("Password + 1",  data)

	fmt.Println("EvalWithVariables:",ret2.Num())

}

type User struct {
	Id        int    `json:"id"`
	Username    float64    `json:"username"`
	Password    float64    `json:"password"`
}

func Struct2Map(obj interface{}) map[string]string {
	t := reflect.TypeOf(obj)
	v := reflect.ValueOf(obj)

	var data = make(map[string]string)
	for i := 0; i < t.NumField(); i++ {
		f := v.Field(i)
		data[t.Field(i).Name] = fmt.Sprint(f)
	}
	return data
}


