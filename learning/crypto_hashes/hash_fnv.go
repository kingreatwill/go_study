// GO语言中的hash校验和(hash/fnv包);
// 校验和算法;
package crypto_hashes

import "fmt"
import "hash/fnv"
import "encoding/hex"
import "github.com/kingreatwill/go_study/learning/common"

// FNV-1和FNV-1a(非加密hash函数);
func hashfnv() {
	common.Println("hash_fnv.go")

	//	type Hash interface {
	//		// 通过嵌入的匿名io.Writer接口的Write方法向hash中添加更多数据，永远不返回错误
	//		io.Writer
	//		// 返回添加b到当前的hash值后的新切片，不会改变底层的hash状态
	//		Sum(b []byte) []byte
	//		// 重设hash为无数据输入的状态
	//		Reset()
	//		// 返回Sum会返回的切片的长度
	//		Size() int
	//		// 返回hash底层的块大小；Write方法可以接受任何大小的数据，
	//		// 但提供的数据是块大小的倍数时效率更高
	//		BlockSize() int
	//	}

	//	type Hash32 interface {
	//		Hash
	//		Sum32() uint32
	//	}

	//	type Hash64 interface {
	//		Hash
	//		Sum64() uint64
	//	}

	// func New32() hash.Hash32;
	data := []byte("058kl5j204809d8lk452098")
	hs1 := fnv.New32()
	hs1.Write(data)
	fmt.Println("fnv:", hs1.Sum(nil))                     // 返回结果为整数类型切片;
	fmt.Println("fnv:", string(hs1.Sum(nil)))             // 乱码;
	fmt.Println("fnv:", hex.EncodeToString(hs1.Sum(nil))) // (一般性)转为十六进制编码;
}
