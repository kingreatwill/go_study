// GO语言中的可移植的网络I/O接口(net包);
package nets

import "fmt"
import "net"
import "github.com/kingreatwill/go_study/learning/common"

// type IP []byte;
func net_IP() {
	common.Println("net.go")

	// IP类型是代表单个IP地址的[]byte切片;
	// type IP []byte;

	// 返回包含一个IPv4地址a.b.c.d的IP地址(16字节格式);
	// func IPv4(a, b, c, d byte) IP;
	// 将s解析为IP地址(IPv4格式: "74.125.19.99", 或IPv6格式: 2001:4860:0:2001::68);
	// func ParseIP(s string) IP;
	ip := net.ParseIP("74.125.19.99")
	fmt.Println(ip) // 74.125.19.99;

	// 是否全局单播IP地址;
	// func (ip IP) IsGlobalUnicast() bool;
	// 是否链路本地单播IP地址;
	// func (ip IP) IsLinkLocalUnicast() bool;
	// 是否接口本地组播IP地址;
	// func (ip IP) IsInterfaceLocalMulticast() bool;
	// 是否链路本地组播IP地址;
	// func (ip IP) IsLinkLocalMulticast() bool;
	// 是否组播地址;
	// func (ip IP) IsMulticast() bool;
	// 是否环回地址;
	// func (ip IP) IsLoopback() bool;
	// 是否未指定地址;
	// func (ip IP) IsUnspecified() bool;
	fmt.Println(ip.IsGlobalUnicast()) // true;

	// 返回IP地址的默认子网掩码(只有IPv4有默认子网掩码);
	// func (ip IP) DefaultMask() IPMask;
	fmt.Println(ip.DefaultMask()) // ff000000;
	// 比较IP地址(代表同一地址的IPv4地址和IPv6地址也被认为是相等的);
	// func (ip IP) Equal(x IP) bool;
	fmt.Println(ip.Equal(net.ParseIP("2001:4860:0:2001::68"))) // false;
	// 将IP地址转换为16字节表示;
	// func (ip IP) To16() IP;
	fmt.Println(ip.To16()) // 74.125.19.99;
	// 将IPv4地址转换为4字节表示;
	// func (ip IP) To4() IP;
	fmt.Println(ip.To4()) // 74.125.19.99;
	// 根据子网掩码mask, 返回网络地址部分的IP;
	// func (ip IP) Mask(mask IPMask) IP;
	fmt.Println(ip.Mask(ip.DefaultMask())) // 74.0.0.0;
	// 返回IP地址ip的字符串(IPv4格式: "74.125.19.99", 或IPv6格式: 2001:4860:0:2001::68);
	// func (ip IP) String() string;

	// 实现encoding.TextMarshaler接口, 返回值和String方法一样;
	// func (ip IP) MarshalText() ([]byte, error);
	// 实现encoding.TextUnmarshaler接口(IP地址字符串应该是ParseIP函数可以接受的格式);
	// func (ip *IP) UnmarshalText(text []byte) error;
}

// type IPMask []byte;
func net_IPMask() {
	// IP地址的掩码;
	// type IPMask []byte;

	// 返回一个4字节格式的IPv4掩码a.b.c.d;
	// func IPv4Mask(a, b, c, d byte) IPMask;
	// 返回一个IPMask类型值, 该返回值总共有bits个字位, 其中前ones个字位都是1, 其余字位都是0;
	// func CIDRMask(ones, bits int) IPMask;
	// 返回m的前导的1字位数和总字位数;
	// func (m IPMask) Size() (ones, bits int);
	// 返回m的十六进制格式(没有标点);
	// func (m IPMask) String() string;
}

// type HardwareAddr []byte;
func net_HardwareAddr() {
	// 硬件地址(MAC地址);
	// type HardwareAddr []byte;

	// 解析一个IEEE 802 MAC-48、EUI-48或EUI-64硬件地址;
	//func ParseMAC(s string) (hw HardwareAddr, err error);
	// 返回硬件地址(MAC地址)字符串;
	//func (a HardwareAddr) String() string;
}

// type IPAddr/TCPAddr/UDPAddr/UnixAddr struct;
func net_IPAddr_TCPAddr_UDPAddr_UnixAddr() {
	// type IPAddr struct {
	//    IP   IP
	//    Zone string // IPv6范围寻址域
	// }

	// 将addr作为一个格式为"host"或"ipv6-host%zone"的IP地址来解析(net必须是"ip"、"ip4"或"ip6");
	// func ResolveIPAddr(net, addr string) (*IPAddr, error);
	// 返回地址的网络类型: "ip";
	// func (a *IPAddr) Network() string;
	// 返回地址字符串(IP + Zone);
	// func (a *IPAddr) String() string;
	ipaddr, _ := net.ResolveIPAddr("ip", "www.baidu.com")
	fmt.Println(ipaddr.IP, ipaddr.Zone) // 14.215.177.38 ;
	fmt.Println(ipaddr.Network())       // ip;
	fmt.Println(ipaddr.String())        // 14.215.177.38;

	// type TCPAddr struct {
	//    IP   IP
	//    Port int
	//    Zone string // IPv6范围寻址域
	// }
	// 将addr作为TCP地址解析并返回, 参数addr格式为"host:port"或"[ipv6-host%zone]:port", IPv6地址字面值/名称必须用方括号包起来，如"[::1]:80"、"[ipv6-host]:http"或"[ipv6-host%zone]:80", 解析得到网络名和端口名(net必须是"tcp"、"tcp4"或"tcp6");
	// func ResolveTCPAddr(net, addr string) (*TCPAddr, error);
	// func (a *TCPAddr) Network() string;
	// func (a *TCPAddr) String() string;

	// type UDPAddr struct {
	//    IP   IP
	//    Port int
	//    Zone string // IPv6范围寻址域
	// }
	// 将addr作为UDP地址解析并返回, 参数addr格式为"host:port"或"[ipv6-host%zone]:port"，IPv6地址字面值/名称必须用方括号包起来，如"[::1]:80"、"[ipv6-host]:http"或"[ipv6-host%zone]:80", 解析得到网络名和端口名(net必须是"udp"、"udp4"或"udp6");
	// func ResolveUDPAddr(net, addr string) (*UDPAddr, error);
	// func (a *UDPAddr) Network() string;
	// func (a *UDPAddr) String() string;

	// type UnixAddr struct {
	//    Name string
	//    Net  string
	// }
	// 将addr作为Unix域socket地址解析, 参数net指定网络类型："unix"、"unixgram"或"unixpacket";
	// func ResolveUnixAddr(net, addr string) (*UnixAddr, error);
	// func (a *UnixAddr) Network() string;
	// func (a *UnixAddr) String() string;
}

// type IPNet struct;
func net_IPNet() {
	// type IPNet struct {
	//    IP   IP     // 网络地址
	//    Mask IPMask // 子网掩码
	// }

	// 将s作为一个CIDR(无类型域间路由)的IP地址和掩码字符串;
	// func ParseCIDR(s string) (IP, *IPNet, error);
	ip, ipnet, _ := net.ParseCIDR("192.168.100.1/24")
	fmt.Println(ip, ipnet, ipnet.Mask) // 192.168.100.1 192.168.100.0/24 ffffff00;
	// 该网络是否包含IP地址;
	// func (n *IPNet) Contains(ip IP) bool;
	fmt.Println(ipnet.Contains(net.ParseIP("192.168.100.2"))) // true;
	// 返回网络类型名: ip+net;
	// func (n *IPNet) Network() string;
	fmt.Println(ipnet.Network()) // ip+net;
	// 返回n的CIDR表示(192.168.100.1/24或2001:DB8::/48);
	// func (n *IPNet) String() string;
	fmt.Println(ipnet.String()) // 192.168.100.0/24;
}

// 两个重要接口(interface);
func net_Summary01() {
	// #############################################################连接#############################################################
	// 通用的面向流的网络连接;
	// type Conn interface {
	//    // Read从连接中读取数据
	//    // Read方法可能会在超过某个固定时间限制后超时返回错误，该错误的Timeout()方法返回真
	//    Read(b []byte) (n int, err error)
	//    // Write从连接中写入数据
	//    // Write方法可能会在超过某个固定时间限制后超时返回错误，该错误的Timeout()方法返回真
	//    Write(b []byte) (n int, err error)
	//    // Close方法关闭该连接
	//    // 并会导致任何阻塞中的Read或Write方法不再阻塞并返回错误
	//    Close() error
	//    // 返回本地网络地址
	//    LocalAddr() Addr
	//    // 返回远端网络地址
	//    RemoteAddr() Addr
	//    // 设定该连接的读写deadline，等价于同时调用SetReadDeadline和SetWriteDeadline
	//    // deadline是一个绝对时间，超过该时间后I/O操作就会直接因超时失败返回而不会阻塞
	//    // deadline对之后的所有I/O操作都起效，而不仅仅是下一次的读或写操作
	//    // 参数t为零值表示不设置期限
	//    SetDeadline(t time.Time) error
	//    // 设定该连接的读操作deadline，参数t为零值表示不设置期限
	//    SetReadDeadline(t time.Time) error
	//    // 设定该连接的写操作deadline，参数t为零值表示不设置期限
	//    // 即使写入超时，返回值n也可能>0，说明成功写入了部分数据
	//    SetWriteDeadline(t time.Time) error
	// }

	// 在网络network上连接地址address, 并返回一个Conn接口, 可用的网络类型有: "tcp"、"tcp4"、"tcp6"、"udp"、"udp4"、"udp6"、"ip"、"ip4"、"ip6"、"unix"、"unixgram"、"unixpacket";
	// func Dial(network, address string) (Conn, error);
	// 类似Dial但采用了超时(timeout参数如果必要可包含名称解析);
	// func DialTimeout(network, address string, timeout time.Duration) (Conn, error);
	// 创建一个内存中的同步、全双工网络连接, 连接的两端都实现了Conn接口, 一端的读取对应另一端的写入, 直接将数据在两端之间作拷贝(没有内部缓冲);
	// func Pipe() (Conn, Conn);

	// 对TCP和UDP网络, 地址格式是host:port或[host]:port;
	// conn, _ := net.Dial("tcp", "12.34.56.78:80")
	// conn, _ := net.Dial("tcp", "google.com:http")
	// conn, _ := net.Dial("tcp", "[2001:db8::1]:http")
	// conn, _ := net.Dial("tcp", "[fe80::1%lo0]:80")
	// 对IP网络, network必须是"ip"、"ip4"、"ip6"后跟冒号和协议号或者协议名, 地址必须是IP地址字面值;
	// conn, _ := net.Dial("ip4:1", "127.0.0.1")
	// conn, _ := net.Dial("ip6:ospf", "::1")

	// #############################################################侦听#############################################################
	// 面向流的网络协议的公用的网络监听器接口;
	// type Listener interface {
	//    // Addr返回该接口的网络地址
	//    Addr() Addr
	//    // Accept等待并返回下一个连接到该接口的连接
	//    Accept() (c Conn, err error)
	//    // Close关闭该接口，并使任何阻塞的Accept操作都会不再阻塞并返回错误。
	//    Close() error
	// }

	// 返回在一个本地网络地址laddr上监听的Listener, 网络类型参数net必须是面向流的网络: "tcp"、"tcp4"、"tcp6"、"unix"或"unixpacket");
	// func Listen(net, laddr string) (Listener, error);

	// #############################################################其他#############################################################
	// 通用的面向数据包的网络连接;
	// type PacketConn interface {
	//    // ReadFrom方法从连接读取一个数据包，并将有效信息写入b
	//    // ReadFrom方法可能会在超过某个固定时间限制后超时返回错误，该错误的Timeout()方法返回真
	//    // 返回写入的字节数和该数据包的来源地址
	//    ReadFrom(b []byte) (n int, addr Addr, err error)
	//    // WriteTo方法将有效数据b写入一个数据包发送给addr
	//    // WriteTo方法可能会在超过某个固定时间限制后超时返回错误，该错误的Timeout()方法返回真
	//    // 在面向数据包的连接中，写入超时非常罕见
	//    WriteTo(b []byte, addr Addr) (n int, err error)
	//    // Close方法关闭该连接
	//    // 会导致任何阻塞中的ReadFrom或WriteTo方法不再阻塞并返回错误
	//    Close() error
	//    // 返回本地网络地址
	//    LocalAddr() Addr
	//    // 设定该连接的读写deadline
	//    SetDeadline(t time.Time) error
	//    // 设定该连接的读操作deadline，参数t为零值表示不设置期限
	//    // 如果时间到达deadline，读操作就会直接因超时失败返回而不会阻塞
	//    SetReadDeadline(t time.Time) error
	//    // 设定该连接的写操作deadline，参数t为零值表示不设置期限
	//    // 如果时间到达deadline，写操作就会直接因超时失败返回而不会阻塞
	//    // 即使写入超时，返回值n也可能>0，说明成功写入了部分数据
	//    SetWriteDeadline(t time.Time) error
	// }

	// 监听本地网络地址laddr。网络类型net必须是面向数据包的网络类型: "ip"、"ip4"、"ip6"、"udp"、"udp4"、"udp6"、或"unixgram";
	// func ListenPacket(net, laddr string) (PacketConn, error);
}

// 四种Conn(struct);
func net_Summary02() {
	// IPConn类型代表IP网络连接(实现了Conn和PacketConn接口);
	// type IPConn struct { ... };
	// func DialIP(netProto string, laddr, raddr *IPAddr) (*IPConn, error);
	// func ListenIP(netProto string, laddr *IPAddr) (*IPConn, error);
	// func (c *IPConn) LocalAddr() Addr;
	// func (c *IPConn) RemoteAddr() Addr;
	// func (c *IPConn) SetReadBuffer(bytes int) error;
	// func (c *IPConn) SetWriteBuffer(bytes int) error;
	// func (c *IPConn) SetDeadline(t time.Time) error;
	// func (c *IPConn) SetReadDeadline(t time.Time) error;
	// func (c *IPConn) SetWriteDeadline(t time.Time) error;
	// func (c *IPConn) Read(b []byte) (int, error);
	// func (c *IPConn) ReadFrom(b []byte) (int, Addr, error);
	// func (c *IPConn) ReadFromIP(b []byte) (int, *IPAddr, error);
	// func (c *IPConn) ReadMsgIP(b, oob []byte) (n, oobn, flags int, addr *IPAddr, err error);
	// func (c *IPConn) Write(b []byte) (int, error);
	// func (c *IPConn) WriteTo(b []byte, addr Addr) (int, error);
	// func (c *IPConn) WriteToIP(b []byte, addr *IPAddr) (int, error);
	// func (c *IPConn) WriteMsgIP(b, oob []byte, addr *IPAddr) (n, oobn int, err error);
	// func (c *IPConn) Close() error;
	// func (c *IPConn) File() (f *os.File, err error);

	// TCPConn类型代表TCP网络连接(实现了Conn接口);
	// type TCPConn struct { ... };
	// func DialTCP(net string, laddr, raddr *TCPAddr) (*TCPConn, error);
	// func (c *TCPConn) LocalAddr() Addr;
	// func (c *TCPConn) RemoteAddr() Addr;
	// func (c *TCPConn) SetReadBuffer(bytes int) error;
	// func (c *TCPConn) SetWriteBuffer(bytes int) error;
	// func (c *TCPConn) SetDeadline(t time.Time) error;
	// func (c *TCPConn) SetReadDeadline(t time.Time) error;
	// func (c *TCPConn) SetWriteDeadline(t time.Time) error;
	// func (c *TCPConn) SetKeepAlive(keepalive bool) error;
	// func (c *TCPConn) SetKeepAlivePeriod(d time.Duration) error;
	// func (c *TCPConn) SetLinger(sec int) error;
	// func (c *TCPConn) SetNoDelay(noDelay bool) error;
	// func (c *TCPConn) Read(b []byte) (int, error);
	// func (c *TCPConn) ReadFrom(r io.Reader) (int64, error);
	// func (c *TCPConn) Write(b []byte) (int, error);
	// func (c *TCPConn) Close() error;
	// func (c *TCPConn) CloseRead() error;
	// func (c *TCPConn) CloseWrite() error;
	// func (c *TCPConn) File() (f *os.File, err error);

	// UDPConn类型代表UDP网络连接(实现了Conn和PacketConn接口);
	// type UDPConn struct { ... };
	// func DialUDP(net string, laddr, raddr *UDPAddr) (*UDPConn, error);
	// func ListenMulticastUDP(net string, ifi *Interface, gaddr *UDPAddr) (*UDPConn, error);
	// func ListenUDP(net string, laddr *UDPAddr) (*UDPConn, error);
	// func (c *UDPConn) LocalAddr() Addr;
	// func (c *UDPConn) RemoteAddr() Addr;
	// func (c *UDPConn) SetReadBuffer(bytes int) error;
	// func (c *UDPConn) SetWriteBuffer(bytes int) error;
	// func (c *UDPConn) SetDeadline(t time.Time) error;
	// func (c *UDPConn) SetReadDeadline(t time.Time) error;
	// func (c *UDPConn) SetWriteDeadline(t time.Time) error;
	// func (c *UDPConn) Read(b []byte) (int, error);
	// func (c *UDPConn) ReadFrom(b []byte) (int, Addr, error);
	// func (c *UDPConn) ReadFromUDP(b []byte) (n int, addr *UDPAddr, err error);
	// func (c *UDPConn) ReadMsgUDP(b, oob []byte) (n, oobn, flags int, addr *UDPAddr, err error);
	// func (c *UDPConn) Write(b []byte) (int, error);
	// func (c *UDPConn) WriteTo(b []byte, addr Addr) (int, error);
	// func (c *UDPConn) WriteToUDP(b []byte, addr *UDPAddr) (int, error);
	// func (c *UDPConn) WriteMsgUDP(b, oob []byte, addr *UDPAddr) (n, oobn int, err error);
	// func (c *UDPConn) Close() error;
	// func (c *UDPConn) File() (f *os.File, err error);

	// UnixConn类型代表Unix域socket连接(实现了Conn和PacketConn接口);
	// type UnixConn struct { ... };
	// func DialUnix(net string, laddr, raddr *UnixAddr) (*UnixConn, error);
	// func ListenUnixgram(net string, laddr *UnixAddr) (*UnixConn, error);
	// func (c *UnixConn) LocalAddr() Addr;
	// func (c *UnixConn) RemoteAddr() Addr;
	// func (c *UnixConn) SetReadBuffer(bytes int) error;
	// func (c *UnixConn) SetWriteBuffer(bytes int) error;
	// func (c *UnixConn) SetDeadline(t time.Time) error;
	// func (c *UnixConn) SetReadDeadline(t time.Time) error;
	// func (c *UnixConn) SetWriteDeadline(t time.Time) error;
	// func (c *UnixConn) Read(b []byte) (int, error);
	// func (c *UnixConn) ReadFrom(b []byte) (int, Addr, error);
	// func (c *UnixConn) ReadFromUnix(b []byte) (n int, addr *UnixAddr, err error);
	// func (c *UnixConn) ReadMsgUnix(b, oob []byte) (n, oobn, flags int, addr *UnixAddr, err error);
	// func (c *UnixConn) Write(b []byte) (int, error);
	// func (c *UnixConn) WriteTo(b []byte, addr Addr) (n int, err error);
	// func (c *UnixConn) WriteToUnix(b []byte, addr *UnixAddr) (n int, err error);
	// func (c *UnixConn) WriteMsgUnix(b, oob []byte, addr *UnixAddr) (n, oobn int, err error);
	// func (c *UnixConn) Close() error;
	// func (c *UnixConn) CloseRead() error;
	// func (c *UnixConn) CloseWrite() error;
	// func (c *UnixConn) File() (f *os.File, err error);
}

// 两种监听者(struct);
func net_Summary03() {
	// TCP网络的监听者(使用者应尽量使用Listener接口而不是假设网络连接为TCP);
	// type TCPListener struct { ... };
	// func ListenTCP(net string, laddr *TCPAddr) (*TCPListener, error);
	// func (l *TCPListener) Addr() Addr;
	// func (l *TCPListener) SetDeadline(t time.Time) error;
	// func (l *TCPListener) Accept() (Conn, error);
	// func (l *TCPListener) AcceptTCP() (*TCPConn, error);
	// func (l *TCPListener) Close() error;
	// func (l *TCPListener) File() (f *os.File, err error);

	// Unix域scoket的监听者(使用者应尽量使用Listener接口而不是假设网络连接为Unix域scoket);
	// type UnixListener struct { ... };
	// func ListenUnix(net string, laddr *UnixAddr) (*UnixListener, error);
	// func (l *UnixListener) Addr() Addr;
	// func (l *UnixListener) SetDeadline(t time.Time) (err error);
	// func (l *UnixListener) Accept() (c Conn, err error);
	// func (l *UnixListener) AcceptUnix() (*UnixConn, error);
	// func (l *UnixListener) Close() error;
	// func (l *UnixListener) File() (f *os.File, err error);
}

// 各种查询;
func net_Summary04() {
	// 将格式为"host:port"、"[host]:port"或"[ipv6-host%zone]:port"的网络地址分割为host或ipv6-host%zone和port两个部分, Ipv6的文字地址或者主机名必须用方括号括起来, 如"[::1]:80"、"[ipv6-host]:http"、"[ipv6-host%zone]:80";
	// func SplitHostPort(hostport string) (host, port string, err error);
	// 将host和port合并为一个网络地址, 一般格式为"host:port", 如果host含有冒号或百分号, 格式为"[host]:port";
	// func JoinHostPort(host, port string) string;

	// 查询指定网络和服务的(默认)端口;
	// func LookupPort(network, service string) (port int, err error);
	// 查询name的规范DNS名;
	// func LookupCNAME(name string) (cname string, err error);
	// 查询主机的网络地址序列;
	// func LookupHost(host string) (addrs []string, err error);
	// 查询主机的ipv4和ipv6地址序列;
	// func LookupIP(host string) (addrs []IP, err error);
	// 查询某个地址, 返回映射到该地址的主机名序列(本函数和LookupHost不互为反函数);
	// func LookupAddr(addr string) (name []string, err error);
	// 返回指定主机的按Pref字段排好序的DNS MX记录;
	// func LookupMX(name string) (mx []*MX, err error);
	// 返回指定主机的DNS NS记录;
	// func LookupNS(name string) (ns []*NS, err error);
	// 尝试执行指定服务、协议、主机的SRV查询(协议proto为"tcp" 或"udp"), 返回的记录按Priority字段排序, 同一优先度按Weight字段随机排序;
	// func LookupSRV(service, proto, name string) (cname string, addrs []*SRV, err error);
	// 返回指定主机的DNS TXT记录;
	// func LookupTXT(name string) (txt []string, err error);
}

// 最终常用的连接与侦听;
func net_Summary09() {
	// #############################################################连接#############################################################
	// 在网络network上连接地址address, 并返回一个Conn接口, 可用的网络类型有: "tcp"、"tcp4"、"tcp6"、"udp"、"udp4"、"udp6"、"ip"、"ip4"、"ip6"、"unix"、"unixgram"、"unixpacket";
	// func Dial(network, address string) (Conn, error);
	// 类似Dial但采用了超时(timeout参数如果必要可包含名称解析);
	// func DialTimeout(network, address string, timeout time.Duration) (Conn, error);
	// 创建一个内存中的同步、全双工网络连接, 连接的两端都实现了Conn接口, 一端的读取对应另一端的写入, 直接将数据在两端之间作拷贝(没有内部缓冲);
	// func Pipe() (Conn, Conn);

	// 客户端的连接示例;
	// conn, err := net.Dial("tcp", "google.com:80")
	// if err != nil {
	//	   fmt.Println(err)
	//	   return
	// }
	// fmt.Fprintf(conn, "GET / HTTP/1.0\r\n\r\n")
	// status, err := bufio.NewReader(conn).ReadString('\n')

	// #############################################################侦听#############################################################
	// 返回在一个本地网络地址laddr上监听的Listener, 网络类型参数net必须是面向流的网络: "tcp"、"tcp4"、"tcp6"、"unix"或"unixpacket");
	// func Listen(net, laddr string) (Listener, error);

	// 服务器端的侦听示例;
	// ln, err := net.Listen("tcp", ":8080")
	// if err != nil {
	//	   fmt.Println(err)
	//	   return
	// }
	// for {
	//   conn, err := ln.Accept()
	//   if err != nil {
	//	   fmt.Println(err)
	//	   continue
	//   }
	//   go handleConnection(conn)
	// }
}
