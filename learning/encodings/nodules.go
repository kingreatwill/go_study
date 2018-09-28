// GO语言的encodings小结;
package encodings

// Test 测试;
func Test() {
	// unicode01;
	unicode01()
	// utf801;
	utf801()
	// utf1601;
	utf1601()

	// ascii85_;
	ascii85_()
	// base32_;
	base32_()
	// base64_;
	base64_()
	// binary01;
	binary01()

	// csv01;
	csv01()
	// csv02;
	csv02()
	// csv03;
	csv03()
	// csv04;
	csv04()
	// csv00;
	csv00()
	// gob01;
	gob01()
	// hex01;
	hex01()

	// json01;
	struct2json()
	// json2struct;
	json2struct()

	// pem_Generate_Rsa_Key;
	pem_Generate_Rsa_Key()
	// pem_Rsa_Crypt;
	pem_Rsa_Crypt()

	// xml01;
	struct2xml()
	// xml02;
	xml2struct()
}
