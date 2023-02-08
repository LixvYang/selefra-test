package main

import (
	"log"
	"os"
	"os/signal"
	"selefra-demo/internal/model"
	"selefra-demo/internal/router"
	"selefra-demo/internal/utils"
	"syscall"
)

func main() {
	utils.InitSetting.Do(utils.Init)
	model.InitDB.Do(model.InitDb)

	signalch := make(chan os.Signal, 1)

	router.InitRouter()

	signal.Notify(signalch, os.Interrupt, syscall.SIGTERM)
	signalType := <-signalch
	signal.Stop(signalch)
	//cleanup before exit
	log.Printf("On Signal <%s>", signalType)
	log.Println("Exit command received. Exiting...")
}
