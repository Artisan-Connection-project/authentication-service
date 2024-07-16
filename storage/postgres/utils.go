package postgres

import (
	"crypto/rand"
	"encoding/hex"
	"log"

	"golang.org/x/crypto/bcrypt"
)

type Hasher interface {
	Hash(password string) string
	Compare(hashedPassword, password string) bool
}

type bcryptHasherImple struct {
	cost int
}

func NewBcryptHasher(cost int) Hasher {
	return &bcryptHasherImple{cost: cost}
}

func (h *bcryptHasherImple) Hash(password string) string {

	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(password), h.cost)
	return string(hashedPassword)
}

func (h *bcryptHasherImple) Compare(hashedPassword, password string) bool {
	log.Println("Comparing passwords")
	log.Println("Hashed Password:", hashedPassword)
	log.Println("Plain Password:", password)
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
	if err != nil {
		log.Println("Password comparison error:", err)
	}
	return err == nil
}

func GenerateRandomToken() (string, error) {
	bytes := make([]byte, 16)
	if _, err := rand.Read(bytes); err != nil {
		return "", err
	}
	return hex.EncodeToString(bytes), nil
}
