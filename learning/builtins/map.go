// GO语言中的字典(map);
package builtins

import "fmt"
import "github.com/kingreatwill/go_study/learning/common"

// 取值;
func map01() {
	common.Println("map.go")
	// 字典声明;
	colors := map[string]string{
		"bird":  "blue",
		"snake": "green",
		"cat":   "black",
	}
	// 取值;
	c := colors["snake"]
	fmt.Println(c)
}

// 赋值;
func map02() {
	// 字典声明;
	names := map[int]string{}
	// 赋值;
	names[990] = "file.txt"
	names[1009] = "data.xls"
	names[1209] = "image.jpg"
	fmt.Println(len(names))
}

// 删除;
func map03() {
	// 字典声明;
	ids := map[string]int{}
	// 赋值;
	ids["steve"] = 10
	ids["mark"] = 20
	ids["adnan"] = 30
	// 删除;
	delete(ids, "steve")
	fmt.Println(len(ids))
}

// 遍历;
func map04() {
	// 字典声明;
	animals := map[string]string{}
	animals["cat"] = "Mittens"
	animals["dog"] = "Spot"
	// 遍历;
	for key, value := range animals {
		fmt.Println(key, "=", value)
	}
}

// 遍历并拷贝KEY;
func map05() {
	// 字典声明;
	sizes := map[string]int{
		"XL": 20,
		"L":  10,
		"M":  5,
	}
	// 切片声明;
	keys := []string{}
	// 遍历并拷贝KEY;
	for key, _ := range sizes {
		keys = append(keys, key)
	}
	fmt.Println(keys)
}

// 引用类型必须使用make函数创建;
func map06() {
	// 引用类型必须使用make函数创建;
	lookup := make(map[string]int, 200)
	// 赋值;
	lookup["cat"] = 10
	// 取值;
	fmt.Println(lookup["cat"])
}

// 其他补充;
func map07() {
	// 字典声明;
	var _map map[string]string
	// 引用类型必须使用make函数创建;
	_map = make(map[string]string)
	// 初始化数据;
	_map["France"] = "Paris"
	_map["Italy"] = "Rome"
	_map["Japan"] = "Tokyo"
	_map["India"] = "New Delhi"
	// 字典遍历;
	for key := range _map {
		fmt.Println("Capital of ", key, " is "+_map[key]+".")
	}
	// 根据key获取value(方式1);
	_value, b := _map["United States"]
	if b {
		fmt.Println("Capital of United States is " + _value + ".")
	} else {
		fmt.Println("Capital of United States is not present.")
	}
	// 根据key获取value(方式2);
	_value = _map["United States"]
	if _value != "" {
		fmt.Println("Capital of United States is " + _value + ".")
	} else {
		fmt.Println("Capital of United States is not present.")
	}
}
