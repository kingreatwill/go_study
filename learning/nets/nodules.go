// GO语言的nets小结;
package nets

// Test 测试;
func Test() {
	// html01;
	html01()

	// neturl01;
	neturl01()
	// neturl02;
	neturl02()
	// neturl03;
	neturl03()

	// tabwriter01;
	tabwriter01()
	// 文本模板数据驱动输出(简单例子);
	text_template01()
	// 其他函数;
	text_template02()
	// 文本模板数据驱动输出;
	text_template03()
	// 文本模板数据驱动输出(常用);
	text_template04()
	// 文本模板数据驱动输出(FuncMap);
	text_template05()

	// expvar01;
	expvar01()

	// type IP []byte;;
	net_IP()
	// type IPMask []byte;
	net_IPMask()
	// type HardwareAddr []byte;
	net_HardwareAddr()
	// type IPAddr/TCPAddr/UDPAddr/UnixAddr struct;
	net_IPAddr_TCPAddr_UDPAddr_UnixAddr()
	// type IPNet struct;
	net_IPNet()
	// 最终常用的连接与侦听;
	net_Summary09()

	// 简单示例;
	nethttp_main()
}
