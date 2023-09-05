package server

import (
	"net/http"
	"strconv"
	"vago/vago/log/server"
)

func Serve(port int, directory string) {
	var logger server.ServerLogger
	logger.Init()

	logger.Info("Starting...")
	logger.Info("Listening for requests on port %s", strconv.Itoa(port))
	err := http.ListenAndServe(":"+strconv.Itoa(port), RouterHandler(http.Dir(directory)))
	if err != nil {
		logger.Error(err)
	}
	logger.Warning("Closing server and port...")
}
