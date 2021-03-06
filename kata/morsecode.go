package kata

import (
	"fmt"
	"strings"
)

//DecodeMorse is a simple morse decoder
func DecodeMorse(morseCode string) string {
	words := strings.Split(morseCode, "   ")
	res := ""
	for _, word := range words {
		letter := strings.Split(word, " ")
		for _, l := range letter {
			res = res //+ MORSE_CODE[l]
			fmt.Println(l)
		}
		res = res + " "
	}
	return strings.TrimSpace(res)
}
