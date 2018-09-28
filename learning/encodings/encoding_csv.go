// GO语言中的读写逗号分隔值(csv)的文件(encoding/csv包);
package encodings

import "fmt"
import "os"
import "io"
import "strings"
import "bytes"
import "encoding/csv"
import "github.com/kingreatwill/go_study/learning/common"

// 项目实践中的用法(演示追加);
func csv00() {
	record := strings.Join([]string{"91", "92", "93"}, ",")
	root, _ := os.Getwd()
	_path := strings.Join([]string{root, "doc", "temp", "encoding_csv01.csv"}, string(os.PathSeparator))
	file, _ := os.OpenFile(_path, os.O_CREATE|os.O_APPEND|os.O_RDWR, 0660)
	defer file.Close()
	file.WriteString(record)
}

// 读取CSV文件(逐行);
func csv01() {
	common.Println("encoding_csv.go")

	// 返回一个从r读取的*Reader;
	// func NewReader(r io.Reader) *Reader;

	// 从r读取一条记录, 返回值record是字符串的切片, 每个字符串代表一个字段;
	// func (r *Reader) Read() (record []string, err error);

	root, _ := os.Getwd()
	_path := strings.Join([]string{root, "doc", "temp", "encoding_csv01.csv"}, string(os.PathSeparator))
	file, err := os.Open(_path)
	if err != nil {
		fmt.Println("Csv01 Error:", err)
		return
	}
	defer file.Close()
	reader := csv.NewReader(file)
	for {
		record, err := reader.Read()
		if err == io.EOF {
			break
		} else if err != nil {
			fmt.Println("Csv01 Error:", err)
			return
		}
		fmt.Println(record)
	}
}

// 读取CSV文件(全行);
func csv02() {
	// 返回一个从r读取的*Reader;
	// func NewReader(r io.Reader) *Reader;

	// 从r中读取所有剩余的记录, 每个记录都是字段的切片, 成功的调用返回值err为nil而不是EOF;
	// func (r *Reader) ReadAll() (records [][]string, err error);

	root, _ := os.Getwd()
	_path := strings.Join([]string{root, "doc", "temp", "encoding_csv01.csv"}, string(os.PathSeparator))
	file, err := os.Open(_path)
	if err != nil {
		fmt.Println("Csv02 Error:", err)
		return
	}
	defer file.Close()
	reader := csv.NewReader(file)
	records, err := reader.ReadAll()
	if err != nil {
		fmt.Println("Csv02 Error:", err)
		return
	}
	for i, v := range records {
		fmt.Println("Row", i, ":", v) // 也可继续嵌套循环输出每列;
	}
}

// 写入CSV文件(逐行);
func csv03() {
	// 返回一个写入w的*Writer;
	// func NewWriter(w io.Writer) *Writer;

	// 向w中写入一条记录, 会自行添加必需的引号, 记录是字符串切片, 每个字符串代表一个字段;
	// func (w *Writer) Write(record []string) (err error);

	// 将缓存中的数据写入底层的io.Writer, 要检查Flush时是否发生错误的话, 应调用Error方法;
	// func (w *Writer) Flush();

	// 返回在之前的Write方法和Flush方法执行时出现的任何错误;
	// func (w *Writer) Error() error;

	buf := new(bytes.Buffer)
	writer := csv.NewWriter(buf)
	writer.Write([]string{"41", "42", "43"})
	writer.Flush()
	e := writer.Error()
	if e != nil {
		fmt.Println("Csv03 Error:", e)
		return
	}

	root, _ := os.Getwd()
	_path := strings.Join([]string{root, "doc", "temp", "encoding_csv01.csv"}, string(os.PathSeparator))
	file, err := os.OpenFile(_path, os.O_CREATE|os.O_APPEND|os.O_RDWR, 0660)
	if err != nil {
		fmt.Println("Csv03 Error:", err)
		return
	}
	defer file.Close()
	file.WriteString(buf.String())
}

// 写入CSV文件(全行);
func csv04() {
	// 返回一个写入w的*Writer;
	// func NewWriter(w io.Writer) *Writer;

	// WriteAll方法使用Write方法向w写入多条记录, 并在最后调用Flush方法清空缓存;
	// func (w *Writer) WriteAll(records [][]string) (err error);

	buf := new(bytes.Buffer)
	writer := csv.NewWriter(buf)
	e := writer.WriteAll([][]string{{"51", "52", "53"}, {"61", "62", "63"}})
	if e != nil {
		fmt.Println("Csv04 Error:", e)
		return
	}

	root, _ := os.Getwd()
	_path := strings.Join([]string{root, "doc", "temp", "encoding_csv01.csv"}, string(os.PathSeparator))
	file, err := os.OpenFile(_path, os.O_CREATE|os.O_APPEND|os.O_RDWR, 0660)
	if err != nil {
		fmt.Println("Csv03 Error:", err)
		return
	}
	defer file.Close()
	file.WriteString(buf.String())
}
