package cipher

import (
	"strings"
)

func findIndex(char string, alphabet [27]string) int {
	for n := 0; n < len(alphabet); n++ {
		if alphabet[n] == char {
			return n
		}
	}

	return 404

}

// base 4
func Caesar_C(message string) string {
	alphabet := [...]string{"a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "k", "l", "m", "n",
		"o", "p", "q", "r", "s", "t", "u", "v", "w", "x", "y", "z", " "}
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
	alphabet := [...]string{"a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "k", "l", "m", "n",
		"o", "p", "q", "r", "s", "t", "u", "v", "w", "x", "y", "z", " "}
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
