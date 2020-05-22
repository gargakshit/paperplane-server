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
	"github.com/joho/godotenv"
)

func main() {
	log.Println("Loading the config files...")
	err := godotenv.Load()
	if err != nil {
		log.Fatalln(err.Error())
	}
	log.Println("Config files loaded")

	log.Println("Starting the server...")

	cfg := config.GetDefaultConfig()
	config.GlobalConfig = &cfg

	httpAddress := fmt.Sprintf("%s:%d", cfg.HTTPConfig.ListenAddress, cfg.HTTPConfig.Port)
	tcpAddress := fmt.Sprintf("%s:%d", cfg.TCPConfig.ListenAddress, cfg.TCPConfig.Port)

	log.Println("Compiling the RegExp(s)")
	base64RegEx, err := regexp.Compile(`^(?:[A-Za-z0-9+\/]{2}[A-Za-z0-9+\/]{2})*(?:[A-Za-z0-9+\/]{2}==|[A-Za-z0-9+\/]{3}=)?$`)
	if err != nil {
		log.Fatalln(err.Error())
	}
	utils.Base64Regex = base64RegEx
	log.Println("RegExp(s) compiled")

	log.Println("Trying to connect to RethinkDB at", cfg.DatabaseConfig.RethinkDBConfig.Address)
	database.ConnectToRethink(
		cfg.DatabaseConfig.RethinkDBConfig.Address,
		cfg.DatabaseConfig.RethinkDBConfig.Database,
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
