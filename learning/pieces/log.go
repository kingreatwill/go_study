// GO语言中的日志(log包);
// 主要三个函数: Print/Panic/Fatal;
// Print()行数将生成的格式化字符串输出到标准logger;
// Panic()函数等价于{Print(v...); panic(...)};
// Fatal()函数等价于{Print(v...); os.Exit(1)};
package pieces

import "io"
import "io/ioutil"
import "log"
import "os"
import "github.com/kingreatwill/go_study/learning/common"

var (
	trace   *log.Logger
	info    *log.Logger
	warning *log.Logger
	error   *log.Logger
)

// 初始化;
func logInit(
	traceWriter io.Writer,
	infoWriter io.Writer,
	warningWriter io.Writer,
	errorWriter io.Writer) {
	trace = log.New(traceWriter, "TRACE: ", log.Ldate|log.Ltime|log.Lshortfile)
	info = log.New(infoWriter, "INFO: ", log.Ldate|log.Ltime|log.Lshortfile)
	warning = log.New(warningWriter, "WARNING: ", log.Ldate|log.Ltime|log.Lshortfile)
	error = log.New(errorWriter, "ERROR: ", log.Ldate|log.Ltime|log.Lshortfile)
}

// 写入日志;
func log01() {
	common.Println("log.go")
	// 初始化(当然也可输出到文件);
	logInit(ioutil.Discard, os.Stdout, os.Stdout, os.Stderr)
	// Println将生成的格式化字符串输出到标准logger;
	trace.Println("trace1")
	info.Println("info1")
	warning.Println("warning1")
	// error.Println("error1")
	// info.Panic("info2")
	// info.Fatal("info3")
	// 将不执行下行代码(进程已退出);
	// fmt.Println("0000000000000000000000000")
}
