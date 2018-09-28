// GO语言中的运行时反射(reflect包);
// 典型用法:
// 1.TypeOf: 用静态类型interface{}保存一个值, 通过调用TypeOf获取其动态类型信息, 该函数返回一个Type类型值;
// 2.ValueOf: 调用ValueOf函数返回一个Value类型值, 该值代表运行时的数据;
// 3.Zero: Zero接受一个Type类型参数并返回一个代表该类型零值的Value类型值;
package pieces

import "fmt"
import "reflect"
import "github.com/kingreatwill/go_study/learning/common"

type Panda struct {
	FirstName string `tag_name:"tag 1"`
	LastName  string `tag_name:"tag 2"`
	Age       int    `tag_name:"tag 3"`
}

func (this *Panda) GetName(prefix string) string {
	return prefix + ": " + this.FirstName + " " + this.LastName
}

// 遍历字段;
func reflect01() {
	common.Println("reflect.go")

	// 引用类型;
	f := Panda{
		FirstName: "KY",
		LastName:  "Z",
		Age:       20,
	}
	v := reflect.ValueOf(f)

	// 指针类型;
	// f := &Panda{
	//		FirstName: "KY",
	//		LastName:  "Z",
	//		Age:       20,
	// }
	// v := reflect.ValueOf(f).Elem()

	t := reflect.TypeOf(f)
	//t := v.Type()

	for i := 0; i < v.NumField(); i++ {
		field := t.Field(i)
		fmt.Printf("Field Name: %s,\t Field Value: %v,\t Tag Value: %s\n", field.Name, v.Field(i), field.Tag.Get("tag_name"))
	}

	// Field Name: FirstName,	 Field Value: KY,	 Tag Value: tag 1
	// Field Name: LastName,	 Field Value: Z,	 Tag Value: tag 2
	// Field Name: Age,          Field Value: 20,	 Tag Value: tag 3
}

// 遍历方法;
func reflect02() {
	f := new(Panda)
	v := reflect.ValueOf(f)
	t := reflect.TypeOf(f)
	for i := 0; i < v.NumMethod(); i++ {
		fmt.Printf("Method Name: %s\n", t.Method(i).Name)
	}
}

// 动态调用方法;
func reflect03() {
	f := &Panda{
		FirstName: "KY",
		LastName:  "Z",
		Age:       20,
	}
	r := _DynamicCall(f, "GetName", "TESTING")
	fmt.Println(r[0])
}

func _DynamicCall(_struct interface{}, methodName string, params ...interface{}) []reflect.Value {
	ins := make([]reflect.Value, len(params))
	for i, v := range params {
		ins[i] = reflect.ValueOf(v)
	}
	return reflect.ValueOf(_struct).MethodByName(methodName).Call(ins)
}
