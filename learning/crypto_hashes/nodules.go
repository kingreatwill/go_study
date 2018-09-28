// GO语言的crypto_hashes小结;
package crypto_hashes

// 测试;
func Test() {
	// Adler-32校验和算法;
	hashadler32()
	// 32位循环冗余校验(CRC-32)的校验和算法;
	hashcrc32()
	// 64位循环冗余校验(CRC-64)的校验和算法;
	hashcrc64()
	// FNV-1和FNV-1a(非加密hash函数);
	hashfnv()
	// HMAC加密哈希信息认证码算法;
	hmac_()
	// MD5哈希算法;
	md5_01()
	// MD5哈希算法;
	md5_02()
	// SHA1哈希算法;
	sha1_01()
	// SHA1哈希算法;
	sha1_02()
	// SHA224哈希算法;
	sha224()
	// SHA256哈希算法;
	sha256_()
	// SHA384哈希算法;
	sha384()
	// SHA512哈希算法;
	sha512_()
}
