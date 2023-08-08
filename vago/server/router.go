package server

import (
	"net/http"
	"os"
	"path"
)

const NotFoundPage = "./error/404.html"

func notFound(w http.ResponseWriter, r *http.Request) {
	// Here you can send your custom 404 back.
	http.ServeFile(w, r, NotFoundPage)
}

func RouterHandler(fs http.FileSystem) http.Handler {
	fileServer := http.FileServer(fs)
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		_, err := fs.Open(path.Clean(r.URL.Path)) // Do not allow path traversals.
		if os.IsNotExist(err) {
			notFound(w, r)
			return
		}
		fileServer.ServeHTTP(w, r)
	})
}

//func findPage(name string, folder string) os.DirEntry {
//	files, err := os.ReadDir(folder)
//	if err != nil {
//		panic(err)
//		return nil
//	}
//	for _, file := range files {
//		if file.Name() == name {
//			return file
//		}
//	}
//
//	return nil
//}
