package main

import (
	//"fmt"
	"github.com/urfave/cli" // renamed from codegansta
)

const (
	VERSION string = "2016-09-02"
)

func main() {
	app := cli.NewApp()
	app.Name = "influx-killer"
	app.Version = VERSION
	app.Authors = []cli.Author{
		cli.Author{
			Name:  "Odd E. Ebbesen",
			Email: "odd.ebbesen@wirelesscar.com",
		},
	}
	app.Usage = "Stresstest InfluxDB"

	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:  "url, u",
			Usage: "Full URL (with port) to Influx endpoint",
		},
		cli.StringFlag{
			Name:  "db",
			Usage: "Name of database to write to",
		},
		cli.StringFlag{
			Name:  "hostname-prefix, h",
			Usage: "Prefix for generated hostnames",
		},
		cli.UintFlag{
			Name:  "num-hosts, n",
			Usage: "Number of hosts to simulate traffic from",
		},
		cli.Float64Flag{
			Name:  "interval, i",
			Usage: "How long (in seconds, fractions allowed) between sending metrics",
		},
		cli.Float64Flag{
			Name:  "timeout, t",
			Usage: "How long in seconds (fractions allowed) to run the test",
		},
		cli.StringFlag{
			Name:  "measurement, m",
			Usage: "The Influx unit to report on, e.g. 'cpu'",
		},
		cli.StringFlag{
			Name:  "field, f",
			Usage: "The field to set the generated value on, e.g. 'load'",
		},
	}
}
