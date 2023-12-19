package zcrypto

import "golang.org/x/crypto/bcrypt"

const DefaultCost = bcrypt.DefaultCost

func Hash(password string, cost int) (string, error) {
	if cost == 0 {
		cost = DefaultCost
	}
	fromPassword, err := bcrypt.GenerateFromPassword([]byte(password), cost)
	if err != nil {
		return "", err
	}

	return string(fromPassword), nil
}

func Compare(hashedPass, plainPass string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPass), []byte(plainPass))
}
