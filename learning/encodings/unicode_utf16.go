// GO语言中的utf-16序列的编解码(unicode/utf16包);
package encodings

import "fmt"
import "unicode/utf16"
import "github.com/kingreatwill/go_study/learning/common"

// utf1601;
func utf1601() {
	common.Println("unicode_utf16.go")

	// 返回r是否可以编码为一个utf-16的代理对;
	// func IsSurrogate(r rune) bool;

	// 将unicode码值r编码为一个utf-16的代理对, 如果不能编码, 会返回(U+FFFD, U+FFFD);
	// func EncodeRune(r rune) (r1, r2 rune);
	// 将utf-16代理对(r1, r2)解码为unicode码值, 如果代理对不合法, 会返回U+FFFD;
	// func DecodeRune(r1, r2 rune) rune;

	// 将unicode码值序列编码为utf-16序列;
	// func Encode(s []rune) []uint16;
	s := utf16.Encode([]rune("golang中文unicode编码"))
	fmt.Println("utf16.Encode:", s)

	// 将utf-16序列解码为unicode码值序列;
	// func Decode(s []uint16) []rune;
	fmt.Println("utf16.Decode:", string(utf16.Decode(s)))
}
