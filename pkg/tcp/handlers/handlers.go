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

	conn.Write([]byte(serverBase64PubKey))

	sharedEmphKeys := new([32]byte)

	buffer := make([]byte, 131072) // 128kb

	// stage is a variable used to store the state for the TCP handshake process
	// 0 means the server has sent it's public key and is awaiting the client's public key
	// 1 means the client's key has been recieved and a secure communication channel has been established, and the server is awaiting the client's ID
	// 2 means the client has been sent a challange and is awaiting the response. If the challange fails, close the connection, else continue to stage 3
	// Challange's response is also encrypted with the server's master key, so verifying it will also verify the server's identity
	// 3 means the client has been verified and is ready to send / recieve messages over the encrypted channel
	stage := 0

	for {
		switch stage {
		case 0:
			size, err := conn.Read(buffer)

			if err == nil {
				data := buffer[:size]

				if utils.IsBase64Valid(string(data)) {
					if peerKeyBytes, err := utils.FromBase64(string(data)); err != nil {
						conn.Close()
					} else {
						peerKey := new([32]byte)
						copy(peerKeyBytes[:], data[:32])

						box.Precompute(sharedEmphKeys, peerKey, serverPrivKey)

						stage++
					}
				} else {
					conn.Close()
				}
			} else {
				conn.Close()
			}
		default:
			conn.Close()
		}
	}
}
