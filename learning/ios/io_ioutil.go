// Package ios GO语言中的IO工具函数(io/ioutil包);
package ios

import "fmt"
import "os"
import "io/ioutil"
import "github.com/kingreatwill/go_study/learning/common"

// ioutil;
func ioutil01() {
	common.Println("io_ioutil.go")

	// 用一个无操作的Close方法包装r返回一个ReadCloser接口;
	// func NopCloser(r io.Reader) io.ReadCloser;

	// (常用)读取文件;
	// func ReadAll(r io.Reader) ([]byte, error);
	file31, err31 := os.Open("E:/研发人员项目文件/GolangingCloud-SVN/GolangingCloud-Domain-Layer/branches/V1.0.1.0/GolangingCloud.Service.Cluster/doc/temp/os_osfile02.txt")
	defer file31.Close()
	if err31 != nil {
		fmt.Println(err31)
	}
	bytes, _ := ioutil.ReadAll(file31)
	fmt.Println(string(bytes))

	// (常用)读取文件;
	// func ReadFile(filename string) ([]byte, error);
	data, _ := ioutil.ReadFile("E:/研发人员项目文件/GolangingCloud-SVN/GolangingCloud-Domain-Layer/branches/V1.0.1.0/GolangingCloud.Service.Cluster/doc/temp/os_osfile02.txt")
	fmt.Println(string(data))

	// (常用)写入文件(文件不存在则创建文件, 文件存在则清空再写入文件);
	// func WriteFile(filename string, data []byte, perm os.FileMode) error;
	ioutil.WriteFile("E:/研发人员项目文件/GolangingCloud-SVN/GolangingCloud-Domain-Layer/branches/V1.0.1.0/GolangingCloud.Service.Cluster/doc/temp/io_ioutil01.txt", []byte("hello,world"), 0660)

	// (常用)遍历目录里的内容(文件夹/文件);
	// func ReadDir(dirname string) ([]os.FileInfo, error);
	fis, _ := ioutil.ReadDir("E:/研发人员项目文件/GolangingCloud-SVN/GolangingCloud-Domain-Layer/branches/V1.0.1.0/GolangingCloud.Service.Cluster/doc/temp/")
	for i, fi := range fis {
		fmt.Println(i, fi.Name())
	}

	// 临时文件/目录;
	// func TempDir(dir, prefix string) (name string, err error);
	// func TempFile(dir, prefix string) (f *os.File, err error);
}
