// GO语言中的解析URL(net/url包);
package nets

import "fmt"
import "net/url"
import "github.com/kingreatwill/go_study/learning/common"

// neturl01;
func neturl01() {
	common.Println("net_url.go")

	// 对s进行转码使之可以安全的用在URL里;
	// func QueryEscape(s string) string;
	s := `http://www.bing.com/search?q=dotnet`
	escape := url.QueryEscape(s)
	fmt.Println(escape)

	// 将QueryEscape转码的字符串还原;
	// func QueryUnescape(s string) (string, error);
	unescape, _ := url.QueryUnescape(escape)
	fmt.Println(s)
	fmt.Println(unescape)
}

// neturl02;
func neturl02() {

	// type URL struct {
	//    Scheme   string
	//    Opaque   string    // 编码后的不透明数据
	//    User     *Userinfo // 用户名和密码信息
	//    Host     string    // host或host:port
	//    Path     string
	//    RawQuery string // 编码后的查询字符串，没有'?'
	//    Fragment string // 引用的片段（文档位置），没有'#'
	// }

	// 格式1: Scheme://[User@]Host/Path[?RawQuery][#Fragment];
	// 格式2: Scheme:Opaque[?RawQuery][#Fragment];

	// 解析rawurl为一个URL结构体;
	// func Parse(rawurl string) (url *URL, err error);
	// 解析rawurl为一个URL结构体;
	// func ParseRequestURI(rawurl string) (url *URL, err error);
	u, _ := url.Parse("http://bing.com/search?q=dotnet#d1")
	fmt.Println("Scheme:", u.Scheme)     // http;
	fmt.Println("Opaque:", u.Opaque)     // ;
	fmt.Println("User:", u.User)         // ;
	fmt.Println("Host:", u.Host)         // bing.com;
	fmt.Println("Path:", u.Path)         // /search;
	fmt.Println("RawQuery:", u.RawQuery) // q=dotnet;
	fmt.Println("Fragment:", u.Fragment) // d1;

	// 是否为绝对URL;
	// func (u *URL) IsAbs() bool;
	fmt.Println("IsAbs:", u.IsAbs()) // true;

	// 解析RawQuery字段并返回其表示的Values类型键值对;
	// func (u *URL) Query() Values;
	q := u.Query()
	q.Set("q", "golang")
	u.RawQuery = q.Encode()
	u.Fragment = "d2"

	// 返回编码好的path?query或opaque?query字符串;
	// func (u *URL) RequestURI() string;
	fmt.Println("RequestURI:", u.RequestURI()) // /search?q=dotnet;

	// 将URL重构为一个合法URL字符串;
	// func (u *URL) String() string;
	fmt.Println("String:", u.String())

	// func (u *URL) Parse(ref string) (*URL, error);
	// func (u *URL) ResolveReference(ref *URL) *URL;
}

// neturl03;
func neturl03() {
	// Values将建映射到值的列表, 用于查询的参数和表单的属性;
	// type Values map[string][]string;

	// 解析一个URL编码的查询字符串, 并返回可以表示该查询的Values类型的字典;
	// func ParseQuery(query string) (m Values, err error);
	// 获取key对应的值集的第一个值;
	// func (v Values) Get(key string) string;
	// 将key对应的值集设为只有value, 它会替换掉已有的值集;
	// func (v Values) Set(key, value string);
	// 将value添加到key关联的值集里原有的值的后面;
	// func (v Values) Add(key, value string);
	// 删除key关联的值集;
	// func (v Values) Del(key string);
	// 将v编码为url编码格式("bar=baz&foo=quux"), 编码时会以键进行排序;
	// func (v Values) Encode() string;

	v := url.Values{}
	v.Set("name", "Ava")
	v.Add("friend", "Jess")
	v.Add("friend", "Sarah")
	v.Add("friend", "Zoe")

	fmt.Println(v.Get("name"))   // Ava;
	fmt.Println(v.Get("friend")) // Jess;
	fmt.Println(v["friend"])     // [Jess Sarah Zoe];
	fmt.Println(v.Encode())      // friend=Jess&friend=Sarah&friend=Zoe&name=Ava;
}
