package main

import (
	"flag"
	"log"
	"net"

	"github.com/ZAF07/telco/cmd/app"
	"github.com/ZAF07/telco/config"
	"github.com/soheilhy/cmux"
)

func main() {

	filePath := flag.String("config", "config.yml", "path to config file")
	flag.Parse()
	appConfig := config.InitConfig(*filePath)

	listener, err := net.Listen("tcp", appConfig.Port)
	if err != nil {
		log.Fatalf("🚨🚨🚨 server error: %=v 🚨🚨🚨", err)
	}
	mux := cmux.New(listener)
	httpListener := mux.Match(cmux.HTTP1())
	RPCListener := mux.Match(cmux.Any())

	telcoApp := app.NewTelcoApplication(httpListener, RPCListener, appConfig)
	telcoApp.InitApplication()
	telcoApp.Start()

	if mErr := mux.Serve(); mErr != nil {
		log.Fatalf("🚨🚨🚨mux error: %+v 🚨🚨🚨", mErr)
	}

}
