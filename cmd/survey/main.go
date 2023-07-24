package main

import (
	"log"

	"github.com/AlecAivazis/survey/v2"
)

var testQuestions = []*survey.Question{
	{
		Name:      "name",
		Prompt:    &survey.Input{Message: "What is your name?"},
		Validate:  survey.Required,
		Transform: survey.Title,
	},
	{
		Name: "color",
		Prompt: &survey.Select{
			Message: "Choose a color:",
			Options: []string{"red", "blue", "green"},
			Default: "red",
		},
	},
	{
		Name:   "age",
		Prompt: &survey.Input{Message: "How old are you?"},
	},
}

func main() {
	answers := struct {
		Name          string
		FavoriteColor string `survey:"color"`
		Age           int
	}{}

	err := survey.Ask(testQuestions, &answers)
	if err != nil {
		log.Fatalf("failed to ask survey questions: %v", err)
	}

	log.Println(answers)
}
