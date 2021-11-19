package caesar

import (
	"golang.org/x/text/runes"
	"golang.org/x/text/transform"
	"golang.org/x/text/unicode/norm"
	"unicode"
)

type Caesar struct {
	Key int
}

func (c *Caesar) RemoveSpecialChar(text string) string {
	t := transform.Chain(norm.NFD, runes.Remove(runes.In(unicode.Mn)), norm.NFC)
	result, _, _ := transform.String(t, text)

	return result
}

func (c *Caesar) Crypt(text string) string {
	var crypted string

	text = c.RemoveSpecialChar(text)

	for _, char := range text {
		crypted += c.rotateChar(uint8(char), c.Key)
	}

	return crypted
}

func (c *Caesar) Decrypt(cryptedText string) string {
	var crypted string

	cryptedText = c.RemoveSpecialChar(cryptedText)

	for _, char := range cryptedText {
		crypted += c.rotateChar(uint8(char), -c.Key)
	}

	return crypted
}

func (c *Caesar) getMinMaxMod(char uint8) (min, max, mod int) {
	if char >= 'A' && char <= 'Z' {
		min, max, mod = 'A', 'Z', 'Z' - 'A' + 1
	}

	if char >= 'a' && char <= 'z' {
		min, max, mod = 'a', 'z', 'z' - 'a' + 1
	}

	if char >= '0' && char <= '9' {
		min, max, mod = '0', '9', '9' - '0' + 1
	}

	return min, max, mod
}

func (c *Caesar) rotateChar(char uint8, key int) string {
	var rotated int

	min, max, mod := c.getMinMaxMod(char)

	if mod == 0 {
		return string(char)
	}

	if key > 0 {
		rotated = ((int(char) - min + key) % mod) + min
	} else {
		rotated = ((int(char) - max + key) % mod) + max
	}

	return string(rune(rotated))
}
