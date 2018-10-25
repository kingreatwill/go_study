package main

import (
	"fmt"
	"github.com/soniah/evaler"
	"math/big"
	"reflect"
	"strings"
)

func main(){

	result, _ := evaler.Eval("1+2")
	r,ok:=result.Float64()
	fmt.Println("r:",ok,r)

	user := User{5, 6.1, 3.2}
	data := Struct2Map(user)
	data2 := Struct2Map2(&user)
	fmt.Printf("struct2map得到的map内容为:%v \n", data)
	fmt.Printf("struct2map2得到的map内容为:%v \n", data2)
	ret, _ := evaler.EvalWithVariables("x + 1",  map[string]string{"x": "5"})
	fmt.Println("EvalWithVariables:",ret.Num())

	ret2, _ := evaler.EvalWithVariables("Password + 1",  data)

	fmt.Println("EvalWithVariables:",ret2.Num())

	fmt.Println("计算前：",user.Username)
	StructFormulaEqual(&user,"Username=Password + 1")

	fmt.Println("计算后：",user.Username)

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
		data[t.Field(i).Name] = fmt.Sprint(f.Interface())// fmt.Sprint(f.Interface()) 同 fmt.Sprint(f)
	}
	return data
}

func Struct2Map2(obj interface{}) map[string]string {
	var data = make(map[string]string)
	s := reflect.ValueOf(obj).Elem()
	typeOfT := s.Type()
	for i := 0; i < s.NumField(); i++ {
		f := s.Field(i)
		data[typeOfT.Field(i).Name] = fmt.Sprint(f.Interface())
	}
	return data
}
// StructFormula 计算公式 "x+1"
func StructFormula(obj interface{},formula string) (result *big.Rat, err error){
	smap := Struct2Map2(obj)
	return evaler.EvalWithVariables(formula,  smap)
}

// StructFormulaEqual 计算等式 "y=x+1"
func StructFormulaEqual(obj interface{}, formula string) {
	fs := strings.Split(formula, "=")
	if len(fs)==2 {
		if r,error := StructFormula(obj,fs[1]); error==nil{
			s := reflect.ValueOf(obj).Elem()  // 反射获取测试对象对应的struct枚举类型
			//strings.Replace(str, " ", "", -1)
			field := s.FieldByName(fs[0])
			switch t := field.Interface().(type) {
			case int:
				field.SetInt(r.Num().Int64())
			case float64:
				r,_:=r.Float64()
				field.SetFloat(r)
			default:
				_ = t
				panic("非数字")
			}
		}
	}

}
