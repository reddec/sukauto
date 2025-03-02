package main

import (
	"fmt"
	"github.com/jessevdk/go-flags"
	"log"
	"os"
	"sukauto/controler"
	"sukauto/integration"
	"sukauto/utils"
	"time"
)

var config struct {
	Bind          string                 `long:"bind" env:"BIND" description:"Binding address" default:":8080"`
	ConfigFile    string                 `long:"config-file" env:"CONFIG_FILE" description:"Path to configuration file" default:"config.json"`
	UpdCmd        string                 `long:"updcmd" env:"UPDCMD" description:"command for update" default:"git pull origin master"`
	CORS          integration.CorsConfig `group:"cors" env-namespace:"CORS" namespace:"cors"`
	CheckInterval time.Duration          `long:"check-interval" env:"CHECK_INTERVAL" description:"Background check interval" default:"15s"`
	StatusScript  string                 `long:"status-script" env:"STATUS_SCRIPT" description:"Script to run for services events"`
	// plugins
	Telegram integration.ExtraTelegram `group:"telegram plugin" env-namespace:"TG" namespace:"tg"`
}

func main() {
	_, err := flags.Parse(&config)
	if err != nil {
		os.Exit(1)
	}
	fmt.Println(utils.Logo)
	fmt.Println("SUKAUTO - monitoring system")
	monitor := controler.NewServiceControllerByPath(config.ConfigFile, config.UpdCmd)
	// setup listeners
	events := monitor.Events()
	events = controler.WithBackgroundCheck(events, config.CheckInterval, monitor)
	events = controler.WithStateFilter(events)
	if config.StatusScript != "" {
		events = controler.WithScriptRunner(events, config.StatusScript)
	}
	events, httpEvents := controler.Tee(events)
	events, tgEvents := controler.Tee(events)
	// ....
	go func() {
		for event := range events {
			log.Println(event.Name, event.Type.String())
		}
	}()
	// plugins
	go func() {
		if err := config.Telegram.Run(tgEvents); err != nil {
			log.Println("telegram plugin failed:", err)
		}
	}()

	// setup integration
	var access controler.Access = monitor
	router := integration.NewHTTP(monitor, access, config.CORS, httpEvents)

	panic(router.Run(config.Bind))
}
