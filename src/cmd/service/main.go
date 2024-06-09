package main

import (
	"fmt"
	"net"

	"github.com/alifnaufalyasin/boilerplate-be-golang/src/cmd/service/config"
	"github.com/alifnaufalyasin/boilerplate-be-golang/src/cmd/service/http"
	"github.com/alifnaufalyasin/boilerplate-be-golang/src/internal/utils"
	"github.com/labstack/echo/v4"
)

func main() {
	logger := utils.NewLogger("debug")
	cfg, err := config.GetConfig()
	if err != nil {
		logger.Fatal(err, "GetConfig got error")
	}

	echoMainServer := echo.New()
	apiServer := http.Init(echoMainServer, cfg.Secret, logger)

	apiServer.User()

	// Server Listener
	logger.Fatal(apiServer.EchoServer.Start(fmt.Sprintf(":%s", cfg.Port)), "error to start http server")
	logger.Infof("Port is: %d", apiServer.EchoServer.Listener.Addr().(*net.TCPAddr).Port)
}
