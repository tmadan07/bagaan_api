package util

import (
	"database/sql"
	"fmt"
	"math/rand"
	"strings"
	"time"
)

const alphabet = "abcdefghijklmnopqrstuvwxyz"
const numbers = "0123456789"

func init() {
	rand.Seed(time.Now().UnixNano())
}

// RandomInt generates a random integer between min and max
func RandomInt64(min, max int64) int64 {
	return min + rand.Int63n(max-min+1)
}

// RandomInt generates a random integer between 0 and max
func RandomInt32(min int32, max int32) int32 {
	return min + rand.Int31n(max-min+1)
}

//returns time
func RandomTime() time.Time {
	return time.Now()
}

// RandomHttpStatus generates a random http status code
func RandomHttpStatus() int32 {
	return 200
}

// RandomString generates a random string of length n
func RandomString(n int) string {
	var sb strings.Builder
	k := len(alphabet)

	for i := 0; i < n; i++ {
		c := alphabet[rand.Intn(k)]
		sb.WriteByte(c)
	}

	return sb.String()
}

func RandomNumericString(n int) string {
	var sb strings.Builder
	k := len(numbers)

	for i := 0; i < n; i++ {
		c := numbers[rand.Intn(k)]
		sb.WriteByte(c)
	}

	return sb.String()
}

// RandomNullString generates a random nullable string of length n
func RandomNullString(n int) sql.NullString {
	var sb strings.Builder
	k := len(alphabet)

	for i := 0; i < n; i++ {
		c := alphabet[rand.Intn(k)]
		sb.WriteByte(c)
	}

	return sql.NullString{String: sb.String(), Valid: true}
}

// RandomOwner generates a random owner name
func RandomOwner() string {
	return RandomString(6)
}

// RandomMoney generates a random amount of money
func RandomMoney() int64 {
	return RandomInt64(0, 1000)
}

// RandomCurrency generates a random currency code
func RandomCurrency() string {
	return "USD"
}

// RandomEmail generates a random email
func RandomEmail() string {
	return fmt.Sprintf("%s@email.com", RandomString(6))
}
