// Demonstrate using viper for app configuration.
package main

import (
	"log"

	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

func main() {
	conf := viper.New()

	// Environment variables.
	conf.SetEnvPrefix("demo") // DEMO_
	conf.BindEnv("server")    // DEMO_SERVER
	conf.BindEnv("port")      // DEMO_PORT
	conf.SetDefault("server", "demo.example.com")
	conf.SetDefault("port", 8080)

	// Config file.
	// NOTE: Env vars take precedence!
	// See precedence order: https://github.com/spf13/viper#why-viper
	conf.SetConfigName("config")
	conf.AddConfigPath(".")
	conf.AddConfigPath("/tmp")
	if err := conf.ReadInConfig(); err != nil {
		log.Println(err)
	}

	// Print initial app settings.
	printConf(conf)

	// Watch config file and run callback function on changes.
	conf.WatchConfig()
	conf.OnConfigChange(func(e fsnotify.Event) {
		if e.Op == fsnotify.Write {
			printConf(conf)
		}
	})

	// Wait forever for receive on chan (exit with Ctrl-C).
	<-make(chan bool)
}

// printConf prints the current app configuration settings.
func printConf(conf *viper.Viper) {
	log.Printf("server: %s, port: %d\n",
		conf.GetString("server"),
		conf.GetInt("port"),
	)
}
