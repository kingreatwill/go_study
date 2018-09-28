// Package ios GO语言中的不依赖平台的操作系统函数(os/user包);
package ios

import "fmt"
import "os/user"
import "github.com/kingreatwill/go_study/learning/common"

// os;
func osuser01() {
	common.Println("os.user.go")

	// type User struct {
	//    Uid      string // 用户ID
	//    Gid      string // 初级组ID
	//    Username string
	//    Name     string
	//    HomeDir  string
	//}

	// 返回当前的用户帐户;
	u, _ := user.Current()
	fmt.Println(u)                       // &{S-1-5-21-3826313030-1448048710-1277678651-500 S-1-5-21-3826313030-1448048710-1277678651-513 20161026-133300\Administrator  C:\Users\Administrator} <nil>;
	fmt.Println("Uid:", u.Uid)           // Uid: S-1-5-21-3826313030-1448048710-1277678651-500;
	fmt.Println("Gid:", u.Gid)           // Gid: S-1-5-21-3826313030-1448048710-1277678651-513;
	fmt.Println("Username:", u.Username) // Username: 20161026-133300\Administrator;
	fmt.Println("Name:", u.Name)         // Name: ;
	fmt.Println("HomeDir:", u.HomeDir)   // HomeDir: C:\Users\Administrator;

	// 根据用户名查询用户;
	// func Lookup(username string) (*User, error);
	u2, _ := user.Lookup("20161026-133300\\Administrator")
	fmt.Println("HomeDir:", u2.HomeDir)

	// 根据用户ID查询用户;
	// func LookupId(uid string) (*User, error);
	u3, _ := user.LookupId("S-1-5-21-3826313030-1448048710-1277678651-500")
	fmt.Println("HomeDir:", u3.HomeDir)
}
