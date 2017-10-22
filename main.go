package main

import (
	"os"
	"runtime"

	"github.com/Sirupsen/logrus"
	"github.com/codegangsta/cli"
	"github.com/spf13/viper"
	"gopkg.in/gin-gonic/gin.v1"
)

var (
	debug bool
	nc    netCfg
)

type netCfg struct {
	user     string
	pass     string
	port     string
	listPath string
}

type httpCfg struct {
	port   string
	uiPath string
	addr   string
}

func configRuntime() {
	nuCPU := runtime.NumCPU()
	runtime.GOMAXPROCS(nuCPU)
	logrus.WithFields(logrus.Fields{
		"nCPU": nuCPU,
	}).Info("Running with number of CPUs")
}

func appCLISetup() *cli.App {
	app := cli.NewApp()
	app.Name = "Godjun - example of interaction with BGP via NETCONF in Junos"
	app.Usage = "Provides a REST API to show bgp neighbor on a Junos router with some helpfult additions. More info https://netopscasts.com/first/"
	app.Email = "egor.krv@gmail.com"
	app.Version = "0.1"
	app.Flags = []cli.Flag{
		cli.BoolFlag{
			Name:        "debug, d",
			Usage:       "debug mode currently just prety prints JSON",
			Destination: &debug,
		},
	}
	return app
}

func readCfg() *httpCfg {
	var (
		hc httpCfg
	)
	viper.SetConfigName("godjun")
	viper.AddConfigPath(".")
	viper.SetConfigType("toml")
	err := viper.ReadInConfig()
	// if a config file not found log and exit with a non zero code
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"topic": "config",
			"event": "open config failure",
		}).Error(err)
		os.Exit(1)
	}
	// NETCONF config
	nc.user = viper.GetString("netconf.user")
	nc.pass = viper.GetString("netconf.password")
	nc.port = viper.GetString("netconf.port")
	nc.listPath = viper.GetString("netconf.listpath")
	// WEB config
	hc.port = viper.GetString("http.port")
	hc.uiPath = viper.GetString("http.uipath")
	hc.addr = viper.GetString("http.address")
	return &hc
}

func main() {
	app := appCLISetup()
	hc := readCfg()
	app.Action = func(c *cli.Context) {
		configRuntime()
		gin.SetMode(gin.ReleaseMode)
		// Routes and handlers
		router := gin.Default()
		router.Static("/ui", hc.uiPath)
		router.GET("/", index)
		router.GET("/bgpnei", getBGP)
		err := router.Run(hc.addr + ":" + hc.port)
		if err != nil {
			logrus.WithFields(logrus.Fields{
				"topic": "http",
				"event": "failure to start http server",
			}).Error(err)
			os.Exit(1)
		}
	}
	app.Run(os.Args)
}
