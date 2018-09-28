// Package ios GO语言中的文件路径操作函数(path/filepath包);
package ios

import "fmt"
import "os"
import "path/filepath"
import "github.com/kingreatwill/go_study/learning/common"

// filepath;
func filepath01() {
	common.Println("filepath.go")

	fmt.Println(string(filepath.Separator))                                                                              // 路径分隔符, windows下返回\;
	fmt.Println(filepath.IsAbs("c:\\wind\\aa\\bb\\b.txt"))                                                               // 是否绝对路径, true;
	fmt.Println(filepath.IsAbs("http://www.baidu.com/aa.jpg"))                                                           // 是否绝对路径, false;
	fmt.Println(filepath.Abs("."))                                                                                       // (常用)绝对路径, D:\Projects\GoPath\source\demo\syntax\path <nil>;
	fmt.Println(filepath.Rel("c:/windows", "c:/windows/system/"))                                                        // 相对路径, //system <nil>;
	fmt.Println(filepath.SplitList("c:/windows/system/abc.exe"))                                                         // 路径分割, [c:/windows/system/abc.exe];
	fmt.Println(filepath.Split("c:/windows/system/abc.exe"))                                                             // 路径分割, c:/windows/system/ abc.exe;
	fmt.Println(filepath.Join("a", "\\bb\\", "cc", "/d", "e\\", "ff.txt"))                                               // (常用)拼接路径, a\bb\cc\d\e\ff.txt;
	fmt.Println(filepath.FromSlash("c:\\windows\\aa//bb/cc//path.exe"))                                                  // (常用)斜杠('/')替换为路径分隔符, c:\windows\aa\\bb\cc\\path.exe;
	fmt.Println(filepath.ToSlash("c:\\windows\\aa/bb/cc/path.exe"))                                                      // (常用)路径分隔符替换为斜杠('/'), c:/windows/aa/bb/cc/path.exe;
	fmt.Println(filepath.VolumeName("c:\\windows\\"))                                                                    // 获取卷标, c:;
	fmt.Println(filepath.Dir("aa/c\\baa.exe"))                                                                           // (常用)上级路径, aa\c;
	fmt.Println(filepath.Base("c:\\aa\\baa.exe"))                                                                        // (常用)路径最后元素, baa.exe;
	fmt.Println(filepath.Ext("./path.exe"))                                                                              // (常用)文件扩展名, .exe;
	fmt.Println(filepath.Clean("c:\\\\aa/c\\baa.exe"))                                                                   // (常用)最短路径, c:\aa\c\baa.exe;
	fmt.Println(filepath.Clean("aa/c\\baa.exe"))                                                                         // (常用)最短路径, aa\c\baa.exe;
	fmt.Println(filepath.EvalSymlinks("./path.exe"))                                                                     // 可以用来判断文件或文件夹是否存在。 //path.exe <nil>
	fmt.Println(filepath.Match("c:/windows/*/", "c:/windows/system/"))                                                   // 匹配路径, true <nil>;
	fmt.Println(filepath.Glob("c:\\windows\\*.exe"))                                                                     // (常用)匹配文件;
	err1 := filepath.Walk("./src/golangingcloud.service.cluster", func(path string, info os.FileInfo, err error) error { // (常用)遍历目录下的文件夹/文件;
		fmt.Println("File:", path, "IsDir:", info.IsDir(), "size:", info.Size())
		return nil
	})
	if err1 != nil {
		fmt.Println(err1)
	}
	fmt.Println(filepath.HasPrefix("c:\\aa\\bb", "c:\\")) // (不推荐使用)是否包含前缀, true;
}
