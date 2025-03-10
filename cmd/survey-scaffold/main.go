package main

import (
	"fmt"
	"os"

	"github.com/AlecAivazis/survey/v2"
)

var questions = map[string][]*survey.Question{
	"type": {
		{
			Name: "TemplateType",
			Prompt: &survey.MultiSelect{
				Message: "Template type:",
				Options: []string{"service", "dashboard", "pipeline"},
				Default: "service",
				Help:    "The type of template to be used",
			},
		},
	},

	"service": {
		{
			Name: "Name",
			Prompt: &survey.Input{
				Message: "Service name:",
				Help:    "Name of the service",
			},
		},
		{
			Name: "Host",
			Prompt: &survey.Input{
				Message: "Host:",
				Default: "localhost",
				Help:    "The hostname for your service",
			},
		},

		{
			Name: "Port",
			Prompt: &survey.Input{
				Message: "Port number:",
				Default: "42",
				Help:    "A port number between 0-65535",
			},
		},
	},
}

type Conf struct {
	TemplateType []string
}

type Service struct {
	Name string
	Host string
	Port uint
}

func main() {
	var conf Conf

	err := survey.Ask(questions["type"], &conf, survey.WithKeepFilter(true))
	if err != nil {
		fmt.Printf("failed to complete type questions: %v\n", err)
		os.Exit(1)
	}

	for _, tt := range conf.TemplateType {
		if tt == "service" {
			var service Service
			err = survey.Ask(questions["service"], &service)
			if err != nil {
				fmt.Printf("failed to complete service questions: %v\n", err)
				os.Exit(1)
			}
			fmt.Println(service)
		}
	}

	fmt.Println(conf.TemplateType)
}
