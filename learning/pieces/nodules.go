// GO语言的pieces小结;
package pieces

// 测试;
func Test() {
	// 创建新异常;
	errors01()

	// 写入日志;
	log01()

	// NewReader + NewReaderSize + Peek;
	bufio01()
	// NewReaderSize + Read;
	bufio02()
	// NewWriter + NewWriterSize + Write;
	bufio03()
	// NewScanner + Scan;
	bufio04()

	// Time;
	time01()
	// Duration;
	time02()
	// Ticker;
	time03()
	// Ticker;
	time04()

	// strings;
	strings01()

	// strconv;
	strconv01()

	// sort;
	sort01()

	// 遍历字段;
	reflect01()
	// 遍历方法;
	reflect02()
	// 动态调用方法;
	reflect03()
}
