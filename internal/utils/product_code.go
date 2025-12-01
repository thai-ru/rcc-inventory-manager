package utils

import (
	"fmt"
	"math/rand"
	"time"
)

func GenerateProductCode(prefix string) string {
	rand.Seed(time.Now().UnixNano())
	number := rand.Intn(900000) + 100000 //6-digit number
	return fmt.Sprintf("%s-%d", prefix, number)
}
