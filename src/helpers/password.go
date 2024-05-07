package helpers

import "golang.org/x/crypto/bcrypt"

func Hash(raw string) string {
	hash, err := bcrypt.GenerateFromPassword([]byte(raw), bcrypt.DefaultCost)

	if err != nil {
		panic(err)
	}

	return string(hash)
}

func VerifyPassword(hash string, raw string) error {
	return bcrypt.CompareHashAndPassword([]byte(hash), []byte(raw))
}
