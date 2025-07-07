// Demonstrate the use of the viper package for configuration.
package main

import (
	"log"
	"runtime"

	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

type Config struct {
	Server string
	Port   uint
}

type App struct {
	v *viper.Viper

	config Config
}

func NewViper() *viper.Viper {
	v := viper.New()

	// Environment variables.
	v.SetEnvPrefix("demo") // DEMO_

	err := v.BindEnv("server") // DEMO_SERVER
	if err != nil {
		log.Fatalf("failed to bind DEMO_SERVER env var: %v", err)
	}
	err = v.BindEnv("port") // DEMO_PORT
	if err != nil {
		log.Fatalf("failed to bind DEMO_PORT env var: %v", err)
	}

	v.SetDefault("server", "demo.example.com")
	v.SetDefault("port", 8080)

	// Config file.
	// NOTE: Env vars take precedence!
	// See precedence order: https://github.com/spf13/viper#why-viper
	v.SetConfigName("config")
	v.AddConfigPath(".")

	return v
}

func main() {
	app := App{v: NewViper()}

	if err := app.v.ReadInConfig(); err != nil {
		log.Println(err)
	}

	if err := app.v.Unmarshal(&app.config); err != nil {
		log.Println(err)
	}

	app.logConfig()

	app.watchConfigAndLogChanges()

	// Wait forever (exit with Ctrl-C).
	runtime.Goexit()
}

func (app *App) logConfig() {
	log.Printf("server: %s, port: %d\n", app.v.GetString("server"), app.v.GetInt("port"))
}

// Watch config file and run callback function on changes.
func (app *App) watchConfigAndLogChanges() {
	app.v.WatchConfig()
	app.v.OnConfigChange(func(e fsnotify.Event) {
		if e.Op == fsnotify.Write {
			app.logConfig()
		}
	})
}
