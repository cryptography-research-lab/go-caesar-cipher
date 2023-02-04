package caesar_cipher

import (
	"testing"
)

func TestCrack(t *testing.T) {

	// 正常加密
	plaintext := "helloworld,itcannotbedecryption"
	encrypt := Encrypt(plaintext)
	t.Log(encrypt)

	// 拿到密文，开始破解
	s := CrackForOne(encrypt)
	t.Log(s)

}
