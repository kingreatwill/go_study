// GO语言中的公共变量的标准接口(expvar包);
// 对这些公共变量的读写操作都是原子级的;
// 为了增加HTTP处理器, 本包注册了如下变量;
// cmdline   os.Args
// memstats  runtime.Memstats
// 导入本包;
// import _ "expvar";
// 具体可参考: http://blog.studygolang.com/2017/06/expvar-in-action/
// 英文版: https://orangetux.nl/post/expvar_in_action/
package nets

import "fmt"
import "expvar"
import "net/http"
import "github.com/kingreatwill/go_study/learning/common"

var visits = expvar.NewInt("visits")

func handler01(w http.ResponseWriter, r *http.Request) {
	visits.Add(1)
	fmt.Println(r.URL.Path[1:])
}

// expvar01;
func expvar01() {
	common.Println("expvar.go")

	// http.HandleFunc("/", handler01)
	// http.ListenAndServe(":8080", nil)

	// 访问以下路径, 即可返回服务监控数据json;
	// http://localhost:8080/debug/vars
}
