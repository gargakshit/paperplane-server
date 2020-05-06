package tcp

import (
	"fmt"
	"log"
	"net"
	"sync"
)

// BootstrapTCP bootstraps the TCP server
func BootstrapTCP(listenAddress string, wg *sync.WaitGroup) {
	defer wg.Done()

	server, err := net.Listen("tcp", listenAddress)

	if err != nil {
		log.Fatalln(err)
	} else {
		log.Println("TCP server is listening at", listenAddress)
		fmt.Println()
	}

	defer server.Close()
}
