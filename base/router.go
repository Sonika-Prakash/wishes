package base

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"runtime/debug"
	"syscall"
	"wishes/public"

	"github.com/CloudyKit/jet/v6"
	"github.com/gorilla/mux"
)

// Server is the server struct
type Server struct {
	host string
	port string
	renderHost string
}

// Application is the application struct
type Application struct {
	appName string
	server  Server
	debug   bool
	infoLog log.Logger
	errLog  log.Logger
	view    *jet.Set
}

// GetApplicationInstance returns the pointer to Application instance
func GetApplicationInstance(appName, host, port, renderHost string) *Application {
	jet := initJet()
	return &Application{
		appName: appName,
		server: Server{
			host: host,
			port: port,
			renderHost: renderHost,
		},
		debug:   true,
		infoLog: *log.New(os.Stdout, "INFO\t", log.Ltime|log.Ldate|log.Lshortfile),
		errLog:  *log.New(os.Stderr, "ERROR\t", log.Ltime|log.Ldate|log.Llongfile),
		view:    jet,
	}
}

// GetRouter returns the mux router
func GetRouter(app *Application) http.Handler {
	router := mux.NewRouter()
	if app.debug {
		router.Use(loggingMiddleware)
	}

	router.HandleFunc("/health", app.HealthHandler).Methods(http.MethodGet)
	router.HandleFunc("/", app.IndexHandler).Methods(http.MethodGet)
	router.HandleFunc("/", app.IndexPostHandler).Methods(http.MethodPost)

	// exposing css and images via /public path which is referenced by html pages
	fileServer := http.FileServer(http.FS(public.Files))
	router.PathPrefix("/public/").Handler(http.StripPrefix("/public", fileServer))

	return router
}

// GetServer returns the HTTP server
func (a *Application) GetServer(h http.Handler) *http.Server {
	url := fmt.Sprintf("%s:%s", a.server.host, a.server.port)
	srv := http.Server{
		Handler: h,
		Addr:    url,
	}
	a.infoLog.Printf("Starting HTTP server on %s", url)
	return &srv
}

// CatchInterruptions catches any interruptions
func (a *Application) CatchInterruptions(errs chan error) {
	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)
	sig := <-c
	a.errLog.Printf("caught %s, so application will attempt to gracefully shutdown", sig.String())
	errs <- fmt.Errorf("%s", sig)
}

// GracefulShutdown shuts down the server gracefully
func (a *Application) GracefulShutdown(srv *http.Server, e error) {
	a.infoLog.Println("Received an error on channel", e.Error())
	// exit gracefully
	if errHTTPServer := srv.Shutdown(context.Background()); errHTTPServer != nil {
		a.errLog.Println("failed to gracefully shutdown HTTP server", errHTTPServer.Error())
	}
	a.infoLog.Println("server shutdown complete, application will now exit")
}

func (a *Application) serverErr(w http.ResponseWriter, err error) {
	trace := fmt.Sprintf("%s\n%s", err.Error(), debug.Stack())
	a.errLog.Output(2, trace)
	http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
}
