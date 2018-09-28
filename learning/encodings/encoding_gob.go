// GO语言中的gob流(encoding/gob包);
// gob是Golang包自带的一个数据结构序列化的编码/解码工具, 编码使用Encoder，解码使用Decoder, 一种典型的应用场景就是RPC(remote procedure calls);
// 通用数据结构序列化格式: json/xml/gob/protobuf等;
// 各个类型的编解码规则:
// 整型: 分为sign int和usign int, int和uint是不能互相编解码的, float和int也是不能互相编解码的;
// Struct/array/slice是可以被编码的, 但是function和channel是不能被编码的;
// bool类型是被当作uint来编码的, 0是false, 1是true;
// 浮点类型的值都是被当作float64类型的值来编码的;
// String和[]byte传递是uint(byte个数) + byte[]的形式编码的;
// Slice和array是按照uint(array个数) + 每个array编码 这样的形式进行编码的;
// Maps是按照 uint(Map个数) + 键值对 这样的形式进行编码的;
// Struct是按照一对对(属性名 + 属性值)来进行编码的, 其中属性值是其自己对应的gob编码;
// Struct如果有一个属性值为0或空, 则这个属性直接被忽略, 每个属性的序号是由编码时候顺序决定的, 从0开始顺序递增;
// Struct在序列化前会以-1代表序列化的开始, 以0代表序列化结束, 即Struct的序列化是按照 “-1 （0 属性1名字 属性1值） （1 属性2名字 属性2值） 0 ”来进行编码的;
// 非常重要的一点: Struct中的属性应该是public的, 即应该是大写字母开头;
// 可参考: https://www.cnblogs.com/yjf512/archive/2012/08/24/2653697.html
package encodings

import "fmt"
import "bytes"
import "encoding/gob"
import "github.com/kingreatwill/go_study/learning/common"

type A struct {
	X, Y, Z int
	Name    string
}

type B struct {
	X, Y *int32
	Name string
}

// gob01;
func gob01() {
	common.Println("encoding_gob.go")

	var network bytes.Buffer

	// 编码;
	// 返回一个将编码后数据写入w的*Encoder;
	// func NewEncoder(w io.Writer) *Encoder;
	// 将e编码后发送, 并且会保证所有的类型信息都先发送;
	// func (enc *Encoder) Encode(e interface{}) error;
	a := A{3, 4, 5, "Pythagoras"}
	encoder := gob.NewEncoder(&network)
	err1 := encoder.Encode(a)
	// err1 := encoder.Encode(&a)
	if err1 != nil {
		fmt.Println("encode error:", err1)
	}
	fmt.Println(network)
	fmt.Println(network.String())

	// 解码:
	// 返回一个从r读取数据的*Decoder;
	// func NewDecoder(r io.Reader) *Decoder;
	// 从输入流读取下一个之并将该值存入e;
	// func (dec *Decoder) Decode(e interface{}) error;
	decoder := gob.NewDecoder(&network)
	var b B
	err2 := decoder.Decode(&b)
	if err2 != nil {
		fmt.Println("Decode error:", err2)
	}
	fmt.Println(b)
	fmt.Println(*b.X, *b.Y, b.Name)
}
