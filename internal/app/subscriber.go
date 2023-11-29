package app

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/nats-io/stan.go"
	"github.com/sirupsen/logrus"
	"log"
	"order/config"
	v1 "order/internal/controller/http/v1"
	"order/internal/nats"
	"order/internal/repo"
	"order/internal/service"
	"order/pkg/hasher"
	"order/pkg/httpserver"
	"order/pkg/postgres"
	"order/pkg/validator"
	"os"
	"os/signal"
	"sync"
	"syscall"
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

	// Postgres
	logrus.Info("Initializing postgres...")
	pg, err := postgres.New(cfg.URL, postgres.MaxPoolSize(cfg.MaxPoolSize))
	if err != nil {
		log.Fatal(fmt.Errorf("app - Run - pgdb.NewServices: %s", err))
	}
	defer pg.Close()

	// Repositories
	logrus.Info("Initializing repositories...")
	repositories := repo.NewRepositories(pg)

	// Services dependencies
	logrus.Info("Initializing services...")
	deps := service.ServicesDependencies{
		Repos:    repositories,
		Hasher:   hasher.NewSHA1Hasher(cfg.Salt),
		SignKey:  cfg.SignKey,
		TokenTTL: cfg.TokenTTL,
	}
	services := service.NewServices(deps)

	valid := validator.New()
	// Connect to nats-streaming server
	natsStreaming := nats.NewNats(services, valid)

	sc, err := natsStreaming.Connect(
		"test-cluster",
		"subs-1",
		"localhost:4222",
	)
	if err != nil {
		return
	}
	defer func(sc stan.Conn) {
		err = sc.Close()
		if err != nil {
			logrus.Fatal(fmt.Errorf("error while closing connection to nats-streaming: %s", err))
		}
	}(sc)

	// Subs to the nats subj "orders"
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		err = natsStreaming.Subscribe(&wg, sc, "orders")
		if err != nil {
			return
		}
	}()
	logrus.Printf("Successfilly subscribed to nats-streaming subject orders")

	// Echo handler
	logrus.Info("Initializing handlers and routes...")
	handler := echo.New()
	// setup handler validator as lib validator
	//handler.Validator = validator.NewCustomValidator()
	v1.NewRouter(handler, services)

	// HTTP server
	logrus.Info("Initializing server...")
	logrus.Debugf("Server port: %s", cfg.Port)
	httpServer := httpserver.New(handler, httpserver.Port(cfg.Port))

	// Waiting signal
	logrus.Info("Configuring graceful shutdown...")
	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt, syscall.SIGTERM)

	select {
	case s := <-interrupt:
		logrus.Info("app - Run - signal: " + s.String())
	case err = <-httpServer.Notify():
		logrus.Error(fmt.Errorf("app - Run - httpServer.Notify: %w", err))
	}

	// Graceful shutdown
	logrus.Info("Shutting down...")
	err = httpServer.Shutdown()
	if err != nil {
		logrus.Error(fmt.Errorf("app - Run - httpServer.Shutdown: %w", err))
	}

	if err = sc.Close(); err != nil {
		logrus.Errorf("error while nats-streaming close connection: %s", err.Error())
	}

	wg.Wait()
}
