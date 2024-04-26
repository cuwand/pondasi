package aes

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"strings"
	"unicode"
)

func PKCS5PaddingX(ciphertext []byte, blockSize int, after int) []byte {
	padding := (blockSize - len(ciphertext)%blockSize)
	padtext := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(ciphertext, padtext...)
}

func EncryptHex(plaintext string, key string, iv string, blockSize int) string {
	bKey := []byte(key)
	bIV := []byte(iv)
	bPlaintext := PKCS5PaddingX([]byte(plaintext), blockSize, len(plaintext))
	block, err := aes.NewCipher(bKey)
	if err != nil {
		panic(err)
	}
	ciphertext := make([]byte, len(bPlaintext))
	mode := cipher.NewCBCEncrypter(block, bIV)
	mode.CryptBlocks(ciphertext, bPlaintext)
	return hex.EncodeToString(ciphertext)
}

func EncryptBase64(plaintext string, key string, iv string, blockSize int) string {
	bKey := []byte(key)
	bIV := []byte(iv)
	bPlaintext := PKCS5PaddingX([]byte(plaintext), blockSize, len(plaintext))
	block, err := aes.NewCipher(bKey)
	if err != nil {
		panic(err)
	}
	ciphertext := make([]byte, len(bPlaintext))
	mode := cipher.NewCBCEncrypter(block, bIV)
	mode.CryptBlocks(ciphertext, bPlaintext)

	return base64.StdEncoding.EncodeToString(ciphertext)
}

func DecryptHex(cipherText string, encKey string, iv string) (decryptedString string) {
	bKey := []byte(encKey)
	bIV := []byte(iv)
	cipherTextDecoded, err := hex.DecodeString(cipherText)
	if err != nil {
		panic(err)
	}

	block, err := aes.NewCipher(bKey)
	if err != nil {
		panic(err)
	}

	mode := cipher.NewCBCDecrypter(block, bIV)
	mode.CryptBlocks([]byte(cipherTextDecoded), []byte(cipherTextDecoded))

	cipherTextDecodedStr := string(cipherTextDecoded)

	// Clean String from non printable
	cipherTextDecodedStr = strings.Map(func(r rune) rune {
		if unicode.IsPrint(r) {
			return r
		}
		return -1
	}, cipherTextDecodedStr)

	return cipherTextDecodedStr
}

func DecryptBase64(cipherText string, encKey string, iv string) (decryptedString string) {
	var decodedByte, _ = base64.StdEncoding.DecodeString(cipherText)

	bKey := []byte(encKey)
	bIV := []byte(iv)

	block, err := aes.NewCipher(bKey)
	if err != nil {
		panic(err)
	}

	mode := cipher.NewCBCDecrypter(block, bIV)
	mode.CryptBlocks([]byte(decodedByte), []byte(decodedByte))

	decodedByteStr := string(decodedByte)

	// Clean String from non printable
	decodedByteStr = strings.Map(func(r rune) rune {
		//if unicode.IsSymbol(r) {
		//	return -1
		//}

		if unicode.IsPrint(r) {
			return r
		}

		return -1
	}, decodedByteStr)

	return decodedByteStr
}

func EncryptPayload(key string, iv string, payload interface{}) string {
	marshaledPayload, err := json.Marshal(payload)

	if err != nil {
		panic(err)
	}

	fmt.Println(string(marshaledPayload))

	return EncryptHex(string(marshaledPayload[:]), key, iv, aes.BlockSize)
}

func DecryptPayload(key string, iv string, cipherText string, payload interface{}) {
	plainText := DecryptHex(cipherText, key, iv)

	err := json.Unmarshal([]byte(plainText), &payload)

	if err != nil {
		panic(err)
	}
}
