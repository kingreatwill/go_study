// GO语言中的hash校验和(hash/adler32包);
// 校验和算法;
package crypto_hashes

import "fmt"
import "hash/adler32"
import "github.com/kingreatwill/go_study/learning/common"

// Adler-32校验和算法;
func hashadler32() {
	common.Println("hash_adler32.go")

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

	// 返回数据data的Adler-32校验和;
	// import "hash/adler32";
	// func Checksum(data []byte) uint32;
	data := []byte("058kl5j204809d8lk452098")
	fmt.Println("adler32:", adler32.Checksum(data))
}
