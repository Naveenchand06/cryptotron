package main

import (
	"fmt"

	"github.com/Naveenchand06/cryptotron/decrypt"
	"github.com/Naveenchand06/cryptotron/encrypt"
)

func main() {
	original := "Shiva Shiva"
	fmt.Println("Original: ", original)
	encrypted := encrypt.Nimbus(original)
	fmt.Println("Encrypted: ", encrypted)
	decrypted := decrypt.Dimbus(encrypted)
	fmt.Println("Decrypted: ", decrypted)

}
