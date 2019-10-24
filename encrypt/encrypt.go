package encrypt

import (
	"encoding/base64"
	"golang.org/x/crypto/scrypt"
	"log"
)

func Encrypt(pwd string) string {
	salt := []byte{0xcf, 0xf8, 0xf2, 0x58, 0xa2, 0x6b, 0xad, 0x7b}
	dk, err := scrypt.Key([]byte(pwd), salt, 1<<15, 8, 1, 32)
	if err != nil {
		log.Println(err)
	}
	return base64.StdEncoding.EncodeToString(dk)
}
