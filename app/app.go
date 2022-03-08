package app

import "github.com/urfave/cli"

func CreateApp() *cli.App {
	app := cli.NewApp()

	app.Name = "Command Line App"
	app.Usage = "Find IP's and Servers Names"

	flags := []cli.Flag{
		cli.StringFlag{
			Name: "host",
			Value: "google.com",
		},
	}

	app.Commands = []cli.Command{
		{
			Name: "ip",
			Usage: "Find IP's",
			Flags: flags,
			Action: findIp,
		},
		{
			Name: "server",
			Usage: "Find Server Names",
			Flags: flags,
			Action: findServers,
		},
	}

	return app
}

