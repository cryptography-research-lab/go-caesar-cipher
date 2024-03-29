package caesar_cipher

import (
	"github.com/stretchr/testify/assert"
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

func TestCrackOffset(t *testing.T) {
	ciphertext := "khoorzruog,lwfdqqrwehghfubswlrq"
	offset, s := CrackOffset(ciphertext)
	assert.Equal(t, 3, offset)
	assert.Equal(t, "helloworld,itcannotbedecryption", s)
}
