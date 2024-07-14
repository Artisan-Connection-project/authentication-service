package postgres

import "golang.org/x/crypto/bcrypt"

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

func (h *bcryptHasherImple) Compare(hashedPassword string, password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
	return err == nil
}
