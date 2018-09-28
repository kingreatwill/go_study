// GO语言中的ascii85数据编码(encoding/ascii85包);
// 5个ascii字符表示4个字节;
// 用于btoa工具和Adobe的PostScript语言和PDF文档格式;
package encodings

// import "encoding/ascii85"
import "github.com/kingreatwill/go_study/learning/common"

// ascii85_;
func ascii85_() {
	common.Println("encoding_ascii85.go")

	// 将src编码成最多MaxEncodedLen(len(src))数据写入dst, 返回实际写入的字节数;
	// func Encode(dst, src []byte) int;
	// 将src解码后写入dst, 返回写入dst的字节数、从src解码的字节数;
	// func Decode(dst, src []byte, flush bool) (ndst, nsrc int, err error);

	// 创建一个将数据编码为ascii85流写入w的编码器, Ascii85编码算法操作32位块, 写入结束后, 必须调用Close方法将缓存中保留的不完整块刷新到w里;
	// func NewEncoder(w io.Writer) io.WriteCloser;
	// 创建一个从r解码ascii85流的解码器;
	// func NewDecoder(r io.Reader) io.Reader;
}
