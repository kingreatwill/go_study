// GO语言中的DES加密算法(crypto/des包);
// crypto/cipher包实现了多个标准的用于包装底层块加密算法的加密算法实现;
package crypto_others

import "fmt"
import "crypto/des"
import "crypto/cipher"
import "encoding/hex"
import "github.com/kingreatwill/go_study/learning/common"

// CTR加密(DES);
func des_ctr_encrypt() {
	common.Println("crypto_des.ctr.go")

	plaintext := []byte("NIST Special Publication 800-38A.")
	ciphertext := make([]byte, len(plaintext))
	// key := []byte("1cd5b4c8")              // 8字节;
	key := []byte("1cd5b4c81cd5b4c81cd5b4c8") // 24字节;
	iv := []byte("e6db2718")                  // 8字节;
	// block, err := des.NewCipher(key)       // DES算法;
	block, err := des.NewTripleDESCipher(key) // TDEA算法;
	if err != nil {
		panic(err)
	}
	stream := cipher.NewCTR(block, iv)
	stream.XORKeyStream(ciphertext, plaintext)
	fmt.Println("DES:", ciphertext)                     // 返回结果为整数类型切片;
	fmt.Println(string(ciphertext))                     // 乱码;
	fmt.Println("DES:", hex.EncodeToString(ciphertext)) // (一般性)转为十六进制编码;
	// Output: e90eecff2cc0134107338f486af54458cc07b043e564d47d1d95f8f9064ba5e5db;
}

// CTR解密(DES);
func des_ctr_decrypt() {
	ciphertext, _ := hex.DecodeString("e90eecff2cc0134107338f486af54458cc07b043e564d47d1d95f8f9064ba5e5db") // 16字节;
	plaintext := make([]byte, len(ciphertext))                                                              // 16字节;
	// key := []byte("1cd5b4c8")              // 8字节;
	key := []byte("1cd5b4c81cd5b4c81cd5b4c8") // 24字节;
	iv := []byte("e6db2718")                  // 8字节;
	// block, err := des.NewCipher(key)       // DES算法;
	block, err := des.NewTripleDESCipher(key) // TDEA算法;
	if err != nil {
		panic(err)
	}
	stream := cipher.NewCTR(block, iv)
	stream.XORKeyStream(plaintext, ciphertext)
	fmt.Println("DES:", string(plaintext))
}
