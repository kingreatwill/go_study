// GO语言中的AES加密算法(crypto/aes包);
// crypto/cipher包实现了多个标准的用于包装底层块加密算法的加密算法实现;
package crypto_others

import "fmt"
import "crypto/aes"
import "crypto/cipher"
import "encoding/hex"
import "github.com/kingreatwill/go_study/learning/common"

// GCM加密(AES-128/AES-192/AES-256);
func aes_gcm_encrypt() {
	common.Println("crypto_aes.gcm.go")

	plaintext := []byte("NIST Special Publication 800-38A.")
	// key := []byte("1cd5b4cf89949200")              // 16字节对应AES-128算法;
	// key := []byte("1cd5b4cf8994920012345678")      // 24字节对应AES-192算法;
	key := []byte("1cd5b4cf899492001234567812345678") // 32字节对应AES-256算法;
	nonce := []byte("1cd5b4cf8994")                   // 12字节;
	block, err1 := aes.NewCipher(key)
	if err1 != nil {
		panic(err1)
	}
	aead, err2 := cipher.NewGCM(block)
	if err2 != nil {
		panic(err2)
	}
	ciphertext := aead.Seal(nil, nonce, plaintext, nil)
	fmt.Println("AES-256:", ciphertext)                     // 返回结果为整数类型切片;
	fmt.Println(string(ciphertext))                         // 乱码;
	fmt.Println("AES-256:", hex.EncodeToString(ciphertext)) // (一般性)转为十六进制编码;
	// Output: 52794e4a9e4c705c1ace16909824e659c003fa68c962bf7b1242a63648ee96321f508802e7e9c6e2ba48f8ae03cbb376d2;
}

// GCM解密(AES-128/AES-192/AES-256);
func aes_gcm_decrypt() {
	ciphertext, _ := hex.DecodeString("52794e4a9e4c705c1ace16909824e659c003fa68c962bf7b1242a63648ee96321f508802e7e9c6e2ba48f8ae03cbb376d2")
	// key := []byte("1cd5b4cf89949200")              // 16字节对应AES-128算法;
	// key := []byte("1cd5b4cf8994920012345678")      // 24字节对应AES-192算法;
	key := []byte("1cd5b4cf899492001234567812345678") // 32字节对应AES-256算法;
	nonce := []byte("1cd5b4cf8994")                   // 12字节;
	block, err1 := aes.NewCipher(key)
	if err1 != nil {
		panic(err1)
	}
	aead, err2 := cipher.NewGCM(block)
	if err2 != nil {
		panic(err2)
	}
	plaintext, err3 := aead.Open(nil, nonce, ciphertext, nil)
	if err3 != nil {
		panic(err3)
	}
	fmt.Println("AES-256:", string(plaintext))
}
