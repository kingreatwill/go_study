// GO语言中的AES加密算法(crypto/aes包);
// crypto/cipher包实现了多个标准的用于包装底层块加密算法的加密算法实现;
package crypto_others

import "fmt"
import "crypto/aes"
import "crypto/cipher"
import "encoding/hex"
import "github.com/kingreatwill/go_study/learning/common"

// CBC加密(AES-128/AES-192/AES-256);
func aes_cbc_encrypt() {
	common.Println("crypto_aes.cbc.go")

	plaintext := []byte("exampleplaintext")    // 16*N字节;
	ciphertext := make([]byte, len(plaintext)) // 16*N字节;
	// key := []byte("1cd5b4cf89949200")              // 16字节对应AES-128算法;
	// key := []byte("1cd5b4cf8994920012345678")      // 24字节对应AES-192算法;
	key := []byte("1cd5b4cf899492001234567812345678") // 32字节对应AES-256算法;
	iv := []byte("e6db271db12d4d80")                  // 16字节;
	block, err := aes.NewCipher(key)
	if err != nil {
		panic(err)
	}
	mode := cipher.NewCBCEncrypter(block, iv)
	mode.CryptBlocks(ciphertext, plaintext)
	fmt.Println("AES-256:", ciphertext)                     // 返回结果为整数类型切片;
	fmt.Println(string(ciphertext))                         // 乱码;
	fmt.Println("AES-256:", hex.EncodeToString(ciphertext)) // (一般性)转为十六进制编码;
	// Output: d39691541280030d032f96bc7b8ae61b;
}

// CBC解密(AES-128/AES-192/AES-256);
func aes_cbc_decrypt() {
	ciphertext, _ := hex.DecodeString("d39691541280030d032f96bc7b8ae61b") // 16*N字节;
	plaintext := make([]byte, len(ciphertext))                            // 16*N字节;
	// key := []byte("1cd5b4cf89949200")              // 16字节对应AES-128算法;
	// key := []byte("1cd5b4cf8994920012345678")      // 24字节对应AES-192算法;
	key := []byte("1cd5b4cf899492001234567812345678") // 32字节对应AES-256算法;
	iv := []byte("e6db271db12d4d80")                  // 16字节;
	block, err := aes.NewCipher(key)
	if err != nil {
		panic(err)
	}
	mode := cipher.NewCBCDecrypter(block, iv)
	mode.CryptBlocks(plaintext, ciphertext)
	fmt.Println("AES-256:", string(plaintext))
}
