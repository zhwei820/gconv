package grand

import (
	"crypto/rand"
	"crypto/sha256"
	"fmt"
	"math"
	"math/big"
	"strings"
)

const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
const digitBytes = "01234567890"
const (
	letterIdxBits = 6                    // 6 bits to represent a letter index
	letterIdxMask = 1<<letterIdxBits - 1 // All 1-bits, as many as letterIdxBits
	letterIdxMax  = 63 / letterIdxBits   // # of letter indices fitting in 63 bits
)

func Intn(max int64) int64 {
	maxn := *big.NewInt(max)
	n, err := rand.Int(rand.Reader, &maxn)
	if err != nil {
		return 0
	}
	return n.Int64()
}

// RandStringBytesMaskImprSrcSB
// ref: https://stackoverflow.com/questions/22892120/how-to-generate-a-random-string-of-a-option_trade-length-in-go
func Letters(n int) string {
	sb := strings.Builder{}
	sb.Grow(n)
	// A Intn(math.MaxInt64) generates 63 random bits, enough for letterIdxMax characters!
	for i, cache, remain := n-1, Intn(math.MaxInt64), letterIdxMax; i >= 0; {
		if remain == 0 {
			cache, remain = Intn(math.MaxInt64), letterIdxMax
		}
		if idx := int(cache & letterIdxMask); idx < len(letterBytes) {
			sb.WriteByte(letterBytes[idx])
			i--
		}
		cache >>= letterIdxBits
		remain--
	}

	return sb.String()
}

// RandStringBytesMaskImprSrcSB
// ref: https://stackoverflow.com/questions/22892120/how-to-generate-a-random-string-of-a-option_trade-length-in-go
func Digits(n int) string {
	sb := strings.Builder{}
	sb.Grow(n)
	// A Intn(math.MaxInt64) generates 63 random bits, enough for letterIdxMax characters!
	for i, cache, remain := n-1, Intn(math.MaxInt64), letterIdxMax; i >= 0; {
		if remain == 0 {
			cache, remain = Intn(math.MaxInt64), letterIdxMax
		}
		if idx := int(cache & letterIdxMask); idx < len(digitBytes) {
			sb.WriteByte(digitBytes[idx])
			i--
		}
		cache >>= letterIdxBits
		remain--
	}

	return sb.String()
}

func HashString(s string, llen ...int) string {
	if len(llen) == 0 || llen[0] <= 0 {
		llen = []int{32} // default is 32
	}

	sha1Inst := sha256.New()
	sha1Inst.Write([]byte(s))
	return fmt.Sprintf("%x", sha1Inst.Sum([]byte("")))[:llen[0]]
}

func HashAndKeepRawPart(s string, llen ...int) string {
	ss := HashString(s, llen...)
	if len(s) > llen[0] {
		s = s[0:llen[0]]
	}
	return s + "-" + ss
}
