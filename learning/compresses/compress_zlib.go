// GO语言中的zlib格式压缩数据的读写(compress/zlib包);
// 在读取时解压和写入时压缩的滤镜;
// 忽略, 不研究;
package compresses

import (
	"bytes"
	"compress/zlib"
	"fmt"
	"io"
)

// 压缩;
func _zlib01(src []byte) []byte {
	var in bytes.Buffer
	w := zlib.NewWriter(&in)
	// defer w.Close() // 这样写会压缩失败;
	w.Write(src)
	w.Close() // 要这样写;
	return in.Bytes()
}

// 解压;
func _zlib02(compressSrc []byte) []byte {
	var out bytes.Buffer
	b := bytes.NewReader(compressSrc)
	r, _ := zlib.NewReader(b)
	io.Copy(&out, r)
	return out.Bytes()
}

// zlib解压缩;
func zlib01() {
	b := _zlib01([]byte("hello, world."))
	fmt.Println(b)
	fmt.Println(string(_zlib02(b)))
}
