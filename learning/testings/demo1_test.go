// GO语言中的自动化测试(testing包);
// 单元测试;
package testings

import "testing"

// 单元测试(包下全部);
// go test golangingcloud.service.cluster/learning/testings
// 编译输出:
// --- FAIL: TestAdd2 (0.00s)
//	 demo_test.go:14: 测试没通过.
// FAIL
// FAIL	golangingcloud.service.cluster/learning/testings	0.299s

// 单元测试(指定用例);
// go test -run=TestAdd1 golangingcloud.service.cluster/learning/testings
// 编译输出:
// ok  	golangingcloud.service.cluster/learning/testings	0.336s
func TestAdd1(t *testing.T) {
	if i, e := Add(1, 1); i != 2 || e != nil {
		t.Error("测试没通过.")
	} else {
		t.Log("测试通过.")
	}
}

// 单元测试(包下全部);
// go test -v golangingcloud.service.cluster/learning/testings

// 编译输出:
// === RUN   TestAdd1
// --- PASS: TestAdd1 (0.00s)
//	 demo_test.go:9: 测试通过.
// === RUN   TestAdd2
// --- FAIL: TestAdd2 (0.00s)
//	 demo_test.go:14: 测试没通过.
// FAIL
// FAIL	golangingcloud.service.cluster/learning/testings	0.319s

// 单元测试(指定用例);
// go test -v -run=TestAdd2 golangingcloud.service.cluster/learning/testings
// 编译输出:
// === RUN   TestAdd2
// --- FAIL: TestAdd2 (0.00s)
//	 demo1_test.go:41: 测试没通过.
// FAIL
// FAIL	golangingcloud.service.cluster/learning/testings	0.355s
// func TestAdd2(t *testing.T) {
//	 t.Error("测试没通过.")
// }
