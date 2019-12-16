package crypto

import (
	"core/logger"
	"crypto/md5"
	"encoding/hex"
	"os"
)

var log = logger.Logger{Context: "CRYPTO"}
var secretKey = []byte(os.Getenv("SECRET_KEY_FOR_PASSWORD"))
var secretToken = []byte(os.Getenv("SECRET_KEY_FOR_TOKEN"))

func getSalt(token []byte) byte {
	var tok byte

	lengthOfToken := len(token)
	lastByte := token[lengthOfToken-1]
	preLastByte := token[lengthOfToken-2]

	if lastByte == preLastByte {
		tok = lastByte + preLastByte/2
	} else if lastByte > preLastByte {
		tok = lastByte - preLastByte
	} else {
		tok = preLastByte - lastByte
	}

	if tok == 1 {
		tok += 2
	}

	if tok != 255 {
		tok += 1
	}

	return tok / 2
}

func mockBytes(token []byte) []byte {
	mockToken := make([]byte, len(token)+len(secretToken))
	iterToken := 0
	iterSecret := 0

	for key := range mockToken {
		if key%2 == 0 {
			if iterToken > len(token)-1 {
				break
			}

			mockToken[key] = token[iterToken]
			iterToken++
		} else if key%3 == 0 || key%5 == 0 {
			if iterSecret > len(secretToken) {
				iterSecret = 0
			}

			mockToken[key] = secretToken[iterSecret]
			iterSecret++
		} else {
			if iterSecret > len(secretToken) {
				iterSecret = 0
			}
			if iterToken > len(token)-1 {
				break
			}
			additionalToken := token[iterToken] + (secretToken[iterSecret] / 2)
			if additionalToken == 0 {
				additionalToken = token[iterToken] - (secretToken[iterSecret] / 2)
			}

			mockToken[key] = additionalToken
			iterSecret++
			iterToken++
		}
	}

	lengthOfEmptyBytes := 0

	for _, val := range mockToken {
		if val == 0 {
			lengthOfEmptyBytes++
		}
	}

	return mockToken[:len(mockToken)-lengthOfEmptyBytes]
}

func omitEmptyBytes(bytes []byte) []byte {
	newBytes := make([]byte, 0, len(bytes))

	for _, val := range bytes {
		if val != 0 {
			newBytes = append(newBytes, val)
		}
	}

	return newBytes
}

func unmockBytes(token []byte) []byte {
	unmockToken := make([]byte, len(token))
	iterSecret := 0

	for key := range unmockToken {
		if key%2 == 0 {
			unmockToken[key] = token[key]
		} else if key%3 == 0 || key%5 == 0 {
			if iterSecret > len(secretToken) {
				iterSecret = 0
			}
			unmockToken[key] = 0
			iterSecret++
		} else {
			if iterSecret > len(secretToken) {
				iterSecret = 0
			}
			additionalToken := (token[key] - (secretToken[iterSecret] / 2)) + secretToken[iterSecret]/2
			if additionalToken == 0 {
				additionalToken = token[key] + (secretToken[iterSecret] / 2)
			} else {
				additionalToken = token[key] - (secretToken[iterSecret] / 2)
			}
			unmockToken[key] = additionalToken
			iterSecret++
		}
	}

	return omitEmptyBytes(unmockToken)
}

func encodeByte(token *byte, key int, additionalToken byte) {
	if key%2 == 0 {
		if additionalToken+2 != 0 {
			*token = additionalToken + 2
		}
	} else if key%3 == 0 {
		if additionalToken-4 != 0 {
			*token = additionalToken - 4
		}
	} else {
		if additionalToken-1 != 0 {
			*token = additionalToken + 1
		}
	}
}

func decodeByte(token *byte, key int, additionalToken byte) {
	if key%2 == 0 {
		*token = additionalToken - 2
	} else if key%3 == 0 {
		*token = additionalToken + 4
	} else {
		*token = additionalToken - 1
	}
}

func isOutOfRangeOfByte(additionalToken byte) bool { return additionalToken == 0 }

func EncodeToken(token []byte, err interface{}) []byte {
	if err != nil {
		panic(err)
	}
	secretSalt := getSalt(token)

	encodedToken := token

	for key := range token {
		additionalToken := encodedToken[key] + secretSalt
		if isOutOfRangeOfByte(additionalToken) {
			additionalToken = encodedToken[key] - secretSalt
		}
		encodeByte(&encodedToken[key], key, additionalToken)
	}

	encodedToken = append(encodedToken, secretSalt-10)

	return mockBytes(encodedToken)
}

func DecodeToken(token []byte) []byte {
	decodedToken := unmockBytes(token)
	secretSalt := decodedToken[len(decodedToken)-1] + 10
	decodedToken = decodedToken[:len(decodedToken)-1]

	for key := range decodedToken {
		additionalToken := decodedToken[key] - secretSalt
		if isOutOfRangeOfByte(additionalToken) {
			additionalToken = decodedToken[key] + secretSalt
		}
		decodeByte(&decodedToken[key], key, additionalToken)
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
