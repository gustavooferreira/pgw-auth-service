package main

import (
	"fmt"
	"io/ioutil"
	"os"

	"github.com/gustavooferreira/pgw-auth-service/pkg/api"
	"github.com/gustavooferreira/pgw-auth-service/pkg/core"
	"github.com/gustavooferreira/pgw-auth-service/pkg/core/log"
	"github.com/gustavooferreira/pgw-auth-service/pkg/core/repository"
	"github.com/gustavooferreira/pgw-auth-service/pkg/lifecycle"
)

func main() {
	retCode := mainLogic()
	os.Exit(retCode)
}

func mainLogic() int {
	// Setup logger
	logger := core.NewAppLogger(os.Stdout, log.INFO)
	defer logger.Sync()

	logger.Info("APP starting")

	// Read config
	logger.Info("reading configuration", log.Field("type", "setup"))
	config := core.NewConfig()
	if err := config.LoadConfig(); err != nil {
		logger.Error(err.Error(), log.Field("type", "setup"))
		return 1
	}

	// TODO: Set log level after reading config
	// something like this:
	// logger.SetLevel(config.Options.LogLevel)

	// Read credentials from file and pass data to credsHolder
	yamlContent, err := ioutil.ReadFile(config.Database.Filename)
	if err != nil {
		logger.Error(err.Error(), log.Field("type", "setup"))
		return 1
	}

	credsHolder := repository.NewCredsHolder()
	err = credsHolder.Load(yamlContent)
	if err != nil {
		logger.Error(err.Error(), log.Field("type", "setup"))
		return 1
	}

	server := api.NewServer(config.Webserver.Host, config.Webserver.Port, config.Options.DevMode, logger, &credsHolder)

	// Spawn SIGINT listener
	go lifecycle.TerminateHandler(logger, server)

	// Listen for incoming requests -- app blocks here
	logger.Info("listenning for incoming requests", log.Field("type", "setup"))
	err = server.ListenAndServe()
	if err != nil {
		logger.Error(fmt.Sprintf("unexpected error while serving HTTP: %s", err))
		return 1
	}

	logger.Info("APP gracefully terminated")
	return 0
}
