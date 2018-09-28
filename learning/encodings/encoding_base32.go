// GO语言中的RFC 4648规定的base32编码(encoding/base32包);
package encodings

import "fmt"
import "encoding/base32"
import "github.com/kingreatwill/go_study/learning/common"

// base32_;
func base32_() {
	common.Println("encoding_base32.go")

	// 使用给出的字符集生成一个*Encoding, 字符集必须是32字节的字符串;
	// func NewEncoding(encoder string) *Encoding;
	// RFC 4648定义的“扩展Hex字符集”(用于DNS);
	// var HexEncoding = NewEncoding(encodeHex);
	// RFC 4648定义的标准base32编码字符集;
	// var StdEncoding = NewEncoding(encodeStd);

	// 返回将src编码后的字符串;
	// func (enc *Encoding) EncodeToString(src []byte) string;
	fmt.Println(base32.StdEncoding.EncodeToString([]byte("any + old & data")))

	// 返回base32编码的字符串s代表的数据;
	// func (enc *Encoding) DecodeString(s string) ([]byte, error);
	data, _ := base32.StdEncoding.DecodeString("MFXHSIBLEBXWYZBAEYQGIYLUME======")
	fmt.Println(string(data))

	// 创建一个新的base32流编码器;
	// func NewEncoder(enc *Encoding, w io.Writer) io.WriteCloser;
	// 创建一个新的base32流解码器;
	// func NewDecoder(enc *Encoding, r io.Reader) io.Reader;
}
