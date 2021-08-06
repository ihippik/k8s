package main

import (
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
	"github.com/urfave/cli/v2"
)

func main() {
	hostName, err := os.Hostname()
	if err != nil {
		logrus.Errorln(err)
	}

	logrus.SetLevel(logrus.DebugLevel)
	logger := logrus.WithField("host", hostName)

	app := &cli.App{
		Name:  "boom",
		Usage: "make an explosive entrance",
		Flags: []cli.Flag{
			&cli.IntFlag{
				Name:    "port",
				Aliases: []string{"p"},
				Usage:   "http server port",
				Value:   8080,
			},
		},
		Commands: []*cli.Command{
			{
				Name:    "job",
				Aliases: []string{"j"},
				Usage:   "launching a fake job",
				Flags: []cli.Flag{
					&cli.DurationFlag{
						Name:    "sleep",
						Aliases: []string{"s"},
						Usage:   "pause in second",
						Value:   5,
					},
				},
				Action: func(c *cli.Context) error {
					logger.Infoln("starting job..")

					sleep := c.Duration("sleep")

					logger.WithField("duration", sleep).Debugln("set sleep")

					<-time.NewTimer(sleep).C

					logger.Infoln("job was finished")

					return nil
				},
			},
			{
				Name:    "server",
				Aliases: []string{"s"},
				Usage:   "start http server",
				Action: func(c *cli.Context) error {
					e := echo.New()
					e.HideBanner = true
					e.HidePort = true

					e.GET("/start", func(c echo.Context) error {
						return c.String(http.StatusOK, "You`ve hit "+hostName+"\n")
					})

					e.GET("/healthcheck", func(c echo.Context) error {
						return c.String(http.StatusOK, "OK")
					})

					addr := fmt.Sprintf(":%d", c.Int("port"))

					logger.WithField("addr", addr).Infoln("server was started")

					return e.Start(addr)
				},
			},
		},
	}

	if err := app.Run(os.Args); err != nil {
		logrus.Fatal(err)
	}
}
