package main

import (
	"os"
	"net/url"
	"regexp"
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

				_, err := url.ParseRequestURI(c.Args().First())

				if err != nil {
					color.Red("Error: Invalid URL. Be sure you specified fully qualified URL like " +
						"https://google.com.tr")

					return nil
				}

				email := findEnvironmentVariableByKey("EMAIL")

				if !validateEmail(email) {
					color.Red("Error: Invalid Email address.")

					return nil
				}

				resp, err := resty.R().
					SetBody(Link{Url: c.Args().First(), Email: email}).
					Post(ENDPOINT)

				if err != nil {
					color.Red("Error: Something went wrong while sharing URL.")
				}

				if resp.StatusCode() == 422 {
					color.Red("Error: Please check EMAIL(from your env.) and URL you've provided.")
				} else if resp.StatusCode() == 500 {
					color.Red("Error: API fucked up! Please getting touch with: info@istanbulphp.org")
				}

				color.Green("Thank you!")

				return nil
			},
		},
	}

	app.Run(os.Args)
}

func findEnvironmentVariableByKey(key string) string {
	value := os.Getenv(key)

	if len(value) == 0 {
		color.Yellow("Next time If you set EMAIL= to your environment, we'll send a gift to you!")

		return "info@istanbulphp.org"
	}

	return value
}

func validateEmail(email string) bool {
	Re := regexp.MustCompile(`^[a-z0-9._%+\-]+@[a-z0-9.\-]+\.[a-z]{2,4}$`)

	return Re.MatchString(email)
}