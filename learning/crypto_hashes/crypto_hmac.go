// GO语言中的HMAC加密哈希信息认证码算法(crypto/hmac包);
// HMAC是使用key标记信息的加密hash, 接收者使用相同的key逆运算来认证hash;
package crypto_hashes

import "fmt"
import "crypto/sha256"
import "crypto/hmac"
import "encoding/hex"
import "github.com/kingreatwill/go_study/learning/common"

// HMAC加密哈希信息认证码算法;
func hmac_() {
	common.Println("crypto_hmac.go")

	key := []byte("4098bs908a0dgf8gfbx098dfa08gagdfa78gad89g798ad8g")
	data := []byte("Hello World.")
	h := hmac.New(sha256.New, key)
	// h := hmac.New(sha256.New224, key)
	h.Write(data)
	fmt.Println("hmac:", h.Sum(nil))                     // 返回结果为整数类型切片;
	fmt.Println(string(h.Sum(nil)))                      // 乱码;
	fmt.Println("hmac:", hex.EncodeToString(h.Sum(nil))) // (一般性)转为十六进制编码;

	// 比较两个MAC是否相同, 而不会泄露对比时间信息;
	// func Equal(mac1, mac2 []byte) bool;
	// hmac.Equal(messageMAC, expectedMAC)
}
