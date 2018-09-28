// GO语言中的RFC 4648规定的base64编码(encoding/base64包);
package encodings

import "fmt"
import "encoding/base64"
import "github.com/kingreatwill/go_study/learning/common"

// base64_;
func base64_() {
	common.Println("encoding_base64.go")

	// 使用给出的字符集生成一个*Encoding, 字符集必须是64字节的字符串;
	// func NewEncoding(encoder string) *Encoding;
	// RFC 4648定义的标准base64编码字符集;
	// var StdEncoding = NewEncoding(encodeStd);
	// RFC 4648定义的另一base64编码字符集(用于URL和文件名);
	// var URLEncoding = NewEncoding(encodeURL);

	// 返回将src编码后的字符串;
	// func (enc *Encoding) EncodeToString(src []byte) string;
	fmt.Println(base64.StdEncoding.EncodeToString([]byte("any + old & data")))

	// 返回base64编码的字符串s代表的数据;
	// func (enc *Encoding) DecodeString(s string) ([]byte, error);
	data, _ := base64.StdEncoding.DecodeString("YW55ICsgb2xkICYgZGF0YQ==")
	fmt.Println(string(data))

	// 创建一个新的base64流编码器;
	// func NewEncoder(enc *Encoding, w io.Writer) io.WriteCloser;
	// 创建一个新的base64流解码器;
	// func NewDecoder(enc *Encoding, r io.Reader) io.Reader;
}
