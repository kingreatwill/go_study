// GO语言中的hash校验和(hash/crc32包);
// 校验和算法;
package crypto_hashes

import "fmt"
import "hash/crc32"
import "github.com/kingreatwill/go_study/learning/common"

// 32位循环冗余校验(CRC-32)的校验和算法;
func hashcrc32() {
	common.Println("hash_crc32.go")

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

	//	const (
	//		// 最常用的CRC-32多项式；用于以太网、v.42、fddi、gzip、zip、png、mpeg-2……
	//		IEEE = 0xedb88320
	//		// 卡斯塔尼奥利多项式，用在iSCSI；有比IEEE更好的错误探测特性
	//		// http://dx.doi.org/10.1109/26.231911
	//		Castagnoli = 0x82f63b78
	//		// 库普曼多项式；错误探测特性也比IEEE好
	//		// http://dx.doi.org/10.1109/DSN.2002.1028931
	//		Koopman = 0xeb31d82e
	//	)

	// 返回数据data使用IEEE多项式计算出的CRC-32校验和;
	// func ChecksumIEEE(data []byte) uint32;
	data := []byte("058kl5j204809d8lk452098")
	fmt.Println("crc32:", crc32.ChecksumIEEE(data))
	// 等价于fmt.Println("crc32:", crc32.Checksum(data, crc32.MakeTable(crc32.IEEE)))

	// 返回数据data使用tab代表的多项式计算出的CRC-32校验和;
	// func Checksum(data []byte, tab *Table) uint32;
	fmt.Println("crc32:", crc32.Checksum(data, crc32.MakeTable(crc32.Castagnoli)))
}
