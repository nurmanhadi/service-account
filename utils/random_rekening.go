package utils

import (
	"fmt"
	"math/rand"
	"time"
)

func GenerateRandomRekening() string {
	source := rand.NewSource(time.Now().UnixNano())
	r := rand.New(source)
	prefix := "2001"
	randomPart := r.Intn(900000) + 100000
	return fmt.Sprintf("%s%d", prefix, randomPart)
}
