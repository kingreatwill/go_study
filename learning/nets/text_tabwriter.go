// GO语言中的TAB对齐文本(text/tabwriter包);
package nets

import "os"
import "text/tabwriter"
import "github.com/kingreatwill/go_study/learning/common"

// tabwriter01;
func tabwriter01() {
	common.Println("text_tabwriter.go")

	// padchar != '\t'时的选项;
	// const (
	//    // 忽略html标签，并将字符实体（以'&'开始，以';'结束）视为单字符
	//    FilterHTML uint = 1 << iota
	//    // 将转义后文本两端的转义字符去掉
	//    StripEscape
	//    // 强制单元格右对齐，默认是左对齐的
	//    AlignRight
	//    // 剔除空行
	//    DiscardEmptyColumns
	//    // 使用tab而不是空格进行缩进
	//    TabIndent
	//    // 在格式化后在每一列两侧添加'|'并忽略空行
	//    Debug
	// )

	// 初始化一个Writer;
	// func NewWriter(output io.Writer, minwidth, tabwidth, padding int, padchar byte, flags uint) *Writer;
	// 初始化一个Writer;
	// func (b *Writer) Init(output io.Writer, minwidth, tabwidth, padding int, padchar byte, flags uint) *Writer;
	writer := tabwriter.NewWriter(os.Stdout, 0, 8, 0, '\t', 0)
	// writer := new(tabwriter.Writer)
	// writer.Init(os.Stdout, 0, 8, 0, '\t', 0)

	// 将buf写入b;
	// func (b *Writer) Write(buf []byte) (n int, err error);
	// 在最后一次调用Write后, 必须调用Flush方法以清空缓存, 并将格式化对齐后的文本写入生成时提供的output中;
	// func (b *Writer) Flush() (err error);
	writer.Write([]byte("a\tb\tc\td\r\n"))
	writer.Write([]byte("123456789666777\t12345\t1234567\t123\r\n"))
	writer.Flush()
}
