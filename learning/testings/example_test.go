// GO语言中的自动化测试(testing包);
// 样本测试;
package testings

import "fmt"

// ~~~~~~~~~~~~~~~~~~~~~~~~~~~~
// 命名规则(单个样本用例);
// ~~~~~~~~~~~~~~~~~~~~~~~~~~~~
// 包;
// func Example() { ... };
// 函数F;
// func ExampleF() { ... };
// 类型T;
// func ExampleT() { ... };
// 类型T上的方法M;
// func ExampleT_M() { ... };
// ~~~~~~~~~~~~~~~~~~~~~~~~~~~~
// 命名规则(多个样本用例);
// ~~~~~~~~~~~~~~~~~~~~~~~~~~~~
// 包;
// func Example_xxx() { ... }
// 函数F;
// func ExampleF_xxx() { ... }
// 类型T;
// func ExampleT_xxx() { ... }
// 类型T上的方法M;
// func ExampleT_M_xxx() { ... }

// 样本测试(指定用例);
// go test -v -run=ExampleF_add1 golangingcloud.service.cluster/learning/testings
// 编译输出:
// === RUN   ExampleF_add1
// --- PASS: ExampleF_add1 (0.00s)
// PASS
// ok  	golangingcloud.service.cluster/learning/testings	0.406s
func ExampleF_add1() {
	fmt.Println(Add(1, 2))
	// Output: 3 <nil>
}

// 样本测试(指定用例);
// go test -v -run=ExampleF_add2 golangingcloud.service.cluster/learning/testings
// 编译输出:
// === RUN   ExampleF_add2
// --- PASS: ExampleF_add2 (0.00s)
// PASS
// ok  	golangingcloud.service.cluster/learning/testings	(cached)
func ExampleF_add2() {
	for i := 1; i <= 5; i++ {
		fmt.Println(Add(1, i))
	}
	// Unordered output:
	// 2 <nil>
	// 6 <nil>
	// 5 <nil>
	// 3 <nil>
	// 4 <nil>
}
