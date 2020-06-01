package main

import (
	"fmt"
	"log"
	"os"
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
	_, errPriv := os.Stat("./keys/server_key_private")
	_, errPub := os.Stat("./keys/server_key_public")
	_, errPrivBase := os.Stat("./keys/server_key_private_base64")
	_, errPubBase := os.Stat("./keys/server_key_public_base64")

	if !os.IsNotExist(errPriv) && !os.IsNotExist(errPub) && !os.IsNotExist(errPrivBase) && !os.IsNotExist(errPubBase) {

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

		// log.Println("Trying to connect to RethinkDB at", cfg.DatabaseConfig.RethinkDBConfig.Address)
		// database.ConnectToRethink(
		// 	cfg.DatabaseConfig.RethinkDBConfig.Address,
		// 	cfg.DatabaseConfig.RethinkDBConfig.Database,
		// )

		defer database.DisconnectMongo()

		log.Println("Trying to connect to MongoDB at", cfg.DatabaseConfig.MongoDBConfig.URI)
		database.ConnectToMongo(cfg.DatabaseConfig.MongoDBConfig.URI)

		var wg sync.WaitGroup

		wg.Add(1)
		go tcp.BootstrapTCP(tcpAddress, &wg)
		fmt.Println()

		wg.Add(1)
		go http.BootstrapHTTP(httpAddress, &wg)
		fmt.Println()

		wg.Wait()

	} else {
		log.Fatalln("Server keys not found, please run \"make genkeys\"")
	}
}
