package aes

import (
	"crypto/aes"
	"fmt"
	"github.com/cuwand/pondasi/models"
	"strings"
	"testing"
)

const (
	plainText  = "halo ini adalah ichwan almaza halo ini adalah ichwan almazahalo ini adalah ichwan almaza"
	cipherText = "52FDFC072182654F163F5F0F9A621D723EB4B331185F371DF1EAFF475E48AC1F69E469330152B64ECD6790AECB93BC846276574B594036C7572ACA343E56E0C505D3F4C90A3B63B4D1E65F6479A291F2"
	key        = "12345678901234561234567890123456"
)

func TestEncryptPayload(t *testing.T) {
	cipherText := EncryptPayload("1234567891234567", "1234567891234567", models.User{
		Username: "ini username",
		FullName: "ini fullname",
		Identity: "ini identity",
	})

	fmt.Println(strings.ToUpper(cipherText))
}

func TestDecryptPayload(t *testing.T) {
	usr := models.User{}

	DecryptPayload(
		"1234567891234567", "0994567891234567",
		"6273484369345848a293454482f900068bf84b993b8bbd51bac0b3344d70fc88a78ed86fe522962944c09207a6e8f60fb0215b25468bc27d1b1585b1649ba008f3ee7a7e0fb0a8f3fb9e6e884a2d8721ebc16a848d512679d466b13042cec5ce",
		&usr)

	fmt.Println(usr.FullName)
	fmt.Println(usr.Username)
	fmt.Println(usr.Identity)
}

func TestEncrypt(t *testing.T) {
	//cipherText := Encrypt("0000000000000000", "ichwan almaza almaza ichwan")
	//cipherText := Ase256Encode("ichwan", "1234567891234567", "0994567891234567", aes.BlockSize)
	cipherText := EncryptBase64("ichwan",
		"1234567891234567",
		"0994567891234567",
		aes.BlockSize)

	fmt.Println(cipherText)

	//decrypted := Ase256Decode(cipherText, "1234567891234567", "0994567891234567")
	decrypted := DecryptBase64(cipherText, "1234567891234567", "0994567891234567")

	fmt.Println(decrypted)
}

func TestDecrypt(t *testing.T) {
	plainText := DecryptBase64(
		"gd9fwYhQvVaLcS+3cq7mCGzJvUgqob9l5eerXqrwnJDehrqFSqfG1EkrNBsLPh4LcNrfBnc3aQ4QoaMFomdS1VCldiO7PIRrW9pkXVwx5/zs2kaX6gLQ/9wOwDYEliCXwT4/wYxcMnDCH2WPO8TT8ztf93QQW8J8jGSjU7vg/pE=",
		"1234567891234567",
		"0235567891287593")

	fmt.Println(plainText)
}
