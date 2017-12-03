package main

import (
	"gopkg.in/gin-gonic/gin.v1"
)

func index(c *gin.Context) {
	c.Redirect(301, "/dist/index.html")
}

func getBGP(c *gin.Context) {
	hosts, err := readHosts(nc.listPath)
	if err != nil {
		c.JSON(500, err.Error())
		return
	}
	if len(hosts) > 0 {
		bgpData := pollJunipers(hosts)
		c.JSON(200, bgpData)
	} else {
		c.JSON(200, nil)
	}
}
