// Package ios GO语言中的不依赖平台的操作系统函数(os包);
package ios

import "fmt"
import "os"
import "path/filepath"
import "strings"
import "github.com/kingreatwill/go_study/learning/common"

// os;
func osfile01() {
	common.Println("os.file.go")

	// func OpenFile(name string, flag int, perm FileMode) (file *File, err error);
	// func NewFile(fd uintptr, name string) *File;
	// func Pipe() (r *File, w *File, err error);

	// (常用)创建文件(任何人都可读写不可执行, 文件描述符O_RDWR模式);
	// func Create(name string) (file *File, err error);
	root, _ := os.Getwd()                                                       // 当前工作目录;
	sp := string(filepath.Separator)                                            // 路径分隔符;
	path1 := strings.Join([]string{root, "doc", "temp", "os_osfile01.txt"}, sp) // 拼接路径();
	file1, err11 := os.Create(path1)
	defer file1.Close()
	if err11 != nil {
		fmt.Println(err11)
	}

	// (常用)打开文件(文件描述符O_RDONLY模式);
	// func Open(name string) (file *File, err error);
	file2, err12 := os.Open("E:/研发人员项目文件/GolangingCloud-SVN/GolangingCloud-Domain-Layer/branches/V1.0.1.0/GolangingCloud.Service.Cluster/doc/temp/os_osfile02.txt")
	defer file2.Close()
	if err12 != nil {
		fmt.Println(err12)
	}

	// func (f *File) Name() string;
	// func (f *File) Stat() (fi FileInfo, err error);
	// func (f *File) Fd() uintptr;
	// func (f *File) Chdir() error;
	// func (f *File) Chmod(mode FileMode) error;
	// func (f *File) Chown(uid, gid int) error;
	// func (f *File) Truncate(size int64) error;

	// (常用)遍历目录里的内容(文件夹/文件);
	// func (f *File) Readdir(n int) (fi []FileInfo, err error);
	file3, _ := os.Open("E:/研发人员项目文件/GolangingCloud-SVN/GolangingCloud-Domain-Layer/branches/V1.0.1.0/GolangingCloud.Service.Cluster/doc/temp/")
	defer file3.Close()
	fis, err21 := file3.Readdir(0)
	if err21 != nil {
		fmt.Println(err21)
	}
	for i, fi := range fis {
		fmt.Println(i, fi.Name())
	}

	// (常用)遍历目录里的内容(文件夹/文件, 仅要名字字符串);
	// func (f *File) Readdirnames(n int) (names []string, err error);
	file4, _ := os.Open("E:/研发人员项目文件/GolangingCloud-SVN/GolangingCloud-Domain-Layer/branches/V1.0.1.0/GolangingCloud.Service.Cluster/doc/temp/")
	defer file4.Close()
	names, err22 := file4.Readdirnames(0)
	if err22 != nil {
		fmt.Println(err22)
	}
	for i, name := range names {
		fmt.Println(i, name, "Readdirnames")
	}

	// (常用)读取文件;
	// func (f *File) Read(b []byte) (n int, err error);
	// func (f *File) ReadAt(b []byte, off int64) (n int, err error);
	file31, err31 := os.Open("E:/研发人员项目文件/GolangingCloud-SVN/GolangingCloud-Domain-Layer/branches/V1.0.1.0/GolangingCloud.Service.Cluster/doc/temp/os_osfile02.txt")
	defer file31.Close()
	if err31 != nil {
		fmt.Println(err31)
	}
	bytes := make([]byte, 32)
	for {
		n, err32 := file31.Read(bytes)
		if n > 0 {
			fmt.Println(string(bytes[:n]))
		}
		if err32 != nil {
			fmt.Println(err32)
			break
		}
	}

	// (常用)创建+写文件;
	// func (f *File) Write(b []byte) (n int, err error);
	// func (f *File) WriteString(s string) (ret int, err error);
	// func (f *File) WriteAt(b []byte, off int64) (n int, err error);
	root41, _ := os.Getwd()
	path41 := strings.Join([]string{root41, "doc", "temp", "os_osfile02.txt"}, string(filepath.Separator)) // 拼接路径();
	file41, err41 := os.Create(path41)
	defer file41.Close()
	if err41 != nil {
		fmt.Println(err41)
	}
	file41.WriteString("Gin is a HTTP web framework written in Go (Golang). It features a Martini-like API with much better performance -- up to 40 times faster. If you need smashing performance, get yourself some Gin.")
	// file41.Write([]byte("Gin is a HTTP web framework written in Go (Golang). It features a Martini-like API with much better performance -- up to 40 times faster. If you need smashing performance, get yourself some Gin."))
	// file41.Sync()

	// 设置下次读/写的位置;
	// func (f *File) Seek(offset int64, whence int) (ret int64, err error);

	// (常用)创建文件/续写文件(文件不存在则创建文件, 如存在则续写文件);
	file51, err51 := os.OpenFile("E:/研发人员项目文件/GolangingCloud-SVN/GolangingCloud-Domain-Layer/branches/V1.0.1.0/GolangingCloud.Service.Cluster/doc/temp/os_osfile03.txt", os.O_CREATE|os.O_APPEND|os.O_RDWR, 0660)
	defer file51.Close()
	if err51 != nil {
		fmt.Println(err51)
	}
	file51.WriteString("Gin is a HTTP web framework written in Go (Golang). It features a Martini-like API with much better performance -- up to 40 times faster. If you need smashing performance, get yourself some Gin.")
	file51.Sync()
}
