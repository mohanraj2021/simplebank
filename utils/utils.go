package utils

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
)

const str = "abcdefghijklmnopqrstuvwxyz"

func RandomINT(min, max int64) int64 {
	rand.Seed(time.Now().UnixNano())
	randv := rand.Int63n(max - min + int64(rand.Intn(int(max))))
	val := min + randv
	return val
}

func RandomAccountName(n int) string {
	var sb strings.Builder
	k := len(str)
	for i := 0; i <= n; i++ {
		sb.WriteByte(str[rand.Intn(k)])
	}
	return sb.String()
}

func RandomCurrency() string {
	cur := []string{"USD", "INR", "PAK,"}
	lc := len(cur)
	return cur[rand.Intn(lc)]
}

func RandomString(n int) string {
	var sb strings.Builder
	k := len(str)
	for i := 0; i <= n; i++ {
		sb.WriteByte(str[rand.Intn(k)])
	}
	return sb.String()
}

func RandomAccountEmail(n int) string {
	var sb strings.Builder
	k := len(str)
	for i := 0; i <= n; i++ {
		sb.WriteByte(str[rand.Intn(k)])
	}
	return fmt.Sprintf("%s@mail.com", sb.String())
}
