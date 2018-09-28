// GO语言中的SHA224和SHA256哈希算法(crypto/sha256包);
// 校验和算法;
package crypto_hashes

import "fmt"
import "crypto/sha256"
import "encoding/hex"
import "github.com/kingreatwill/go_study/learning/common"

// SHA224哈希算法;
func sha224() {
	common.Println("crypto_sha256.go")

	data := []byte("Hello World.")
	h := sha256.New224()
	h.Write(data)
	fmt.Println("SHA224:", h.Sum(nil))                     // 返回结果为整数类型切片;
	fmt.Println(string(h.Sum(nil)))                        // 乱码;
	fmt.Println("SHA224:", hex.EncodeToString(h.Sum(nil))) // (一般性)转为十六进制编码;
}

// SHA256哈希算法;
func sha256_() {
	data := []byte("Hello World.")
	h := sha256.New()
	h.Write(data)
	fmt.Println("SHA256:", h.Sum(nil))                     // 返回结果为整数类型切片;
	fmt.Println(string(h.Sum(nil)))                        // 乱码;
	fmt.Println("SHA256:", hex.EncodeToString(h.Sum(nil))) // (一般性)转为十六进制编码;
}
