package caesar_cipher

import (
	"reflect"
	"testing"
)

func TestEncrypt(t *testing.T) {
	type args struct {
		plaintext string
		offset    []int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		// TODO: Add test cases.
		{"", args{
			plaintext: "cc11001100",
			offset:    []int{20},
		}, "ww11001100"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Encrypt(tt.args.plaintext, tt.args.offset...); got != tt.want {
				t.Errorf("Encrypt() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDecrypt(t *testing.T) {
	type args struct {
		encryptText string
		offset      []int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		// TODO: Add test cases.
		{"", args{
			encryptText: "ww11001100",
			offset:      []int{20},
		}, "cc11001100"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Decrypt(tt.args.encryptText, tt.args.offset...); got != tt.want {
				t.Errorf("Decrypt() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestEncryptBytes(t *testing.T) {
	type args struct {
		originalBytes []byte
		offset        []int
	}
	tests := []struct {
		name string
		args args
		want []byte
	}{
		// TODO: Add test cases.
		{
			"",
			args{
				originalBytes: []byte{0, 1, 2, 3, 4, 5, 6, 7},
				offset:        []int{1},
			},
			[]byte{1, 2, 3, 4, 5, 6, 7, 8},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := EncryptBytes(tt.args.originalBytes, tt.args.offset...); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("EncryptBytes() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDecryptBytes(t *testing.T) {
	type args struct {
		encryptBytes []byte
		offset       []int
	}
	tests := []struct {
		name string
		args args
		want []byte
	}{
		// TODO: Add test cases.
		{
			"",
			args{
				encryptBytes: []byte{1, 2, 3, 4, 5, 6, 7, 8},
				offset:       []int{1},
			},
			[]byte{0, 1, 2, 3, 4, 5, 6, 7},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := DecryptBytes(tt.args.encryptBytes, tt.args.offset...); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("DecryptBytes() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestEncryptUnicode(t *testing.T) {
	type args struct {
		text   string
		offset []int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		// TODO: Add test cases.
		{
			args: args{
				text:   "中文凯撒移位",
				offset: []int{1},
			},
			want: "丮斈凰撓秼低",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := EncryptUnicode(tt.args.text, tt.args.offset...); got != tt.want {
				t.Errorf("EncryptUnicode() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDecryptUnicode(t *testing.T) {
	type args struct {
		encryptText string
		offset      []int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		// TODO: Add test cases.
		{
			args: args{
				encryptText: "丮斈凰撓秼低",
				offset:      []int{1},
			},
			want: "中文凯撒移位",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := DecryptUnicode(tt.args.encryptText, tt.args.offset...); got != tt.want {
				t.Errorf("DecryptUnicode() = %v, want %v", got, tt.want)
			}
		})
	}
}
