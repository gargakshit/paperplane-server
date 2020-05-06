package main

import (
	"fmt"
	"log"
	"sync"

	"github.com/gargakshit/paperplane-server/config"
	"github.com/gargakshit/paperplane-server/pkg/http"
	"github.com/gargakshit/paperplane-server/pkg/tcp"
)

func main() {
	log.Println("Starting the server...")
	fmt.Println()

	config := config.GetDefaultConfig()

	httpAddress := fmt.Sprintf("%s:%d", config.HTTPConfig.ListenAddress, config.HTTPConfig.Port)
	tcpAddress := fmt.Sprintf("%s:%d", config.TCPConfig.ListenAddress, config.TCPConfig.Port)

	log.Printf("HTTP Server is configured to listen on %s\n", httpAddress)
	log.Printf("TCP Server is configured to listen on %s\n", tcpAddress)

	var wg sync.WaitGroup

	wg.Add(1)
	go tcp.BootstrapTCP(tcpAddress, &wg)
	fmt.Println()

	wg.Add(1)
	go http.BootstrapHTTP(httpAddress, &wg)
	fmt.Println()

	wg.Wait()
}
