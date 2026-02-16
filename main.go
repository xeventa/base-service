package main

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"github.com/xeventa/base-service/app"
	"github.com/xeventa/base-service/core/environment"
)

func main() {
	// initialize config
	conf, err := environment.ProvideConfig()
	if err != nil {
		panic(err.Error())
	}

	logrus.Info(conf.AppName, " is running on ", conf.AppHost, ":", conf.AppPort, " ", conf.AppEnv, " mode")

	// Force log's color
	gin.ForceConsoleColor()
	router := gin.New()
	router.Use(gin.Logger())

	// Recovery middleware recovers from any panics and writes a 500 if there was one.
	router.Use(gin.Recovery())

	routes, err := app.InjectRoutes()
	if err != nil {
		panic(err.Error())
	}

	for _, route := range routes {
		route.Register(router)
	}

	// Start server using Listen so we honor configured host/port/protocol
	Listen(conf, router)
}

// Listen is a func to start http server
func Listen(conf *environment.Config, handler http.Handler) {
	addr := conf.AppHost + ":" + strconv.Itoa(conf.AppPort)
	var err error
	if conf.AppProtocol == "http" {
		err = http.ListenAndServe(addr, handler)
	} else {
		err = errors.New("Unsupported protocol: " + conf.AppProtocol)
	}
	if err != nil {
		panic("Failed to start server: " + err.Error())
	}
}
