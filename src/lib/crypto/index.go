package crypto

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
)

var secretKey = []byte("qiWyR.Uj2mc0Cs21KA21321Sd<A/D}S")

func EncodeToken(token []byte, err interface{}) []byte {
	if err != nil {
		panic(err)
	}

	encodedToken := token

	for key := range token {
		if key%2 == 0 {
			encodedToken[key] = encodedToken[key] + 2
		} else if key%3 == 0 {
			encodedToken[key] = encodedToken[key] - 4
		}
	}

	return encodedToken
}

func DecodeToken(token []byte) []byte {
	decodedToken := token

	for key := range token {
		fmt.Println(key)
		if key%2 == 0 {
			decodedToken[key] = decodedToken[key] - 2
		} else if key%3 == 0 {
			decodedToken[key] = decodedToken[key] + 4
		}
	}

	return decodedToken
}

func WrapPassWithSecretKey(password string) []byte {
	secretPassword := make([]byte, 64)

	lenOfPassword := len(password)
	lenOfSecret := len(secretKey)

	for i := 0; true; i++ {
		if i >= lenOfPassword && i >= lenOfSecret {
			break
		}

		if i >= lenOfSecret && i <= lenOfPassword {
			secretPassword[i] = password[i]
		} else if i >= lenOfPassword && i <= lenOfSecret {
			secretPassword[i] = secretKey[i]
		} else if i%2 == 0 {
			secretPassword[i] = password[i] + secretKey[i]
		} else {
			secretPassword[i] = secretKey[i] + password[i]
		}
	}

	return secretPassword
}

func GenerateHash(password string) string {
	hash := md5.Sum(WrapPassWithSecretKey(password))
	return hex.EncodeToString(hash[:])
}
