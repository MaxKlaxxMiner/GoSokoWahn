package main

import (
	"fmt"
	"github.com/MaxKlaxxMiner/GoSokoWahn/sokoLib/tool"
	"github.com/rs/cors"
	"net/http"
	"os"
	"path/filepath"
	"sokoServer/api"
	"strconv"
	"strings"
)

const port = "9000"

func main() {
	//test()
	//return

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path == "/" {
			w.Header().Set("Content-Type", "text/html")
			tsCss := tool.FileLastModifiedTs("../sokoClient/html/css/base.css")
			tsWasm := tool.FileLastModifiedTs("../sokoClient/html/main.wasm")
			version := tool.Max(tsCss, tsWasm)
			indexHtmlStr := string(tool.PanicErrorP1(os.ReadFile("../sokoClient/html/index.html")))
			indexHtmlStr = strings.ReplaceAll(indexHtmlStr, "{{version}}", strconv.Itoa(int(version)))
			_, _ = w.Write([]byte(indexHtmlStr))
			return
		}
		filePath := filepath.Join("../sokoClient/html", r.URL.Path)
		if _, err := os.Stat(filePath); err == nil {
			version := r.URL.Query().Get("v")
			if version != "" {
				etag := version
				if versionTs, _ := strconv.ParseInt(version, 10, 64); versionTs > 0 {
					w.Header().Set("ETag", etag)
					if match := r.Header.Get("If-None-Match"); match != "" && strings.Contains(match, etag) {
						w.WriteHeader(http.StatusNotModified)
						return
					}
				}
			}
		}

		// Wenn die Datei existiert, als statische Datei zurückgeben
		api.GzipHandler(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			http.ServeFile(w, r, filePath)
		})).ServeHTTP(w, r)
	})

	fmt.Println("run server: localhost:" + port)
	handler := cors.AllowAll().Handler(http.DefaultServeMux)
	err := http.ListenAndServe(":"+port, handler)
	if err != nil {
		fmt.Println("Failed to start server", err)
		return
	}
}
