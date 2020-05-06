package main

import (
	"fmt"
	"log"
	"regexp"
	"sync"

	"github.com/gargakshit/paperplane-server/config"
	"github.com/gargakshit/paperplane-server/database"
	"github.com/gargakshit/paperplane-server/pkg/http"
	"github.com/gargakshit/paperplane-server/pkg/tcp"
	"github.com/gargakshit/paperplane-server/utils"
)

func main() {
	log.Println("Starting the server...")

	config := config.GetDefaultConfig()

	httpAddress := fmt.Sprintf("%s:%d", config.HTTPConfig.ListenAddress, config.HTTPConfig.Port)
	tcpAddress := fmt.Sprintf("%s:%d", config.TCPConfig.ListenAddress, config.TCPConfig.Port)

	log.Println("Compiling the RegExp(s)")
	base64RegEx, err := regexp.Compile(`^(?:[A-Za-z0-9+\/]{2}[A-Za-z0-9+\/]{2})*(?:[A-Za-z0-9+\/]{2}==|[A-Za-z0-9+\/]{3}=)?$`)
	if err != nil {
		log.Fatalln(err.Error())
	}
	utils.Base64Regex = base64RegEx
	log.Println("RegExp(s) compiled")

	log.Println("Trying to connect to RethinkDB at", config.DatabaseConfig.RethinkDBConfig.Address)
	database.ConnectToRethink(
		config.DatabaseConfig.RethinkDBConfig.Address,
		config.DatabaseConfig.RethinkDBConfig.Database,
	)

	var wg sync.WaitGroup

	wg.Add(1)
	go tcp.BootstrapTCP(tcpAddress, &wg)
	fmt.Println()

	wg.Add(1)
	go http.BootstrapHTTP(httpAddress, &wg)
	fmt.Println()

	wg.Wait()
}
