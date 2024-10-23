package main

import (
	"flag"
	"fmt"
	"log/slog"
	"net/http"
)

func main() {

	addr := flag.String("addr", ":4000", "listening port no")
	flag.Parse()
	mux := http.NewServeMux()
	mux.HandleFunc("/", home)
	srv := &http.Server{
		Addr:    *addr,
		Handler: mux,
	}
	slog.Info("server listening on port ", "port", *addr)
	err := srv.ListenAndServe()
	if err != nil {
		slog.Error("error listening to the server", "port", err)
	}
}

func home(w http.ResponseWriter, r *http.Request) {
	slog.Info("innn")
	if r.Method != http.MethodGet {
		fmt.Fprintf(w, "use only get")
		slog.Error("invalid method request")
		return
	}
	w.Write([]byte("hello world"))
}
