package main

import (
	"ciphers/cipher"
	"fmt"
)

func main() {
	message := "atacar base sul"
	c := cipher.Caesar_C(message)
	fmt.Println("cripo: ", c)
	fmt.Println("descripto: ", cipher.Caesar_D(c))
}
