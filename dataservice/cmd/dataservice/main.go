package main

import (
	"os"
	"os/signal"
	"syscall"

	"github.com/aclk/goblog/dataservice/cmd"
	"github.com/aclk/goblog/dataservice/internal/app/dbclient"
	"github.com/aclk/goblog/dataservice/internal/app/service"
	"github.com/alexflint/go-arg"
	"github.com/sirupsen/logrus"
)

var appName = "dataservice"

func main() {
	logrus.SetFormatter(&logrus.JSONFormatter{})
	logrus.Infof("Starting %v\n", appName)

	// Initialize config struct and populate it froms env vars and flags.
	cfg := cmd.DefaultConfiguration()
	arg.MustParse(cfg)

	server := service.NewServer(service.NewHandler(dbclient.NewGormClient(cfg)), cfg)
	server.SetupRoutes()

	handleSigterm(func() {
		logrus.Infoln("Captured Ctrl+C")
		server.Close()
	})
	server.Start()
}

// Handles Ctrl+C or most other means of "controlled" shutdown gracefully. Invokes the supplied func before exiting.
func handleSigterm(handleExit func()) {
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	signal.Notify(c, syscall.SIGTERM)
	go func() {
		<-c
		handleExit()
		os.Exit(1)
	}()
}
