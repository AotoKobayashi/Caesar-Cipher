package cipher

import (
	"strings"

	"github.com/AotoKobayashi/Xharacter"
)

func findIndex(char string, alphabet []string) int {
	for n := 0; n < len(alphabet); n++ {
		if alphabet[n] == char {
			return n
		}
	}

	return 404

}

// base 4
func Caesar_C(message string) string {
	alphabet := append(Xharacter.AlphabetMi(), " ")
	newmessage := strings.Split(message, "")

	for n := 0; n < len(newmessage); n++ {
		alptindex := findIndex(newmessage[n], alphabet) + 4

		if alptindex != 404 {
			if alptindex > 27 {
				alptindex = alptindex - 27
			}

			newmessage[n] = alphabet[alptindex]

		} else {
			return "error"
		}

	}

	return strings.Join(newmessage, "")
}

func Caesar_D(message string) string {
	alphabet := append(Xharacter.AlphabetMi(), " ")
	newmessage := strings.Split(message, "")

	for n := 0; n < len(newmessage); n++ {
		alptindex := findIndex(newmessage[n], alphabet) - 4

		if alptindex != 404 {
			if alptindex < 0 {
				alptindex = alptindex + 27
			}

			newmessage[n] = alphabet[alptindex]

		} else {
			return "error"
		}

	}

	return strings.Join(newmessage, "")
}
