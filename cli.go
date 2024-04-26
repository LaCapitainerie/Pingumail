package main

import (
	"client"
	"os"
	"strings"

	"github.com/urfave/cli/v2" // imports as package "cli"
)

func CliStart() {

	app := cli.NewApp()
	app.Name = "pingumail"
	app.Usage = "A simple mail server"
	app.Commands = []*cli.Command{
		{
			Name:    "reload",
			Aliases: []string{"r"},
			Usage:   "Reload the unread mails",
			Action: func(c *cli.Context) error {
				println("Reloading unread mails...")
				mails := client.Reload()

				var username = os.Getenv("pinguUserName")

				for _, mail := range mails {
					if mail.To == username {
						println("From", mail.From, ":", mail.Body)
					}
				}
				return nil
			},
		},
		{
			Name:    "send",
			Aliases: []string{"s", "mail", "m"},
			Flags: []cli.Flag{
				&cli.StringFlag{
					Name:    "t",
					Aliases: []string{"to"},
					Usage:   "Receiver of the mail",
				},
				&cli.StringFlag{
					Name:    "b",
					Aliases: []string{"body"},
					Usage:   "Body of the mail",
				},
			},
			Usage: "Send a mail",
			Action: func(c *cli.Context) error {

				println("Sending mail...")

				to := c.String("to")
				body := c.String("body")

				client.SendMail(to, body)

				return nil
			},
		},
		{
			Name:    "login",
			Usage:   "Login as a user",
			Aliases: []string{"l"},
			Action: func(c *cli.Context) error {
				println("Logging in...")

				if c.NArg() != 1 {
					println("Usage: pingumail login <username>")
					return nil
				}

				token := client.Login(c.Args().Get(0))
				println("Token:", token)
				if token != "" {
					envVar := []byte(strings.Join([]string{
						"pinguServerIP=" + os.Getenv("pinguServerIP"),
						"pinguToken=" + token,
					}, "\n"))

					os.WriteFile("client.env", envVar, 0644)
				} else {
					println("Login failed")
				}

				return nil
			},
		},
	}

	app.RunAndExitOnError()
}
