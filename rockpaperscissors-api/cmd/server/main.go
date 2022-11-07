package main

import (
	"flag"
	"rockpaperscissors-api/internal/app"

	"github.com/sirupsen/logrus"
)

func main() {

	debug := flag.Bool("debug", false, "Sets echo as debug mode")

	flag.Parse()

	app := app.NewWithConfig(app.AppConfig{
		Debug: *debug,
	})

	// go app.InitializeHub()

	logrus.Fatal(app.Start("0.0.0.0:4040"))
}
