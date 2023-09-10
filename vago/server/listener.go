package server

import (
	"net/http"
	"strconv"
	"vago/vago/input"
	"vago/vago/log/server"
)

func Serve(port int, config input.IOPath) {
	var logger server.ServerLogger
	directory := config.OutFolder
	homepath := config.OutFolder + config.Home
	logger.Init()

	logger.Info("Starting...")
	logger.Info("Listening for requests on port %s", strconv.Itoa(port))
	err := http.ListenAndServe(":"+strconv.Itoa(port), RouterHandler(http.Dir(directory), homepath))
	if err != nil {
		logger.Error(err)
	}
	logger.Warning("Closing server and port...")
}
