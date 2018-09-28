// GO语言中的SHA1哈希算法(crypto/sha1包);
// 校验和算法;
package crypto_hashes

import "fmt"
import "crypto/sha1"
import "encoding/hex"
import "github.com/kingreatwill/go_study/learning/common"

// SHA1哈希算法;
func sha1_01() {
	common.Println("crypto_sha1.go")

	data := []byte("Hello World.")
	r := sha1.Sum(data)
	fmt.Println("SHA1:", r[:])                     // 返回结果为整数类型切片;
	fmt.Println(string(r[:]))                      // 乱码;
	fmt.Println("SHA1:", hex.EncodeToString(r[:])) // (一般性)转为十六进制编码;
}

// SHA1哈希算法;
func sha1_02() {
	data := []byte("Hello World.")
	h := sha1.New()
	h.Write(data)
	fmt.Println("SHA1:", h.Sum(nil))                     // 返回结果为整数类型切片;
	fmt.Println(string(h.Sum(nil)))                      // 乱码;
	fmt.Println("SHA1:", hex.EncodeToString(h.Sum(nil))) // (一般性)转为十六进制编码;
}
