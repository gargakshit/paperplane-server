package handlers

import (
	"crypto/rand"
	"log"
	"net"

	"github.com/gargakshit/paperplane-server/utils"
	"golang.org/x/crypto/nacl/box"
)

// HandleTCPClient is the base TCP connection handler, used mainly as a goroutine
func HandleTCPClient(conn net.Conn) {
	defer conn.Close()

	conn.Write([]byte("PaperPlane v2\n"))

	serverPubKey, serverPrivKey, err := box.GenerateKey(rand.Reader)

	if err != nil {
		log.Println("NaCl error:", err)

		conn.Write([]byte("Internal Error, closing connection!"))
		conn.Close()
	}

	serverBase64PubKey := utils.ToBase64(serverPubKey[:])
	serverBase64PrivKey := utils.ToBase64(serverPrivKey[:])

	log.Println(serverBase64PrivKey)

	conn.Write([]byte(serverBase64PubKey))
}
