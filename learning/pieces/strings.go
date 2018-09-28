// GO语言中的字符串函数(strings包);
package pieces

import "fmt"
import "strings"
import "github.com/kingreatwill/go_study/learning/common"

// strings;
func strings01() {
	common.Println("strings.go")

	// 比较(忽略大小写), ture/false;
	fmt.Println(strings.EqualFold("Go", "go")) // true;
	// 比较(区分大小写), -1=不同, 0=相同;
	fmt.Println(strings.Compare("Go", "go")) // -1;

	// 前缀;
	fmt.Println(strings.HasPrefix("NLT_abc", "NLT")) // true;
	// 后缀;
	fmt.Println(strings.HasSuffix("NLT_abc", "abc")) // true;
	// 包含;
	fmt.Println(strings.Contains("seafood", "foo")) // true;
	fmt.Println(strings.Contains("seafood", "bar")) // false;
	fmt.Println(strings.Contains("seafood", ""))    // true;

	// 统计次数;
	fmt.Println(strings.Count("cheese", "e")) // 3;
	fmt.Println(strings.Count("five", ""))    // 5;

	// 第一出现的位置;
	fmt.Println(strings.Index("NLT_abc", "abc")) // 4;
	fmt.Println(strings.Index("NLT_abc", "aaa")) // -1;
	fmt.Println(strings.Index("我是中国人", "中"))     // 6;

	// 最后出现的位置;
	fmt.Println(strings.LastIndex("go gopher", "go")) // 3;

	// 重复拼接;
	fmt.Println("ba" + strings.Repeat("na", 2))

	// 替换;
	fmt.Println(strings.Replace("oink oink oink", "k", "ky", 2))      // oinky oinky oink;
	fmt.Println(strings.Replace("oink oink oink", "oink", "moo", -1)) // moo moo moo;

	// Trim();
	fmt.Println(strings.Trim(" !!! Achtung! Achtung! !!! ", "! "))        // 去掉头尾!与空格;
	fmt.Println(strings.TrimSpace(" \t\n a lone gopher \n\t\r\n"))        // a lone gopher;
	fmt.Println(strings.TrimLeft(" !!! Achtung! Achtung! !!! ", "! A"))   // 去掉头!/空格/A;
	fmt.Println(strings.TrimPrefix("! Achtung! Achtung! !!! ", "! "))     // 去掉头!与空格;
	fmt.Println(strings.TrimRight(" !!! Achtung! Achtung! !!! ", "! g"))  // 去掉尾!/空格/g;
	fmt.Println(strings.TrimSuffix(" !!! Achtung! Achtung! !!! ", "! g")) // 去掉尾!/空格/g;

	// 分割;
	fmt.Println(strings.Fields("  foo bar  baz   "))                // [foo bar baz];
	fmt.Println(strings.Split("a,b,c", ","))                        // [a b c];
	fmt.Println(strings.Split("a man a plan a canal panama", "a ")) // [ man  plan  canal panama];
	fmt.Println(strings.Split(" xyz ", ""))                         // [  x y z  ];
	fmt.Println(strings.Split("", "Bernardo O'Higgins"))            // [];

	// 连接;
	x := []string{"foo", "bar", "baz"}
	fmt.Println(strings.Join(x, ",")) // foo,bar,baz;

	// 转换;
	s := "Welcome to The WORld of go."
	// 小写;
	fmt.Println(strings.ToLower(s))
	// 大写;
	fmt.Println(strings.ToUpper(s))
	// 标题;
	fmt.Println(strings.ToTitle(s))
	fmt.Println(strings.Title(s))

}
