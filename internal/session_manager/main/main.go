package main

import (
	sessionmanager "diwashnembang/snippetbox/internal/session_manager"
	"flag"
	"fmt"
	"log/slog"
	"net/http"
)

type application struct {
	SessionMangaer *sessionmanager.SessionManager
}

func main() {

	addr := flag.String("addr", ":4000", "listening port no")
	flag.Parse()
	sm := sessionmanager.NewSessionManager()
	app := &application{
		SessionMangaer: sm,
	}
	mux := http.NewServeMux()
	mux.Handle("/", app.SessionMangaer.AddCookieMiddleWare(http.HandlerFunc(app.home)))
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

func (a *application) home(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		fmt.Fprintf(w, "use only get")
		slog.Error("invalid method request")
		return
	}
	w.Write([]byte("hello world"))
	token, err := r.Cookie("sessionId")
	if err != nil {
		fmt.Print("error", err)
	}
	a.SessionMangaer.Store.AddSessionValue(token.Value, "authorized", true)
	value, _ := a.SessionMangaer.Store.GetSessionValue(token.Value, "authorized")
	fmt.Print(value)

}
