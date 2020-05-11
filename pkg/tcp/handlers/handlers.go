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

	serverPubKey, serverPrivKey, err := box.GenerateKey(rand.Reader)

	if err != nil {
		log.Println("NaCl error:", err)

		conn.Write([]byte("Internal Error, closing connection!"))
		conn.Close()
	}

	serverBase64PubKey := utils.ToBase64(serverPubKey[:])
	serverBase64PrivKey := utils.ToBase64(serverPrivKey[:])

	log.Println(serverBase64PrivKey, serverBase64PubKey)

	conn.Write(serverPubKey[:])

	var buff []byte

	// stage is a variable used to store the state for the TCP handshake process
	// 0 means the server has sent it's public key and is awaiting the client's public key
	// 1 means the client's key has been recieved and a secure communication channel has been established, and the server is awaiting the client's ID
	// 2 means the client has been sent a challange and is awaiting the response. If the challange fails, close the connection, else continue to stage 3
	// 3 means the client has been verified and is ready to send / recieve messages over the encrypted channel
	stage := 0

	for {
		switch stage {
		case 0:
			conn.Read(buff)
			stage++
		default:
			conn.Close()
		}
	}
}
