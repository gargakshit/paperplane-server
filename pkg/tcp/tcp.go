package tcp

import (
	"log"
	"net"
	"sync"
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

	defer server.Close()
}
