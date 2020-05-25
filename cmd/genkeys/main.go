package main

import (
	"crypto/rand"
	"log"
	"os"

	"golang.org/x/crypto/nacl/box"
)

func main() {
	_, err := os.Stat("./keys/server_key_private")
	if os.IsNotExist(err) {
		log.Println("Generating keys for the server...")

		serverPublicKey, serverPrivateKey, err := box.GenerateKey(rand.Reader)
		if err != nil {
			log.Fatalln("Error generating keys:", err.Error())
		}

		publicKeyFile, err := os.Create("./keys/server_key_public")
		if err != nil {
			log.Fatalln("Error creating file:", err.Error())
		}

		privateKeyFile, err := os.Create("./keys/server_key_private")
		if err != nil {
			log.Fatalln("Error creating file:", err.Error())
		}

		_, err = publicKeyFile.Write(serverPublicKey[:])
		if err != nil {
			log.Fatalln("Error writing to file:", err.Error())
		}

		_, err = privateKeyFile.Write(serverPrivateKey[:])
		if err != nil {
			log.Fatalln("Error writing to file:", err.Error())
		}

		err = publicKeyFile.Close()
		if err != nil {
			log.Fatalln("Error closing file:", err.Error())
		}

		err = privateKeyFile.Close()
		if err != nil {
			log.Fatalln("Error closing file:", err.Error())
		}

		log.Println("Done! It is recommended to backup the keys/ folder regularly")
	} else {
		log.Println("Keys already exist. To generate new keys, please delete the old one from keys/")
	}
}
