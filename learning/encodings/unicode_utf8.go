// GO语言中的utf-8文本的常用函数和常数(unicode/utf16包);
// 包括rune和utf-8编码byte序列之间互相翻译的函数;
package encodings

import "fmt"
import "unicode/utf8"
import "github.com/kingreatwill/go_study/learning/common"

// utf801;
func utf801() {
	common.Println("unicode_utf8.go")

	// 判断r是否可以编码为合法的utf-8序列;
	// func ValidRune(r rune) bool;
	fmt.Println(utf8.ValidRune('a'))             // true;
	fmt.Println(utf8.ValidRune(rune(0xfffffff))) // false;

	// 返回r编码后的字节数, 如果r不是一个合法的可编码为utf-8序列的值, 会返回-1;
	// func RuneLen(r rune) int;
	fmt.Println(utf8.RuneLen('a')) // 1;
	fmt.Println(utf8.RuneLen('界')) // 3;

	// 报告字节b是否可以作为某个rune编码后的第一个字节, 第二个即之后的字节总是将左端两个字位设为10;
	buf := []byte("a界")
	fmt.Println(utf8.RuneStart(buf[0])) // true;
	fmt.Println(utf8.RuneStart(buf[1])) // true;
	fmt.Println(utf8.RuneStart(buf[2])) // false;
	fmt.Println(utf8.RuneStart(buf[3])) // false;

	// 报告切片p是否以一个码值的完整utf-8编码开始, 不合法的编码因为会被转换为宽度1的错误码值而被视为完整的;
	// func FullRune(p []byte) bool;
	// 函数类似FullRune但输入参数是字符串;
	// func FullRuneInString(s string) bool;
	buf2 := []byte("世")
	fmt.Println(utf8.FullRune(buf2))     // true;
	fmt.Println(utf8.FullRune(buf2[:2])) //false;

	// 返回p中的utf-8编码的码值的个数, 错误或者不完整的编码会被视为宽度1字节的单个码值;
	// func RuneCount(p []byte) int;
	// 函数类似RuneCount但输入参数是一个字符串;
	// func RuneCountInString(s string) (n int);
	buf3 := []byte("Hello, 世界")
	fmt.Println("bytes =", len(buf3))            // 13;
	fmt.Println("runes =", utf8.RuneCount(buf3)) // 9;

	// 返回切片p是否包含完整且合法的utf-8编码序列;
	// func Valid(p []byte) bool;
	// 报告s是否包含完整且合法的utf-8编码序列;
	// func ValidString(s string) bool;
	fmt.Println(utf8.Valid([]byte("Hello, 世界")))      // true;
	fmt.Println(utf8.Valid([]byte{0xff, 0xfe, 0xfd})) // false;

	// (常用)将r的utf-8编码序列写入p(p必须有足够的长度), 并返回写入的字节数;
	// func EncodeRune(p []byte, r rune) int;
	buf4 := make([]byte, 3)
	n := utf8.EncodeRune(buf4, '世')
	fmt.Println(buf4) // [228 184 150];
	fmt.Println(n)    // 3;

	// (常用)解码p开始位置的第一个utf-8编码的码值, 返回该码值和编码的字节数;
	// func DecodeRune(p []byte) (r rune, size int);
	b := []byte("Hello, 世界")
	for len(b) > 0 {
		r, size := utf8.DecodeRune(b)
		fmt.Printf("%c %v\n", r, size)
		b = b[size:]
	}

	// 函数类似DecodeRune但输入参数是字符串;
	// func DecodeRuneInString(s string) (r rune, size int);
	str := "Hello, 世界"
	for len(str) > 0 {
		r, size := utf8.DecodeRuneInString(str)
		fmt.Printf("%c %v\n", r, size)
		str = str[size:]
	}

	// 函数解码p中最后一个utf-8编码序列, 返回该码值和编码序列的长度;
	// func DecodeLastRune(p []byte) (r rune, size int);
	b2 := []byte("Hello, 世界") // 倒序;
	for len(b2) > 0 {
		r, size := utf8.DecodeLastRune(b2)
		fmt.Printf("%c %v\n", r, size)
		b2 = b2[:len(b2)-size]
	}

	// 函数类似DecodeLastRune但输入参数是字符串;
	// func DecodeLastRuneInString(s string) (r rune, size int);
	str2 := "Hello, 世界"
	for len(str2) > 0 {
		r, size := utf8.DecodeLastRuneInString(str2)
		fmt.Printf("%c %v\n", r, size)
		str2 = str2[:len(str2)-size]
	}
}
