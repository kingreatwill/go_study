// GO语言中的json对象的编解码(encoding/json包);
// 请参考: https://www.oschina.net/translate/advanced-encoding-decoding
package encodings

import "fmt"
import "encoding/json"
import "reflect"
import "github.com/kingreatwill/go_study/learning/common"

// 结构体转换为JSON;
// func Marshal(v interface{}) ([]byte, error);
func struct2json() {
	common.Println("encoding_json.go")

	type ColorGroup struct {
		ID     int      `json:"id" mark:"编号"`
		Name   string   `json:"name" mark:"姓名"`
		Colors []string `json:"colors" mark:"颜色"`
		X      string   `json:"-" mark:"其他"`
	}

	group := ColorGroup{
		ID:     1,
		Name:   "Reds",
		Colors: []string{"Crimson", "Red", "Ruby", "Maroon"},
	}
	// 反射;
	v := reflect.ValueOf(group)
	t := v.Type()
	for i, n := 0, t.NumField(); i < n; i++ {
		fmt.Printf("%s: %v\n", t.Field(i).Tag.Get("mark"), v.Field(i))
	}
	// 转换;
	b, err := json.Marshal(group)
	if err != nil {
		fmt.Println("error:", err)
	}
	fmt.Println(string(b))
}

// JSON转换为结构体;
// func Unmarshal(data []byte, v interface{}) error;
func json2struct() {
	type Animal struct {
		Name  string
		Order string
	}

	var blob = []byte(`[
		{"Name": "Platypus", "Order": "Monotremata"},
		{"Name": "Quoll",    "Order": "Dasyuromorphia"}
	]`)

	var animals []Animal
	err := json.Unmarshal(blob, &animals)
	if err != nil {
		fmt.Println("error:", err)
	}
	fmt.Println(animals[0].Name, animals[0].Order)
}
