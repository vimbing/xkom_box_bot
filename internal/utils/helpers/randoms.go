package helpers

import (
	"math/rand"
	"strings"
	"time"

	"github.com/zhangyunhao116/fastrand"
)

func RandomizeSeed() int64 {
	return time.Now().UnixNano()*int64(fastrand.Int()) - int64(fastrand.Int()) + int64(fastrand.Int())/2
}

func RandomInt(min int, max int) int {
	defer HandlePanic()

	seed := rand.NewSource(RandomizeSeed())

	if max < 0 || max < min {
		return min
	}

	randInt := rand.New(seed)

	return randInt.Intn((max+1)-min) + min
}

type CaseFormat string

type RandStringOptions struct {
	CaseFormat CaseFormat
}

const (
	UPPER_CASE CaseFormat = "upper"
	LOWER_CASE CaseFormat = "lower"
)

func RandomString(strLen int, normalChars bool, numbers bool, specials bool, opts ...RandStringOptions) string {
	usedAlphabet := []rune{}

	if normalChars {
		numbers := []rune{'q', 'w', 'e', 'r', 't', 'y', 'u', 'i', 'o', 'p', 'a', 's', 'd', 'f', 'g', 'h', 'j', 'k', 'l', 'z', 'x', 'c', 'v', 'b', 'n', 'm'}
		usedAlphabet = append(usedAlphabet, numbers...)
	}

	if numbers {
		numbers := []rune{'1', '2', '3', '4', '5', '6', '7', '8', '9', '0'}
		usedAlphabet = append(usedAlphabet, numbers...)
	}

	if specials {
		specials := []rune{'!', '$', '%', '.', ','}
		usedAlphabet = append(usedAlphabet, specials...)
	}

	randString := ""

	for i := 0; i < strLen; i++ {
		randString += string(usedAlphabet[RandomInt(0, len(usedAlphabet)-1)])
	}

	if len(opts) > 0 {
		o := opts[0]

		switch o.CaseFormat {
		case UPPER_CASE:
			randString = strings.ToUpper(randString)
		case LOWER_CASE:
			randString = strings.ToLower(randString)
		default:
		}
	}

	return randString
}
