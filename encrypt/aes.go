package encrypt

import (
	"crypto/aes"
	"crypto/cipher"
)

func AESEncrypt(plantText, key []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}
	plantText = PKCS7Padding(plantText, block.BlockSize())
	ciphertext := make([]byte, len(plantText))
	blockMode := cipher.NewCBCEncrypter(block, key)
	blockMode.CryptBlocks(ciphertext, plantText)
	return ciphertext, nil
}

func AESDecrypt(ciphertext, key []byte) ([]byte, error) {
	keyBytes := []byte(key)
	block, err := aes.NewCipher(keyBytes)
	if err != nil {
		return nil, err
	}
	plaintext := make([]byte, len(ciphertext))
	blockMode := cipher.NewCBCDecrypter(block, keyBytes)
	blockMode.CryptBlocks(plaintext, ciphertext)
	plaintext = PKCS7UnPadding(plaintext, block.BlockSize())
	return plaintext, nil
}
