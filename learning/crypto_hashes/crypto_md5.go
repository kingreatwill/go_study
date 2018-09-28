// GO语言中的MD5哈希算法(crypto/md5包);
// 校验和算法;
package crypto_hashes

import "fmt"
import "crypto/md5"
import "encoding/hex"
import "github.com/kingreatwill/go_study/learning/common"

// MD5哈希算法;
func md5_01() {
	common.Println("crypto_sha1.go")

	data := []byte("Hello World.")
	r := md5.Sum(data)
	fmt.Println("MD5:", r[:])                     // 返回结果为整数类型切片;
	fmt.Println(string(r[:]))                     // 乱码;
	fmt.Println("MD5:", hex.EncodeToString(r[:])) // (一般性)转为十六进制编码;
}

// MD5哈希算法;
func md5_02() {
	data := []byte("Hello World.")
	h := md5.New()
	h.Write(data)
	fmt.Println("MD5:", h.Sum(nil))                     // 返回结果为整数类型切片;
	fmt.Println(string(h.Sum(nil)))                     // 乱码;
	fmt.Println("MD5:", hex.EncodeToString(h.Sum(nil))) // (一般性)转为十六进制编码;
}
