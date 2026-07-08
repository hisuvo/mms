package auth

import (
	"crypto/rand"
	"math/big"
)

type GenerateCode interface {
	GenerateToken(prefiex string,length int) (string, error)
}

type CodeGenerator struct{}

func NewCodeGenerator() GenerateCode {
	return &CodeGenerator{}
}

const charset = "ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

func (g *CodeGenerator) GenerateToken(prefiex string,length int) (string, error) {
	b := make([]byte, length)
	for i := range b {
		num, err := rand.Int(rand.Reader, big.NewInt(int64(len(charset))))
		if err != nil {
			return "", err
		}
		b[i] = charset[num.Int64()]
	}

	if prefiex != "" {
		return prefiex +"-" + string(b), nil
	}
	return string(b), nil
}

// b[i] = charset[rand.Int(rand.Reader, big.NewInt(int64(len(charset))))]