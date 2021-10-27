package crypto

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
)

//PKCS7Padding use PKCS7 to padding
func PKCS7Padding(ciphertext []byte, blockSize int) []byte {
	padding := blockSize - len(ciphertext) % blockSize
	padtext := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(ciphertext, padtext...)
}
//PKCS7UnPadding use PkCS7 to unpadding
func PKCS7UnPadding(origData []byte) []byte {
	length := len(origData)
	unpadding := int(origData[length-1])
	return origData[:(length - unpadding)]
}

// AesCBCEncrypt use aes-cbc,pkcs7 padding to encrypt, support AES-128, AES-192, or AES-256.
func AesCBCEncrypt(origData []byte, key []byte, iv []byte) ([]byte, error) {
	newKey:=make([]byte,16)
	newIv:=make([]byte,16)

	copy(newKey,key)
	copy(newIv,iv)
	block, err := aes.NewCipher(newKey)
	if err != nil {
		return nil, err
	}

	blockSize := block.BlockSize()
	origData = PKCS7Padding(origData, blockSize)

	//使用cbc
	blocMode := cipher.NewCBCEncrypter(block, newIv)
	encrypted := make([]byte, len(origData))

	blocMode.CryptBlocks(encrypted, origData)
	return encrypted, nil
}

// AesCBCDecrypt use decrypt the content, like its name
func AesCBCDecrypt(cypted []byte, key []byte, iv []byte) ([]byte, error) {
	newKey:=make([]byte,16)
	newIv:=make([]byte,16)

	copy(newKey,key)
	copy(newIv,iv)

	block, err := aes.NewCipher(newKey)
	if err != nil {
		return nil, err
	}

	//使用key充当偏移量
	blockMode := cipher.NewCBCDecrypter(block, newIv)
	origData := make([]byte, len(cypted))

	blockMode.CryptBlocks(origData, cypted)

	origData = PKCS7UnPadding(origData)
	if err != nil {
		return nil, err
	}
	return origData, err
}
