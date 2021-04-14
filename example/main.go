package main

import (
	"crowdeAesCbc256"
	"fmt"
)

func main() {
	t := new(crowdeAesCbc256.AesEncrypt)
	t.Init("ijzh84t1w9xa56s9", "4bd393e7a457f9023d9ba95fffb5a2e1")
	encrypted := t.AESEncrypt("tes_enkripsi")
	fmt.Printf("encrypt: %s \n", encrypted)

	decrypt := t.AESDecrypt(encrypted)
	fmt.Printf("decrypt: %s \n", decrypt)
}