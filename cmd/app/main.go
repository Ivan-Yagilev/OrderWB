package main

import "order/internal/app"

const (
	configPath = "config"
	configName = "config"
)

func main() {
	app.Run(configPath, configName)
}
