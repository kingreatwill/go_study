// GO语言中的HTTP客户端和服务端(net/http包);
package nets

import "fmt"
import "net/http"
import "github.com/kingreatwill/go_study/learning/common"

func IndexHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello World.")
}

// 简单示例;
func nethttp_main() {
	common.Println("net_http.go")

	// 注册路由;
	http.HandleFunc("/", IndexHandler)
	// 开启监听;
	http.ListenAndServe(":8080", nil)
}
