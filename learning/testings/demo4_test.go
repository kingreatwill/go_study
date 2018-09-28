// GO语言中的自动化测试(testing包);
// Main函数测试;
package testings

import "fmt"
import "os"
import "testing"

// Main函数测试;
// go test -v golangingcloud.service.cluster/learning/testings
func TestMain(m *testing.M) {
	fmt.Println("begin")
	// flag.Parse()
	os.Exit(m.Run())
	fmt.Println("end")
}
