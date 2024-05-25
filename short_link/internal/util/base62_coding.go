package util

import (
	"errors"
	"fmt"
	"math"
	"strings"
)

const (
	base         uint64 = 62
	characterSet        = "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"
)

var ErrInvalidToken = errors.New("error invalid token")

func ToBase62(num uint64) string {
	encoded := ""
	for num > 0 {
		r := num % base
		num /= base
		encoded = string(characterSet[r]) + encoded

	}
	return encoded
}

func FromBase62(encoded string) (uint64, error) {
	var val uint64
	for index, char := range encoded {
		pow := len(encoded) - (index + 1)
		pos := strings.IndexRune(characterSet, char)
		if pos == -1 {
			return 0, fmt.Errorf("Controller GetByToken token invalid character: "+string(char)+"error: %w", ErrInvalidToken)
		}

		val += uint64(pos) * uint64(math.Pow(float64(base), float64(pow)))
	}

	return val, nil
}
