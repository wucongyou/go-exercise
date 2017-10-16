package encrypt

import (
	"crypto/cipher"
	"crypto/des"
)

func DESEncrypt(plaintext, key []byte) ([]byte, error) {
	block, err := des.NewCipher(key)
	if err != nil {
		return nil, err
	}
	plaintext = PKCS5Padding(plaintext, block.BlockSize())
	ciphertext := make([]byte, len(plaintext))
	blockMode := cipher.NewCBCEncrypter(block, key)
	blockMode.CryptBlocks(ciphertext, plaintext)
	return ciphertext, nil
}

func DESDecrypt(ciphertext, key []byte) ([]byte, error) {
	block, err := des.NewCipher(key)
	if err != nil {
		return nil, err
	}
	plaintext := make([]byte, len(ciphertext))
	blockMode := cipher.NewCBCDecrypter(block, key)
	blockMode.CryptBlocks(plaintext, ciphertext)
	plaintext = PKCS5UnPadding(plaintext, block.BlockSize())
	return plaintext, nil
}
