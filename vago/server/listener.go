package server

import (
	"net/http"
	"strconv"
	generator "vago/vago/log/server"
)

func Serve(port int, directory string) {
	var logger generator.ServerLogger
	logger.Init()

	err := http.ListenAndServe(":"+strconv.Itoa(port), RouterHandler(http.Dir(directory)))
	if err != nil {
		logger.Error(err)
	}
}
