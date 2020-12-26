package service

import "golang.org/x/crypto/bcrypt"

// EncryptString ...
func EncryptString(s string) (string, error) {
	b, err := bcrypt.GenerateFromPassword([]byte(s), bcrypt.MinCost)
	if err != nil {
		return "", err
	}

	return string(b), nil
}

// EncryptCompare ...
func EncryptCompare(encrypt string, str string) bool {
	return bcrypt.CompareHashAndPassword([]byte(encrypt), []byte(str)) == nil
}
