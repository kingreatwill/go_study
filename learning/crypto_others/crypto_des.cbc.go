// GO语言中的DES加密算法(crypto/des包);
// crypto/cipher包实现了多个标准的用于包装底层块加密算法的加密算法实现;
package crypto_others

import "fmt"
import "crypto/des"
import "crypto/cipher"
import "encoding/hex"
import "github.com/kingreatwill/go_study/learning/common"

// CBC加密(DES);
func des_cbc_encrypt() {
	common.Println("crypto_des.cbc.go")

	plaintext := []byte("exampleplaintext12345678") // 8*N字节;
	ciphertext := make([]byte, len(plaintext))      // 8*N字节;
	// key := []byte("1cd5b4c8")              // 8字节;
	key := []byte("1cd5b4c81cd5b4c81cd5b4c8") // 24字节;
	iv := []byte("e6db2718")                  // 8字节;
	// block, err := des.NewCipher(key)       // DES算法;
	block, err := des.NewTripleDESCipher(key) // TDEA算法;
	if err != nil {
		panic(err)
	}
	mode := cipher.NewCBCEncrypter(block, iv)
	mode.CryptBlocks(ciphertext, plaintext)
	fmt.Println("DES:", ciphertext)                     // 返回结果为整数类型切片;
	fmt.Println(string(ciphertext))                     // 乱码;
	fmt.Println("DES:", hex.EncodeToString(ciphertext)) // (一般性)转为十六进制编码;
	// Output: 9a4d16c4743540be362e21c4cab7d06a0ec0e8717f5579c2;
}

// CBC解密(DES);
func des_cbc_decrypt() {
	ciphertext, _ := hex.DecodeString("9a4d16c4743540be362e21c4cab7d06a0ec0e8717f5579c2") // 8*N字节;
	plaintext := make([]byte, len(ciphertext))                                            // 8*N字节;
	// key := []byte("1cd5b4c8")              // 8字节;
	key := []byte("1cd5b4c81cd5b4c81cd5b4c8") // 24字节;
	iv := []byte("e6db2718")                  // 8字节;
	// block, err := des.NewCipher(key)       // DES算法;
	block, err := des.NewTripleDESCipher(key) // TDEA算法;
	if err != nil {
		panic(err)
	}
	mode := cipher.NewCBCDecrypter(block, iv)
	mode.CryptBlocks(plaintext, ciphertext)
	fmt.Println("DES:", string(plaintext))
}
