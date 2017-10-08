package main

import (
	"os"
	"github.com/urfave/cli"
	"gopkg.in/resty.v1"
	"github.com/fatih/color"
)

const ENDPOINT = "https://weekly-share.istanbulphp.org/api/v1/links"

type Link struct {
	Url   string `json:"url"`
	Email string `json:"email"`
}

func main() {
	app := cli.NewApp()
	app.Name = "weekly"
	app.Usage = "u https://emirkarsiyakali.com/php-ve-visual-debt-goruntu-kirliligi-ecbb5267e412"
	app.Version = "0.1.0"

	app.Commands = []cli.Command{
		{
			Name:    "url",
			Aliases: []string{"u"},
			Usage:   "Fully qualified URL you want to share with weekly curators. Ex: https://google.com.tr/",
			Action: func(c *cli.Context) error {
				resp, err := resty.R().
					SetBody(Link{Url: c.Args().First(), Email: findEnvironmentVariableByKey("EMAIL")}).
					Post(ENDPOINT)

				if err != nil {
					color.Red("Something went wrong while sharing URL. " +
						"Please try again and be sure you specified fully qualified URL like https://google.com.tr/.")
				}

				if resp.StatusCode() == 201 {
					color.Green("Thank you!")
				}

				return nil
			},
		},
	}

	app.Run(os.Args)
}

func findEnvironmentVariableByKey(key string) string {
	value := os.Getenv(key)

	if len(value) == 0 {
		color.Yellow("Please set 'EMAIL' on your environment, we'll use your name as a sender")
		return "info@istanbulphp.org"
	}

	return value
}
