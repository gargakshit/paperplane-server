package tcp

import (
	"log"
	"net"
	"sync"

	"github.com/gargakshit/paperplane-server/pkg/tcp/handlers"
)

// BootstrapTCP bootstraps the TCP server
func BootstrapTCP(listenAddress string, wg *sync.WaitGroup) {
	defer wg.Done()

	server, err := net.Listen("tcp", listenAddress)

	if err != nil {
		log.Fatalln(err.Error())
	} else {
		log.Println("TCP server is listening at", listenAddress)
	}

	defer func() {
		err = server.Close()

		if err != nil {
			log.Println("Error shutting down the http server, forcing\n", err)
		}
	}()

	for {
		conn, err := server.Accept()

		if err != nil {
			log.Println("Error accepting TCP connection:", err.Error())
		}

		go handlers.HandleTCPClient(conn)
	}
}
