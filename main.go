package main

import (
	//"fmt"
	"github.com/urfave/cli" // renamed from codegansta
)

const (
	VERSION string = "2016-09-02"

func main() {
	app := cli.NewApp()
	app.Name = "influx-killer"
	app.Version = VERSION
	app.Authors = []cli.Author{
		cli.Author{
			Name: "Odd E. Ebbesen",
			Email: "odd.ebbesen@wirelesscar.com",
		},
	}
	app.Usage = "Stresstest InfluxDB"

	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name: "url, u",
			Usage: "Full URL (with port) to Influx endpoint",
		},
		cli.StringFlag{
			Name: "db",
			Usage: "Name of database to write to",
		},
	}
}
