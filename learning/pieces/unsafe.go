// GO语言中的unsafe(unsafe包);
// unsafe包用于Go编译器, 而不是Go运行时;
// 使用unsafe作为程序包名称只是让你在使用此包是更加小心;
// 使用unsafe.Pointer并不总是一个坏主意, 有时我们必须使用它;
// Golang的类型系统是为了安全和效率而设计的, 但是在Go类型系统中, 安全性比效率更重要, 通常Go是高效的, 但有时安全真的会导致Go程序效率低下;
// unsafe包用于有经验的程序员通过安全地绕过Go类型系统的安全性来消除这些低效;
// unsafe包可能被滥用并且是危险的;
// 可参考: https://studygolang.com/articles/9446
// 可参考: http://www.cnblogs.com/golove/p/5909968.html
package pieces
