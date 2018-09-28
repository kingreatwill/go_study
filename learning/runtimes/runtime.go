// GO语言中的go运行时环境(runtime包);
package runtimes

import "fmt"
import "runtime"
import "github.com/kingreatwill/go_study/learning/common"

// runtime01;
func runtime01() {
	common.Println("runtime.go")

	// GOARCH/GOOS;
	fmt.Println("GOARCH:", runtime.GOARCH) // amd64;
	fmt.Println("GOOS:", runtime.GOOS)     // windows;

	// 返回Go的根目录(如果存在GOROOT环境变量, 返回该变量的值, 否则返回创建Go时的根目录);
	// func GOROOT() string;
	// 返回Go的版本字符串(它要么是递交的hash和创建时的日期, 要么是发行标签如"go1.3");
	// func Version() string;
	// 返回本地机器的逻辑CPU个数;
	// func NumCPU() int;
	fmt.Println("GOROOT:", runtime.GOROOT()) // c:\go;
	fmt.Println("Version:", runtime.Version())
	fmt.Println("NumCPU:", runtime.NumCPU())

	// 设置可同时执行的最大CPU数, 并返回先前的设置(本函数在调度程序优化后会去掉);
	// func GOMAXPROCS(n int) int;
	// 执行一次垃圾回收;
	// func GC();
	// 将x的终止器设置为f;
	// func SetFinalizer(x, f interface{});

	// 将内存申请和分配的统计信息填写进m;
	// func ReadMemStats(m *MemStats);
	memStats := new(runtime.MemStats)
	runtime.ReadMemStats(memStats)
	fmt.Println("ReadMemStats:", memStats)

	// 返回当前内存profile中的记录数n;
	// func MemProfile(p []MemProfileRecord, inuseZero bool) (n int, ok bool);
	// 设置CPU profile记录的速率为平均每秒hz次;
	// func SetCPUProfileRate(hz int);
	// 返回二进制CPU profile堆栈跟踪数据的下一个chunk, 函数会阻塞直到该数据可用;
	// func CPUProfile() []byte;
	// 执行一个断点陷阱;
	// func Breakpoint();
	// 将调用其的go程的调用栈踪迹格式化后写入到buf中并返回写入的字节数;
	// func Stack(buf []byte, all bool) int;
	// 报告当前go程调用栈所执行的函数的文件和行号信息;
	// func Caller(skip int) (pc uintptr, file string, line int, ok bool);
	// 把当前go程调用栈上的调用栈标识符填入切片pc中, 返回写入到pc中的项数;
	// func Callers(skip int, pc []uintptr) int;

	// 返回当前进程执行的cgo调用次数;
	// func NumCgoCall() int64;
	// 返回当前存在的Go程数;
	// func NumGoroutine() int;
	fmt.Println("NumCgoCall:", runtime.NumCgoCall())
	fmt.Println("NumGoroutine:", runtime.NumGoroutine())

	// 终止调用它的go程, 其它go程不会受影响, Goexit会在终止该go程前执行所有defer的函数;
	// func Goexit();
	// 使当前go程放弃处理器, 以让其它go程运行, 它不会挂起当前go程, 因此当前go程未来会恢复执行;
	// func Gosched();

	// 将调用的go程绑定到它当前所在的操作系统线程;
	// func LockOSThread();
	// 将调用的go程解除和它绑定的操作系统线程;
	// func UnlockOSThread();

	// 返回活跃go程的堆栈profile中的记录个数;
	// func GoroutineProfile(p []StackRecord) (n int, ok bool);
	// 返回线程创建profile中的记录个数;
	// func ThreadCreateProfile(p []StackRecord) (n int, ok bool);

	// 控制阻塞profile记录go程阻塞事件的采样频率;
	// func SetBlockProfileRate(rate int);
	// 返回当前阻塞profile中的记录个数;
	// func BlockProfile(p []BlockProfileRecord) (n int, ok bool);
}
