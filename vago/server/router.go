package server

import (
	"fmt"
	"net"
	"net/http"
	"os"
	"path"
	"strings"
	"vago/vago/log/server"
)

const NotFoundPage = "./error/404.html"

func notFound(w http.ResponseWriter, r *http.Request) {
	// Here you can send your custom 404 back.
	http.ServeFile(w, r, NotFoundPage)
}

func home(w http.ResponseWriter, r *http.Request, homepath string) {
	http.ServeFile(w, r, homepath)
}

func RouterHandler(fs http.FileSystem, homepath string) http.Handler {
	var logger server.ServerLogger
	logger.Init()

	fileServer := http.FileServer(fs)
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		pagename := path.Clean(r.URL.Path)
		requesterIp, err := getIP(r)
		if err != nil {
			logger.Warning(err.Error())
		} else {
			logger.Log(pagename, requesterIp)
		}

		if pagename == "/" {
			// If empty, output home.
			home(w, r, homepath)
			logger.Info("Sending home page.")
			return
		}

		_, err = fs.Open(pagename) // Do not allow path traversals.
		if os.IsNotExist(err) {
			notFound(w, r)
			logger.Warning("Page %s not found. Sending 404.", pagename)
			return
		}
		fileServer.ServeHTTP(w, r)
	})
}

func getIP(r *http.Request) (string, error) {
	//Get IP from the X-REAL-IP header
	ip := r.Header.Get("X-REAL-IP")
	netIP := net.ParseIP(ip)
	if netIP != nil {
		return ip, nil
	}

	//Get IP from X-FORWARDED-FOR header
	ips := r.Header.Get("X-FORWARDED-FOR")
	splitIps := strings.Split(ips, ",")
	for _, ip := range splitIps {
		netIP := net.ParseIP(ip)
		if netIP != nil {
			return ip, nil
		}
	}

	//Get IP from RemoteAddr
	ip, _, err := net.SplitHostPort(r.RemoteAddr)
	if err != nil {
		return "", err
	}
	netIP = net.ParseIP(ip)
	if netIP != nil {
		return ip, nil
	}
	return "", fmt.Errorf("No valid IP found")
}
