// GO语言中的AES加密算法(crypto/aes包);
// crypto/cipher包实现了多个标准的用于包装底层块加密算法的加密算法实现;
package crypto_others

import "fmt"
import "crypto/aes"
import "crypto/cipher"
import "encoding/hex"
import "github.com/kingreatwill/go_study/learning/common"

// OFB加密(AES-128/AES-192/AES-256);
func aes_ofb_encrypt() {
	common.Println("crypto_aes.ofb.go")

	plaintext := []byte("NIST Special Publication 800-38A.")
	ciphertext := make([]byte, len(plaintext))
	// key := []byte("1cd5b4cf89949200")              // 16字节对应AES-128算法;
	// key := []byte("1cd5b4cf8994920012345678")      // 24字节对应AES-192算法;
	key := []byte("1cd5b4cf899492001234567812345678") // 32字节对应AES-256算法;
	iv := []byte("e6db271db12d4d80")                  // 16字节;
	block, err := aes.NewCipher(key)
	if err != nil {
		panic(err)
	}
	stream := cipher.NewOFB(block, iv)
	stream.XORKeyStream(ciphertext, plaintext)
	fmt.Println("AES-256:", ciphertext)                     // 返回结果为整数类型切片;
	fmt.Println(string(ciphertext))                         // 乱码;
	fmt.Println("AES-256:", hex.EncodeToString(ciphertext)) // (一般性)转为十六进制编码;
	// Output: 93e113493e81ff3ca9bbedfcadf42755ae9c3e10175ad890ef622b891cb9f6b281;
}

// OFB解密(AES-128/AES-192/AES-256);
func aes_ofb_decrypt() {
	ciphertext, _ := hex.DecodeString("93e113493e81ff3ca9bbedfcadf4275537e353e5874289bbaff13b33d72ad7704c")
	plaintext := make([]byte, len(ciphertext))
	// key := []byte("1cd5b4cf89949200")              // 16字节对应AES-128算法;
	// key := []byte("1cd5b4cf8994920012345678")      // 24字节对应AES-192算法;
	key := []byte("1cd5b4cf899492001234567812345678") // 32字节对应AES-256算法;
	iv := []byte("e6db271db12d4d80")                  // 16字节;
	block, err := aes.NewCipher(key)
	if err != nil {
		panic(err)
	}
	stream := cipher.NewOFB(block, iv)
	stream.XORKeyStream(plaintext, ciphertext)
	fmt.Println("AES-256:", string(plaintext))
}
