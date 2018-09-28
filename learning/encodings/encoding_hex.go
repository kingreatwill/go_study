// GO语言中的16进制字符表示的编解码(encoding/hex包);
package encodings

import "fmt"
import "encoding/hex"
import "github.com/kingreatwill/go_study/learning/common"

// hex01;
func hex01() {
	common.Println("encoding_hex.go")

	src := []byte("Hello World.")

	// 编码;
	// func EncodedLen(n int) int;
	// func Encode(dst, src []byte) int;
	// func EncodeToString(src []byte) string;

	dstlen := hex.EncodedLen(len(src))
	dst := make([]byte, dstlen)
	eclen := hex.Encode(dst, src)
	fmt.Println("src =", src)
	fmt.Println("dstlen =", dstlen, ", eclen =", eclen)
	fmt.Println("hex =", dst)

	// (推荐);
	ecstring := hex.EncodeToString(src)
	fmt.Println("EncodeToString =", ecstring)

	// 解码;
	// func DecodedLen(x int) int;
	// func Decode(dst, src []byte) (int, error);
	// func DecodeString(s string) ([]byte, error);

	dclen := hex.DecodedLen(len(dst))
	dc := make([]byte, dclen)
	hex.Decode(dc, dst)
	fmt.Println("dc =", string(dc), ", dst =", dst)

	// (推荐);
	dcstring, _ := hex.DecodeString("48656c6c6f20576f726c642e")
	fmt.Println("DecodeString =", dcstring)

	// 格式化输出;
	// func Dump(data []byte) string;
	// func Dumper(w io.Writer) io.WriteCloser;
	fmt.Println(hex.Dump(dcstring))
}
