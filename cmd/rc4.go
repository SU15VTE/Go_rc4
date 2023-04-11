package cmd

import "encoding/hex"

func String_to_byte(str string) []byte {
	return []byte(str)
}

// RC4加密函数
func RC4_encrypt(key []byte, plaintext []byte) string {
	s := make([]byte, 256)
	for i := range s {
		s[i] = byte(i)
	}
	j := 0
	for i := range s {
		j = (j + int(s[i]) + int(key[i%len(key)])) % 256
		s[i], s[j] = s[j], s[i]
	}
	ciphertext := make([]byte, len(plaintext))
	i, j := 0, 0
	for k := range plaintext {
		i = (i + 1) % 256
		j = (j + int(s[i])) % 256
		s[i], s[j] = s[j], s[i]
		ciphertext[k] = plaintext[k] ^ s[(int(s[i])+int(s[j]))%256]
	}

	return hex.EncodeToString(ciphertext)
}

// RC4解密函数
func RC4_decrypt(key []byte, ciphertext []byte) string {
	text, _ := hex.DecodeString(string(ciphertext))
	plaintext := RC4_encrypt(key, text)
	result, _ := hex.DecodeString(plaintext)
	return string(result)
}
