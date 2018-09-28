// GO语言中的缓存IO(bufio包);
// NewReaderSize()将rd封装成一个带缓存的bufio.Reader对象, 缓存大小由size指定(如果小于16则会被设置为16);
// NewReader()相当于NewReaderSize(rd, 4096);
// NewWriterSize()将wr封装成一个带缓存的bufio.Writer对象, 缓存大小由size指定(如果小于4096则会被设置为4096);
// NewWriter()相当于NewWriterSize(wr, 4096);
package pieces

import "fmt"
import "bufio"
import "os"
import "github.com/kingreatwill/go_study/learning/common"

// NewReader + NewReaderSize + Peek;
func bufio01() {
	common.Println("bufio.go")
	// 打开文档;
	file, err1 := os.Open(`D:\研发人员项目文件\GolangingCloud-SVN\GolangingCloud-Domain-Layer\branches\V1.0.1.0\GolangingCloud.Service.Cluster\doc\temp\bufio_test.txt`)
	// 延迟关闭;
	defer file.Close()
	// 异常处理;
	if err1 != nil {
		fmt.Println(err1)
	}
	// 创建读的缓存IO;
	//reader := bufio.NewReader(file)
	// 创建读的缓存IO;
	reader := bufio.NewReaderSize(file, 4096)
	// 返回缓存IO的当前位置的下n个字节(不移动读取位置);
	bytes, err2 := reader.Peek(50)
	// 异常处理;
	if err2 != nil {
		fmt.Println(err2)
	}
	fmt.Println(string(bytes))
}

// NewReaderSize + Read;
func bufio02() {
	// 打开文档;
	file, err1 := os.Open(`D:\研发人员项目文件\GolangingCloud-SVN\GolangingCloud-Domain-Layer\branches\V1.0.1.0\GolangingCloud.Service.Cluster\doc\temp\bufio_test.txt`)
	// 延迟关闭;
	defer file.Close()
	// 异常处理;
	if err1 != nil {
		fmt.Println(err1)
	}
	// 创建读的缓存IO(默认缓存区16字节);
	reader := bufio.NewReaderSize(file, 16)
	// 切片(容量32字节);
	buf := make([]byte, 32)
	// while循环;
	for {
		// 循环读;
		n, err2 := reader.Read(buf)
		// 有数据;
		if n > 0 {
			// 最后1次?
			if n < cap(buf) {
				lastBuf := make([]byte, n)
				copy(lastBuf, buf)
				fmt.Println(string(lastBuf))
			} else {
				fmt.Println(string(buf))
			}
		}
		//		// 有数据;
		//		if n > 0 {
		//			fmt.Println(string(buf[:n]))
		//		}
		// EOF;
		if err2 != nil {
			fmt.Println(err2)
			break
		}
	}
}

// NewWriter + NewWriterSize + Write;
func bufio03() {
	// 创建文档;
	file, err1 := os.Create(`D:\研发人员项目文件\GolangingCloud-SVN\GolangingCloud-Domain-Layer\branches\V1.0.1.0\GolangingCloud.Service.Cluster\doc\temp\bufio_bufio03.txt`)
	defer file.Close()
	if err1 != nil {
		fmt.Println(err1)
	}
	// 创建写的缓存IO(默认缓存区4096字节);
	writer := bufio.NewWriter(file)
	bytes := []byte("ky testing...")
	_, err2 := writer.Write(bytes)
	if err2 != nil {
		fmt.Println(err2)
	}
	// 提交缓存中的数据;
	writer.Flush()
}

// NewScanner + Scan;
func bufio04() {
	// 创建文档;
	file, err1 := os.Open(`D:\研发人员项目文件\GolangingCloud-SVN\GolangingCloud-Domain-Layer\branches\V1.0.1.0\GolangingCloud.Service.Cluster\doc\temp\bufio_bufio04.txt`)
	defer file.Close()
	if err1 != nil {
		fmt.Println(err1)
	}
	// 创建逐行读扫描器;
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
}
