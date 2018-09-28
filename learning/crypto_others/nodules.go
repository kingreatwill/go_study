// GO语言的crypto_others小结;
package crypto_others

// 测试;
func Test() {
	// CBC加密(AES-128/AES-192/AES-256);
	aes_cbc_encrypt()
	// CBC解密(AES-128/AES-192/AES-256);
	aes_cbc_decrypt()

	// CFB加密(AES-128/AES-192/AES-256);
	aes_cfb_encrypt()
	// CFB解密(AES-128/AES-192/AES-256);
	aes_cfb_decrypt()

	// CTR加密(AES-128/AES-192/AES-256);
	aes_ctr_encrypt()
	// CTR解密(AES-128/AES-192/AES-256);
	aes_ctr_decrypt()

	// GCM加密(AES-128/AES-192/AES-256);
	aes_gcm_encrypt()
	// GCM解密(AES-128/AES-192/AES-256);
	aes_gcm_decrypt()

	// OFB加密(AES-128/AES-192/AES-256);
	aes_ofb_encrypt()
	// OFB解密(AES-128/AES-192/AES-256);
	aes_ofb_decrypt()

	// CBC加密(DES);
	des_cbc_encrypt()
	// CBC解密(DES);
	des_cbc_decrypt()

	// CFB加密(DES);
	des_cfb_encrypt()
	// CFB加密(DES);
	des_cfb_decrypt()

	// CTR加密(DES);
	des_ctr_encrypt()
	// CTR解密(DES);
	des_ctr_decrypt()

	// GCM加密(DES);
	// des_gcm_encrypt()
	// GCM解密(DES);
	// des_gcm_decrypt()

	// OFB加密(DES);
	des_ofb_encrypt()
	// OFB解密(DES);
	des_ofb_decrypt()
}
