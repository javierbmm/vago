package server

import (
	"net/http"
	"strconv"
)

func Serve(port int, directory string) {
	err := http.ListenAndServe(":"+strconv.Itoa(port), RouterHandler(http.Dir(directory)))
	if err != nil {
		panic(err)
	}
}
