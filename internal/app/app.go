package app

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"log"
	"order/config"
)

func Run(configPath, configName string) {
	// Configuration
	if err := config.InitConfig(configPath, configName); err != nil {
		logrus.Fatalf("config init error: %s", err.Error())
	}

	cfg, err := config.NewConfig()
	if err != nil {
		log.Fatalf("Get config error: %s", err)
	}

	// Logger
	SetLogrus(cfg.Level)

	// Database
	logrus.Info("Initializing postgres...")
	pg, err := postgres.New(cfg.PG.URL, postgres.MaxPoolSize(cfg.PG.MaxPoolSize))
	if err != nil {
		log.Fatal(fmt.Errorf("app - Run - pgdb.NewServices: %s", err))
	}
	defer pg.Close()

	//// Repositories
	//logrus.Info("Initializing repositories...")
	//repositories := repo.NewRepositories(pg)
	//
	//// Services dependencies
	//logrus.Info("Initializing services...")
	//deps := service.ServicesDependencies{
	//	Repos:    repositories,
	//	Hasher:   hasher.NewSHA1Hasher(cfg.Hasher.Salt),
	//	SignKey:  cfg.JWT.SignKey,
	//	TokenTTL: cfg.JWT.TokenTTL,
	//}
	//services := service.NewServices(deps)
	//
	//// Echo handler
	//logrus.Info("Initializing handlers and routes...")
	//handler := echo.New()
	//// setup handler validator as lib validator
	//handler.Validator = validator.NewCustomValidator()
	//v1.NewRouter(handler, services)
	//
	//// HTTP server
	//logrus.Info("Initializing server...")
	//logrus.Debugf("Server port: %s", cfg.HTTP.Port)
	//httpServer := httpserver.New(handler, httpserver.Port(cfg.HTTP.Port))
	//// Waiting signal
	//logrus.Info("Configuring graceful shutdown...")
	//interrupt := make(chan os.Signal, 1)
	//signal.Notify(interrupt, os.Interrupt, syscall.SIGTERM)
	//
	//select {
	//case s := <-interrupt:
	//	logrus.Info("app - Run - signal: " + s.String())
	//case err = <-httpServer.Notify():
	//	logrus.Error(fmt.Errorf("app - Run - httpServer.Notify: %w", err))
	//}
	//
	//// Graceful shutdown
	//logrus.Info("Shutting down...")
	//err = httpServer.Shutdown()
	//if err != nil {
	//	logrus.Error(fmt.Errorf("app - Run - httpServer.Shutdown: %w", err))
	//}
}