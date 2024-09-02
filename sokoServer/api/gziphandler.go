package api

import (
	"compress/gzip"
	"github.com/MaxKlaxxMiner/GoSokoWahn/sokoLib/tool"
	"log"
	"net/http"
	"path/filepath"
	"strings"
)

func GzipHandler(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Überprüfe den Dateityp
		ext := strings.ToLower(filepath.Ext(r.URL.Path))
		switch ext {
		case ".woff2":
			next.ServeHTTP(w, r)
			return
		case ".html", ".map", ".wasm", ".css", ".js", ".ico", "":
		default:
			log.Println("unknown fileext:", ext)
		}

		// Setze den Header, um Gzip zu signalisieren
		w.Header().Set("Content-Encoding", "gzip")
		gz := tool.PanicErrorP1(gzip.NewWriterLevel(w, 1))
		defer gz.Close()

		gzr := gzipResponseWriter{Writer: gz, ResponseWriter: w}
		next.ServeHTTP(gzr, r)
	})
}

type gzipResponseWriter struct {
	http.ResponseWriter
	Writer *gzip.Writer
}

func (w gzipResponseWriter) Write(b []byte) (int, error) {
	return w.Writer.Write(b)
}
