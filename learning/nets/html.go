// GO语言中的转义和解转义HTML文本(html包);
package nets

import "fmt"
import "html"
import "github.com/kingreatwill/go_study/learning/common"

// html01;
func html01() {
	common.Println("html.go")

	_html := `<a href="https://studygolang.com/static/pkgdoc/pkg/runtime.htm?a=1024" target="main">runtime'</a>`

	// 将特定的一些字符转为逸码后的字符实体(如"<"变成"&lt;");
	// 它只会修改五个字符：< > & ' ";
	// func EscapeString(s string) string
	escape := html.EscapeString(_html)
	fmt.Println(html.EscapeString(_html))

	// 将逸码的字符实体如"&lt;"修改为原字符"<";
	// 它会解码一个很大范围内的字符实体, 远比函数EscapeString转码范围大得多;
	// func UnescapeString(s string) string;
	fmt.Println(_html)
	fmt.Println(html.UnescapeString(escape))
}
