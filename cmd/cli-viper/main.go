// Demonstrate the use of the viper package for configuration.
package main

import (
	"fmt"
	"log"

	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

type Config struct {
	Server string
	Port   uint
}

func NewViper() *viper.Viper {
	v := viper.New()

	// Environment variables.
	v.SetEnvPrefix("demo") // DEMO_
	v.BindEnv("server")    // DEMO_SERVER
	v.BindEnv("port")      // DEMO_PORT
	v.SetDefault("server", "demo.example.com")
	v.SetDefault("port", 8080)

	// Config file.
	// NOTE: Env vars take precedence!
	// See precedence order: https://github.com/spf13/viper#why-viper
	v.SetConfigName("config")
	v.AddConfigPath(".")
	v.AddConfigPath("/tmp")

	return v
}

func main() {
	v := NewViper()

	if err := v.ReadInConfig(); err != nil {
		log.Println(err)
	}

	var conf Config
	if err := v.Unmarshal(&conf); err != nil {
		log.Println(err)
	}

	// Print initial app settings.
	fmt.Printf("%#v\n", &conf)
	printConf(v)

	// Watch config file and run callback function on changes.
	v.WatchConfig()
	v.OnConfigChange(func(e fsnotify.Event) {
		if e.Op == fsnotify.Write {
			printConf(v)
		}
	})

	// Wait forever (exit with Ctrl-C).
	select {}
}

func printConf(v *viper.Viper) {
	log.Printf(
		"server: %s, port: %d\n",
		v.GetString("server"),
		v.GetInt("port"),
	)
}
