// Package ios GO语言中的不依赖平台的操作系统函数(os包);
package ios

import "fmt"
import "os"
import "path/filepath"
import "strings"
import "github.com/kingreatwill/go_study/learning/common"

// os;
func osdir01() {
	common.Println("os.dir.go")

	fmt.Println(os.Hostname())                 // 返回内核提供的主机名, 20161026-133300 <nil>;
	fmt.Println(os.Getpagesize())              // 返回底层的系统内存页的尺寸, 4096;
	fmt.Println(os.Environ())                  // 返回环境变量的格式为"key=value"的字符串的切片拷贝, [GOHOSTOS=windows CGO_ENABLED=1 NUMBER_OF_PROCESSORS=4 CXX=g++ GOTOOLDIR=c:\go\pkg\tool\windows_amd64 windows_tracing_flags=3 COMPUTERNAME=20161026-133300 LITEIDE_APP_PATH=C:/Program Files/liteidex33.2.windows-qt4/liteide/bin GOCACHE=C:\Users\Administrator\AppData\Local\go-build GOBIN= LITEIDE_TOOL_PATH=C:/Program Files/liteidex33.2.windows-qt4/liteide/bin LITEIDE_PLUGIN_PATH=C:/Program Files/liteidex33.2.windows-qt4/liteide/lib/liteide/plugins USERDOMAIN=20161026-133300 GOROOT=c:\go LITEIDE_EXECOPT=/C MSMPI_BIN=C:\Program Files\Microsoft MPI\Bin\ SystemRoot=C:\Windows ComSpec=C:\Windows\system32\cmd.exe GORACE= APPDATA=C:\Users\Administrator\AppData\Roaming LITEIDE_EXEC=C:\Windows\system32\cmd.exe USERNAME=Administrator PATHEXT=.COM;.EXE;.BAT;.CMD;.VBS;.VBE;.JS;.JSE;.WSF;.WSH;.MSC GOTMPDIR= PUBLIC=C:\Users\Public VS120COMNTOOLS=D:\Program Files (x86)\Microsoft Visual Studio 12.0\Common7\Tools\ GOGCCFLAGS=-m64 -mthreads -fno-caret-diagnostics -Qunused-arguments -fmessage-length=0 -fdebug-prefix-map=C:\Users\ADMINI~1\AppData\Local\Temp\go-build306983862=/tmp/go-build -gno-record-gcc-switches OS=Windows_NT CommonProgramFiles=C:\Program Files\Common Files CommonProgramFiles(x86)=C:\Program Files (x86)\Common Files Path=c:\mingw64\bin;c:\go\bin;C:\Program Files\Microsoft MPI\Bin\;C:\Windows\system32;C:\Windows;C:\Windows\System32\Wbem;C:\Windows\System32\WindowsPowerShell\v1.0\;C:\Program Files\Microsoft SQL Server\110\Tools\Binn\;C:\Program Files\Microsoft SQL Server\120\DTS\Binn\;C:\Program Files\Microsoft SQL Server\Client SDK\ODBC\110\Tools\Binn\;C:\Program Files (x86)\Microsoft SQL Server\120\Tools\Binn\;C:\Program Files\Microsoft SQL Server\120\Tools\Binn\;C:\Program Files (x86)\Microsoft SQL Server\120\Tools\Binn\ManagementStudio\;C:\Program Files (x86)\Microsoft SQL Server\120\DTS\Binn\;D:\Program Files\TortoiseSVN\bin;C:\Program Files\OpenVPN\bin;C:\Program Files\nodejs\;C:\Program Files\Microsoft SQL Server\130\Tools\Binn\;C:\Program Files\dotnet\;C:\Program Files\Git\cmd;C:\Go\bin;C:\Program Files (x86)\Microsoft VS Code\bin;C:\Users\Administrator\AppData\Roaming\npm;C:/Program Files/liteidex33.2.windows-qt4/liteide/bin;c:/go/bin;c:/go/bin/windows_amd64;C:/Users/Administrator/go/bin;C:/Users/Administrator/go/bin/windows_amd64;E:/研发人员项目文件/GolangingCloud-SVN/GolangingCloud-Domain-Layer/branches/V1.0.1.0/GolangingCloud.Service.Cluster/bin;E:/研发人员项目文件/GolangingCloud-SVN/GolangingCloud-Domain-Layer/branches/V1.0.1.0/GolangingCloud.Service.Cluster/bin/windows_amd64; SystemDrive=C: SHFBROOT=D:\Program Files (x86)\EWSoftware\Sandcastle Help File Builder\ DXROOT=D:\Program Files (x86)\Sandcastle\ PROCESSOR_REVISION=5e03 ALLUSERSPROFILE=C:\ProgramData HOMEPATH=\Users\Administrator TEMP=C:\Users\ADMINI~1\AppData\Local\Temp windir=C:\Windows PSModulePath=C:\Windows\system32\WindowsPowerShell\v1.0\Modules\;C:\Program Files (x86)\Microsoft SQL Server\120\Tools\PowerShell\Modules\ CGO_CPPFLAGS= CGO_CFLAGS=-g -O2 LITEIDE_GDB=gdb64 FSHARPINSTALLDIR=C:\Program Files (x86)\Microsoft SDKs\F#\4.1\Framework\v4.0\ LITEIDE_TERMARGS= GCCGO=gccgo CommonProgramW6432=C:\Program Files\Common Files PKG_CONFIG=pkg-config ProgramFiles(x86)=C:\Program Files (x86) LITEIDE_RES_PATH=C:/Program Files/liteidex33.2.windows-qt4/liteide/share/liteide PROCESSOR_LEVEL=6 CGO_LDFLAGS=-g -O2 GOEXE=.exe LITEIDE_ROOT_PATH=C:/Program Files/liteidex33.2.windows-qt4/liteide USERPROFILE=C:\Users\Administrator FP_NO_HOST_CHECK=NO PROCESSOR_IDENTIFIER=Intel64 Family 6 Model 94 Stepping 3, GenuineIntel ProgramFiles=C:\Program Files CGO_FFLAGS=-g -O2 ProgramData=C:\ProgramData windows_tracing_logfile=C:\BVTBin\Tests\installpackage\csilogfile.log LITEIDE_TERM=C:\Windows\system32\cmd.exe LITEIDE_MAKE=mingw32-make GOPATH=C:\Users\Administrator\go;E:\研发人员项目文件\GolangingCloud-SVN\GolangingCloud-Domain-Layer\branches\V1.0.1.0\GolangingCloud.Service.Cluster LOGONSERVER=\\20161026-133300 LOCALAPPDATA=C:\Users\Administrator\AppData\Local CC=gcc GOHOSTARCH=amd64 TMP=C:\Users\ADMINI~1\AppData\Local\Temp GOOS=windows PROCESSOR_ARCHITECTURE=AMD64 ProgramW6432=C:\Program Files HOMEDRIVE=C: SESSIONNAME=Console GOARCH=amd64 CGO_CXXFLAGS=-g -O2];
	fmt.Println(os.Getenv("GOHOSTARCH"))       // 检索并返回名为key的环境变量的值, amd64;
	fmt.Println(os.Setenv("KY", "testing...")) // 设置名为key的环境变量(临时);

	// os.Clearenv()                           // 删除所有环境变量;
	// os.Exit(0)                              // 程序退出, 状态码0=成功，非0=出错;
	// os.Expand(...)                          // 回调函数替换${var}或$var;
	// os.ExpandEnv(...)                       // 回调函数替换环境变量${var}或$var值;

	fmt.Println(os.Getuid())    // 返回调用者的用户ID;
	fmt.Println(os.Geteuid())   // 返回调用者的有效用户ID;
	fmt.Println(os.Getgid())    // 返回调用者的组ID;
	fmt.Println(os.Getegid())   // 返回调用者的有效组ID;
	fmt.Println(os.Getgroups()) // 返回调用者所属的所有用户组的组ID;
	fmt.Println(os.Getpid())    // 返回调用者所在进程的进程ID;
	fmt.Println(os.Getppid())   // 返回调用者所在进程的父进程的进程ID;

	// (常用)FileInfo的方法;
	fi1, err1 := os.Lstat("E:/研发人员项目文件/GolangingCloud-SVN/GolangingCloud-Domain-Layer/branches/V1.0.1.0/GolangingCloud.Service.Cluster/doc/temp/os_osdir01.txt")
	if err1 == nil {
		fmt.Println(fi1.Name())    // 文件的名字;
		fmt.Println(fi1.Size())    // 普通文件返回值表示其大小, 其他文件的返回值含义各系统不同;
		fmt.Println(fi1.Mode())    // 文件的模式位;
		fmt.Println(fi1.ModTime()) // 文件的修改时间;
		fmt.Println(fi1.IsDir())   // 是否目录;
		fmt.Println(fi1.Sys())     // 底层数据来源(可以返回nil);
	}

	// (常用)文件的模式位FileMode的方法;
	fi2, err2 := os.Lstat("E:/研发人员项目文件/GolangingCloud-SVN/GolangingCloud-Domain-Layer/branches/V1.0.1.0/GolangingCloud.Service.Cluster/doc/temp/os_osdir01.txt")
	if err2 == nil {
		fm := fi2.Mode()
		fmt.Println(fm.IsDir())     // 是否目录;
		fmt.Println(fm.IsRegular()) // 是否普通文件;
		fmt.Println(fm.Perm())      // 返回Unix权限位;
		fmt.Println(fm.String())    // 返回Unix权限位;
	}

	// 某些判断;
	fmt.Println(os.IsPathSeparator('/')) // 字符是否路径分隔符, true;
	fmt.Println(os.IsPathSeparator('|')) // 字符是否路径分隔符, false;
	// os.IsExist(err error) bool        // 返回一个布尔值说明该错误是否表示一个文件或目录已经存在;
	// os.IsNotExist(err error) bool     // 返回一个布尔值说明该错误是否表示一个文件或目录不存在;
	// os.IsPermission(err error) bool   // 返回一个布尔值说明该错误是否表示因权限不足要求被拒绝;

	// (常用)某些修改;
	fmt.Println(os.Getwd())                                                                                                                   // 当前工作目录;
	fmt.Println(os.Chdir("E:/研发人员项目文件/GolangingCloud-SVN/GolangingCloud-Domain-Layer/branches/V1.0.1.0/GolangingCloud.Service.Cluster/doc/")) // 切换当前工作目录;
	fmt.Println(os.Getwd())

	// os.Chmod();   // 修改文件对象的mode;
	// os.Chown();   // 修改文件对象的用户ID和组ID;
	// os.Lchown();  // 修改文件对象的用户ID和组ID;
	// os.Chtimes(); // 修改文件对象的访问时间和修改时间;

	// (常用)存在目录/文件?
	_, err20 := os.Stat("5089/895410/3405089")
	fmt.Println(err20)

	// (常用)创建目录(指定权限);
	os.Chdir("E:/研发人员项目文件/GolangingCloud-SVN/GolangingCloud-Domain-Layer/branches/V1.0.1.0/GolangingCloud.Service.Cluster/") // 切换当前工作目录;
	root, _ := os.Getwd()                                                                                                    // 当前工作目录;
	sp := string(filepath.Separator)                                                                                         // 路径分隔符;
	// path1 := strings.Join([]string{root, "doc", "test"}, sp)                                                              // 拼接路径();
	// err21 := os.Mkdir(path1, os.ModePerm)                                                                                 // 创建目录(指定权限), 上级目录必须已存在, 否则出错;
	path2 := strings.Join([]string{root, "doc", "test", "a", "b", "c"}, sp) // 拼接路径();
	err22 := os.MkdirAll(path2, os.ModePerm)                                // (推荐)创建目录(指定权限);
	if err22 != nil {
		fmt.Println(err22)
	}

	// (常用)重命名或移动(目录/文件), 目标重名则出错;
	err23 := os.Rename(path2, strings.Join([]string{root, "doc", "test", "a", "c"}, sp))
	if err23 != nil {
		fmt.Println(err23)
	}

	// 修文件的大小;
	// func Truncate(name string, size int64) error;

	// (常用)删除目录/文件(不含下级目录);
	//  err24 := os.Remove(strings.Join([]string{root, "doc", "test", "a", "b"}, sp))
	//	if err24 != nil {
	//		fmt.Println(err24)
	//	}
	// (常用)删除目录/文件(含下级目录, 推荐使用);
	err25 := os.RemoveAll(strings.Join([]string{root, "doc", "test", "a"}, sp))
	if err25 != nil {
		fmt.Println(err25)
	}

	// 是否同一个文件?
	// func SameFile(fi1, fi2 FileInfo) bool;

	// 临时文件的默认目录;
	fmt.Println(os.TempDir())

	// 创建一个名为newname指向oldname的符号链接;
	// func Symlink(oldname, newname string) error;
	// 创建一个名为newname指向oldname的硬链接;
	// func Link(oldname, newname string) error;
	// 获取符号链接文件指向的文件的路径;
	// func Readlink(name string) (string, error);
}
