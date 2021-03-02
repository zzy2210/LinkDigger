package main

import (
	"LinkDigger/Global"
	"LinkDigger/Scan"
	"github.com/urfave/cli"
	"log"
	"os"
)

func main()  {
	var url string
	Global.Finder.Flags = append(Global.Finder.Flags,
		cli.StringFlag{
			Name: "url,u",
			Usage: "Target webserver's url",
			Destination: &url,
		},
		cli.BoolFlag{
			Name: "path,p",
			Usage: "to find webserver's path",
		},
		cli.IntFlag{
			Name: "n",
			Usage: "The upper limit of goroutine",
			Value: 5,
			Destination: &Global.Limit,
		},
		cli.BoolFlag{
			Name: "deep,d",
			Usage: "deep scan",
		})

	Global.Finder.Action = func(c *cli.Context){
		Scan.Control(url,c.Bool("deep"))
	}

	err := Global.Finder.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}


}

