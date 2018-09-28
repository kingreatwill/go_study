// GO语言中的自动化测试(testing包);
// 基准测试;
package testings

import "testing"

// 基准测试(指定用例);
// go test -bench=BenchmarkAdd1 golangingcloud.service.cluster/learning/testings
// 编译输出:
// goos: windows
// goarch: amd64
// pkg: golangingcloud.service.cluster/learning/testings
// BenchmarkAdd1-4(用例)   	2000000000(次)	         0.34 ns/op(每次执行平均时间)
// PASS
// ok  	golangingcloud.service.cluster/learning/testings	1.131s(总执行时间)
func BenchmarkAdd1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Add(4, 5)
	}
}

// 基准测试(包下全部);
// go test -bench=".*" golangingcloud.service.cluster/learning/testings
// 编译输出:
// goos: windows
// goarch: amd64
// pkg: golangingcloud.service.cluster/learning/testings
// BenchmarkAdd1-4(用例)   	2000000000(次)	         0.33 ns/op(每次执行平均时间)
// BenchmarkAdd2-4(用例)   	2000000000(次)	         0.34 ns/op(每次执行平均时间)
// PASS
// ok  	golangingcloud.service.cluster/learning/testings	1.821s(总执行时间)
func BenchmarkAdd2(b *testing.B) {
	// 一些初始化的工作, 如读取文件数据/数据库连接之类, 这样不影响我们测试函数本身的性能;
	// ......
	// 重置定时器;
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Add(1, i)
	}
}
