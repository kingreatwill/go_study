// GO语言中的简单的数字与字节序列的转换以及变长值的编解码(encoding/binary包);
package encodings

import "fmt"
import "encoding/binary"
import "bytes"
import "math"
import "github.com/kingreatwill/go_study/learning/common"

// binary01;
func binary01() {
	common.Println("encoding_binary.go")

	// 从r中读取binary编码的数据并赋给data, data必须是一个指向定长值的指针或者定长值的切片;
	// func Read(r io.Reader, order ByteOrder, data interface{}) error;
	var pi1 float64
	b := []byte{0x18, 0x2d, 0x44, 0x54, 0xfb, 0x21, 0x09, 0x40}
	buf1 := bytes.NewReader(b)
	err1 := binary.Read(buf1, binary.LittleEndian, &pi1)
	if err1 != nil {
		fmt.Println("binary.Read failed:", err1)
	}
	fmt.Println(pi1) // 3.141592653589793

	// 将data的binary编码格式写入w, data必须是定长值、定长值的切片、定长值的指针;
	// func Write(w io.Writer, order ByteOrder, data interface{}) error;
	buf2 := new(bytes.Buffer)
	var pi2 float64 = math.Pi
	err2 := binary.Write(buf2, binary.LittleEndian, pi2)
	if err2 != nil {
		fmt.Println("binary.Write failed:", err2)
	}
	fmt.Printf("% x\n", buf2.Bytes())

	// Multi;
	buf3 := new(bytes.Buffer)
	var data3 = []interface{}{
		uint16(61374),
		int8(-54),
		uint8(254),
	}
	for _, v := range data3 {
		err3 := binary.Write(buf3, binary.LittleEndian, v)
		if err3 != nil {
			fmt.Println("binary.Write failed:", err3)
		}
	}
	fmt.Printf("%x\n", buf3.Bytes()) // beefcafe;
}
