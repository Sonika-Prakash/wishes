package main

import "wishes/base"

const (
	appName = "wishes"
	host    = "localhost"
	port    = "10000"
)

func main() {
	// start the server
	app := base.GetApplicationInstance(appName, host, port)
	router := base.GetRouter(app)
	srv := app.GetServer(router)

	errs := make(chan error)
	go func() {
		errs <- srv.ListenAndServe()
	}()
	go func() {
		app.CatchInterruptions(errs)
	}()
	err := <-errs
	app.GracefulShutdown(srv, err)
}
