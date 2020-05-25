package main

import (
	"crypto/rand"
	"log"
	"os"

	"github.com/gargakshit/paperplane-server/utils"
	"golang.org/x/crypto/nacl/box"
)

func main() {
	_, errPriv := os.Stat("./keys/server_key_private")
	_, errPub := os.Stat("./keys/server_key_public")
	_, errPrivBase := os.Stat("./keys/server_key_private_base")
	_, errPubBase := os.Stat("./keys/server_key_public_base")

	if os.IsNotExist(errPriv) || os.IsNotExist(errPub) || os.IsNotExist(errPrivBase) || os.IsNotExist(errPubBase) {
		log.Println("Generating keys for the server...")

		serverPublicKey, serverPrivateKey, err := box.GenerateKey(rand.Reader)
		if err != nil {
			log.Fatalln("Error generating keys:", err.Error())
		}

		serverPublicKeyBase, serverPrivateKeyBase := utils.ToBase64(serverPublicKey[:]), utils.ToBase64(serverPrivateKey[:])

		publicKeyFile, err := os.Create("./keys/server_key_public")
		if err != nil {
			log.Fatalln("Error creating file:", err.Error())
		}

		privateKeyFile, err := os.Create("./keys/server_key_private")
		if err != nil {
			log.Fatalln("Error creating file:", err.Error())
		}

		publicKeyBaseFile, err := os.Create("./keys/server_key_public_base64")
		if err != nil {
			log.Fatalln("Error creating file:", err.Error())
		}

		privateKeyBaseFile, err := os.Create("./keys/server_key_private_base64")
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

		_, err = publicKeyBaseFile.Write([]byte(serverPublicKeyBase))
		if err != nil {
			log.Fatalln("Error writing to file:", err.Error())
		}

		_, err = privateKeyBaseFile.Write([]byte(serverPrivateKeyBase))
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

		err = publicKeyBaseFile.Close()
		if err != nil {
			log.Fatalln("Error closing file:", err.Error())
		}

		err = privateKeyBaseFile.Close()
		if err != nil {
			log.Fatalln("Error closing file:", err.Error())
		}

		log.Println("Done! It is recommended to backup the keys/ folder regularly")
	} else {
		log.Println("Keys already exist. remove old keys, please run \"make clean_keys\"")
	}
}
