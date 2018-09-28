// Package ios GO语言中的路径操作函数(path包);
package ios

import "fmt"
import "path"
import "github.com/kingreatwill/go_study/learning/common"

// path;
func path01() {
	common.Println("path.go")

	fmt.Println(path.IsAbs("/dev/null"))                              // 是否绝对路径(以/开头?), true;
	fmt.Println(path.IsAbs(`E:\研发人员项目文件`))                            // 是否绝对路径(以/开头?), false(无效);
	fmt.Println(path.IsAbs("E:/研发人员项目文件"))                            // 是否绝对路径(以/开头?), false(无效);
	fmt.Println(path.Split("/dev/static/myfile.css"))                 // 分隔路径+文件名, /dev/static/ myfile.css;
	fmt.Println(path.Split("E:/研发人员项目文件/static/myfile.css"))          // 分隔路径+文件名, E:/研发人员项目文件/static/ myfile.css;
	fmt.Println(path.Join("a", "b", "c"))                             // (常用)拼接路径, a/b/c;
	fmt.Println(path.Join("c:", "aa", "bb", "cc.txt"))                // (常用)拼接路径, c:/aa/bb/cc.txt;
	fmt.Println(path.Dir("/a/b/c/"))                                  // (常用)上级路径, /a/b/c;
	fmt.Println(path.Dir("/a/b/c"))                                   // (常用)上级路径, /a/b;
	fmt.Println(path.Base("/a/b/"))                                   // 路径最后元素, b;
	fmt.Println(path.Base("/a/b"))                                    // 路径最后元素, b;
	fmt.Println(path.Base("http://www.baidu.com/file/aa.jpg"))        // 路径最后元素, aa.jpg;
	fmt.Println(path.Ext("/dev/static/myfile.css"))                   // (常用)文件扩展名, .css;
	fmt.Println(path.Ext(`E:\研发人员项目文件\static\myfile.css`))            // (常用)文件扩展名, .css;
	fmt.Println(path.Clean("a/c/b/.."))                               // (常用)最短路径, a/c;
	fmt.Println(path.Clean("/../a/c"))                                // (常用)最短路径, /a/c;
	isMatch, err := path.Match("c:/windows/*/", "c:/windows/system/") // 匹配路径;
	fmt.Println(isMatch, err)                                         // true <nil>;
}
