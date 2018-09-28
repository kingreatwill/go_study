// GO语言中的SHA384和SHA512哈希算法(crypto/sha512包);
// 校验和算法;
package crypto_hashes

import "fmt"
import "crypto/sha512"
import "encoding/hex"
import "github.com/kingreatwill/go_study/learning/common"

// SHA384哈希算法;
func sha384() {
	common.Println("crypto_sha512.go")

	data := []byte("Hello World.")
	h := sha512.New384()
	h.Write(data)
	fmt.Println("SHA384:", h.Sum(nil))                     // 返回结果为整数类型切片;
	fmt.Println(string(h.Sum(nil)))                        // 乱码;
	fmt.Println("SHA384:", hex.EncodeToString(h.Sum(nil))) // (一般性)转为十六进制编码;
}

// SHA512哈希算法;
func sha512_() {
	data := []byte("Hello World.")
	h := sha512.New()
	h.Write(data)
	fmt.Println("SHA512:", h.Sum(nil))                     // 返回结果为整数类型切片;
	fmt.Println(string(h.Sum(nil)))                        // 乱码;
	fmt.Println("SHA512:", hex.EncodeToString(h.Sum(nil))) // (一般性)转为十六进制编码;
}
