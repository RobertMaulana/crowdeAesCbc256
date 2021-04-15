package main

import (
	"crowdeAesCbc256"
	"encoding/base64"
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

func main() {
	/*============================================================
		test with plain text
	============================================================*/
	t := new(crowdeAesCbc256.AesEncrypt)
	t.Init("ijzh84t1w9xa56s9", "4bd393e7a457f9023d9ba95fffb5a2e1")
	//encrypted := t.AESEncrypt("tes_enkripsi")
	//fmt.Printf("encrypt: %s \n", encrypted)
	//
	//decrypt := t.AESDecrypt(encrypted)
	//fmt.Printf("decrypt: %s \n", decrypt)

	/*============================================================
		test with json file
	============================================================*/

	// get current dir path
	pwd, _ := os.Getwd()

	// open json file
	jsonFile, err := os.Open(pwd + "/data/pengajuan_pinjaman.json")
	if err != nil {
		log.Fatal("unable to open json file %#v \n", err)
	}

	defer jsonFile.Close()

	byteValue, _ := ioutil.ReadAll(jsonFile)
	encrypted := t.AESEncrypt(string(byteValue))

	// this encrypted that will send to OJK
	concatWithIv := encrypted + "::" + t.Iv

	// they want encrypted data with base 64
	encoded := base64.StdEncoding.EncodeToString([]byte(concatWithIv))
	fmt.Printf("this data will send to OJK %s \n", encoded)

	// place encrypted data to output folder
	f, err := os.Create(pwd + "/output/pengajuan_pinjaman.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	_, err2 := f.WriteString(encoded)
	if err2 != nil {
		log.Fatal(err2)
	}

	fmt.Println("done!")

	//decrypt := t.AESDecrypt(encrypted)
	//fmt.Printf("decrypt: %s \n", decrypt)
}