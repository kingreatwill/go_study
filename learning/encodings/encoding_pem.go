// GO语言中的简单的数字与字节序列的转换以及变长值的编解码(encoding/pem包);
// PEM数据编码(源自保密增强邮件协议);
// 目前PEM编码主要用于TLS密钥和证书;

// Block代表PEM编码的结构:
// type Block struct {
//    Type    string            // 得自前言的类型（如"RSA PRIVATE KEY"）
//    Headers map[string]string // 可选的头项
//    Bytes   []byte            // 内容解码后的数据，一般是DER编码的ASN.1结构
// }

// PEM编码格式:
// -----BEGIN Type-----
// Headers
// base64-encoded Bytes
// -----END Type-----

package encodings

import "fmt"
import "os"
import "io/ioutil"
import "strings"
import "crypto/rsa"
import "crypto/rand"
import "crypto/x509"
import "encoding/base64"
import "encoding/pem"
import "errors"
import "github.com/kingreatwill/go_study/learning/common"

// pem_Generate_Rsa_Key;
// func Encode(out io.Writer, b *Block) error;
func pem_Generate_Rsa_Key() {
	common.Println("encoding_pem.go")

	// 生成私钥;
	privateKey, err11 := rsa.GenerateKey(rand.Reader, 1024)
	if err11 != nil {
		fmt.Println(err11)
		return
	}
	block1 := &pem.Block{Type: "RSA PRIVATE KEY", Bytes: x509.MarshalPKCS1PrivateKey(privateKey)}
	root1, _ := os.Getwd()
	_path1 := strings.Join([]string{root1, "doc", "temp", "private.pem"}, string(os.PathSeparator))
	file1, err12 := os.Create(_path1)
	if err12 != nil {
		fmt.Println(err12)
		return
	}
	defer file1.Close()
	err13 := pem.Encode(file1, block1)
	if err13 != nil {
		fmt.Println(err13)
		return
	}
	// -----BEGIN RSA PRIVATE KEY-----
	// MIICXQIBAAKBgQC3v7adtZCjwDsjDfjyQbYYYMOdRMlFqmKiPz8yDqFioioLLtWV
	// n2jkckKxNzc5SeE0h7DSMeDly+zU2PukZRbTlyYEtvstwc8iM14UBRFye4z1pHkH
	// Kn5g+d9QhM7jljh+DLuBDNbJd9wR27D+wHDuBydnlAZ3To9sMqTJWuzw8QIDAQAB
	// AoGBAJ0UZe+yXGCwH53a/vtSU3HJntAeEpXcj4811DddrHcePCTtN6c97DSGvR60
	// HiB6WOxIJ0+5VhH8X4yhpZWm2Xg6KdqiADD4bf5ldulNnFmovvjpO9fVZF+JqZi0
	// kIK59GOaruglIKqmbsQqzaNtn7JjAvxjKf1i615jG6BC7xNBAkEA7hcC42ymIPcv
	// yPAb7THJpohp3tq1MlqyRLKezVitaZSSSMu7/AFjXsGxcm/aWR8DIl++wvHufKjd
	// FS6Nm6AFyQJBAMWSPICo4tudfKpCCRbIfdDAav/VZQzL8P2430p+m9wDTxgrwxd0
	// m0ID864Q1pAXkivStNaNeE57RwyQbr00xekCQQCvb89WC0IiumHAm4TIc85V4uC4
	// MMohj4kaa8/uGfC2Ap8D+vPKzq8NLIZ3f4oR73McyhZDiYek3b0ClaoMyY5ZAkAN
	// m67zJa7KdaR6jaXWneSclkRIbshRA1MMBsnKdcICe2/dOFZtrShseZ01Jg/BSP5W
	// amLuTDw3G6KfKOuXuKTRAkAq9wlbbsnWUxBOsoNV31tE7T3jh3vbrdNMNHkxjAc3
	// sbf83HlR38Eh6KoUne5jA2l+gVL+vaNPRtcaB6jvLrMu
	// -----END RSA PRIVATE KEY-----

	// 生成公钥;
	publicKey, err21 := x509.MarshalPKIXPublicKey(&privateKey.PublicKey)
	if err21 != nil {
		fmt.Println(err21)
		return
	}
	block2 := &pem.Block{Type: "PUBLIC KEY", Bytes: publicKey}
	root2, _ := os.Getwd()
	_path2 := strings.Join([]string{root2, "doc", "temp", "public.pem"}, string(os.PathSeparator))
	file2, err22 := os.Create(_path2)
	if err22 != nil {
		fmt.Println(err22)
		return
	}
	defer file2.Close()
	err23 := pem.Encode(file2, block2)
	if err23 != nil {
		fmt.Println(err23)
		return
	}
	// -----BEGIN PUBLIC KEY-----
	// MIGfMA0GCSqGSIb3DQEBAQUAA4GNADCBiQKBgQC3v7adtZCjwDsjDfjyQbYYYMOd
	// RMlFqmKiPz8yDqFioioLLtWVn2jkckKxNzc5SeE0h7DSMeDly+zU2PukZRbTlyYE
	// tvstwc8iM14UBRFye4z1pHkHKn5g+d9QhM7jljh+DLuBDNbJd9wR27D+wHDuBydn
	// lAZ3To9sMqTJWuzw8QIDAQAB
	// -----END PUBLIC KEY-----
}

// pem_Rsa_Crypt;
// func Decode(data []byte) (p *Block, rest []byte);
func pem_Rsa_Crypt() {
	plaintext := []byte("Advanced Encoding and Decoding Techniques in Go.")

	// 加载公钥PEM文件;
	root1, _ := os.Getwd()
	_path1 := strings.Join([]string{root1, "doc", "temp", "public.pem"}, string(os.PathSeparator))
	file1, err11 := os.Open(_path1)
	if err11 != nil {
		fmt.Println(err11)
		return
	}
	defer file1.Close()
	pemBytes1, _ := ioutil.ReadAll(file1)

	// RSA加密(演示);
	block1, _ := pem.Decode(pemBytes1)
	if block1 == nil {
		fmt.Println(errors.New("public key error"))
		return
	}
	pubInterface, err := x509.ParsePKIXPublicKey(block1.Bytes)
	if err != nil {
		fmt.Println(err)
		return
	}
	publicKey := pubInterface.(*rsa.PublicKey)
	ciphertext, _ := rsa.EncryptPKCS1v15(rand.Reader, publicKey, plaintext)
	ciphertextbase64 := base64.StdEncoding.EncodeToString(ciphertext)
	fmt.Println("rsa encrypt base64:" + ciphertextbase64)

	// 加载私钥PEM文件;
	root2, _ := os.Getwd()
	_path2 := strings.Join([]string{root2, "doc", "temp", "private.pem"}, string(os.PathSeparator))
	file2, err22 := os.Open(_path2)
	if err22 != nil {
		fmt.Println(err22)
		return
	}
	defer file2.Close()
	pemBytes2, _ := ioutil.ReadAll(file2)
	// RSA解密(演示);
	block2, _ := pem.Decode(pemBytes2)
	if block2 == nil {
		fmt.Println(errors.New("private key error"))
		return
	}
	privateKey, err23 := x509.ParsePKCS1PrivateKey(block2.Bytes)
	if err23 != nil {
		fmt.Println(err23)
		return
	}
	s, _ := base64.StdEncoding.DecodeString(ciphertextbase64)
	plaintext2, _ := rsa.DecryptPKCS1v15(rand.Reader, privateKey, s)
	fmt.Println("rsa decrypt:" + string(plaintext2))
}
