package caesar_cipher

import (
	variable_parameter "github.com/golang-infrastructure/go-variable-parameter"
	"math"
)

// ------------------------------------------------- --------------------------------------------------------------------

// DefaultOffset 最原始的凯撒加密就是向右偏移3位
const DefaultOffset = 3

// Encrypt 加密
func Encrypt(plaintext string, offset ...int) string {
	offset = variable_parameter.SetDefaultParam(offset, DefaultOffset)

	result := make([]rune, len(plaintext))
	for index, char := range plaintext {
		if char >= 'a' && char <= 'z' {
			// 支持负数偏移量
			target := (int(char-'a') + 26 + offset[0]) % 26
			result[index] = 'a' + rune(target)
		} else if char >= 'A' && char <= 'Z' {
			// 支持负数偏移量
			target := (int(char-'A') + 26 + offset[0]) % 26
			result[index] = 'A' + rune(target)
		} else {
			result[index] = char
		}
	}
	return string(result)
}

// Decrypt 解密
func Decrypt(encryptText string, offset ...int) string {
	// 直接把偏移量取反复用加密逻辑即可
	t := variable_parameter.TakeFirstParamOrDefault(offset, DefaultOffset) * -1
	return Encrypt(encryptText, t)
}

// ------------------------------------------------- --------------------------------------------------------------------

// EncryptBytes 凯撒加密的扩展，对字节进行凯撒加密
func EncryptBytes(originalBytes []byte, offset ...int) []byte {
	offset = variable_parameter.SetDefaultParam(offset, DefaultOffset)

	result := make([]byte, len(originalBytes))
	for index, b := range originalBytes {
		target := (int(b) + 128 + offset[0]) % 128
		result[index] = uint8(target)
	}
	return result
}

// DecryptBytes 对字节进行凯撒解密
func DecryptBytes(encryptBytes []byte, offset ...int) []byte {
	// 直接把偏移量取反复用加密逻辑即可
	t := variable_parameter.TakeFirstParamOrDefault(offset, DefaultOffset) * -1
	return EncryptBytes(encryptBytes, t)
}

// ------------------------------------------------- --------------------------------------------------------------------

// EncryptUnicode 扩展的凯撒加密，在unicode层面加密
func EncryptUnicode(text string, offset ...int) string {
	offset = variable_parameter.SetDefaultParam(offset, DefaultOffset)

	result := make([]rune, 0)
	for _, char := range text {
		target := (int64(char) + int64(offset[0]) + math.MaxInt32) % math.MaxInt32
		result = append(result, rune(target))
	}
	return string(result)
}

// DecryptUnicode 对凯撒unicode方式加密过的密文解密
func DecryptUnicode(encryptText string, offset ...int) string {
	// 直接把偏移量取反复用加密逻辑即可
	t := variable_parameter.TakeFirstParamOrDefault(offset, DefaultOffset) * -1
	return EncryptUnicode(encryptText, t)
}

// ------------------------------------------------- --------------------------------------------------------------------
