package main

import (
	"fmt"
	"strings"
)

func main() {
	sample := "This will-rotatezzza"
	var k int32 = 1
	fmt.Println(caeserCipher(sample, k))

}

func caeserCipher(s string, k int32) string {
	var output []byte

	for _, c := range s {
		switch {
		case c >= 65 && c <= 90:
			{
				output = append(output, rotate(c, k, uppercaseAlphabet))
			}
		case c >= 97 && c <= 122:
			{
				output = append(output, rotate(c, k, lowercaseAlphabet))
			}
		default:
			{
				output = append(output, byte(c))
			}
		}
	}
	return string(output)
}

// 97 122
var lowercaseAlphabet string = "abcdefghijklmnopqrstuvwxyz"

// 65 90
var uppercaseAlphabet string = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"

func rotate(s rune, rotation int32, alphabet string) byte {
	idx := strings.IndexRune(alphabet, s)
	rotatedByte := (idx + int(rotation)) % len(alphabet)
	return alphabet[rotatedByte]
}
