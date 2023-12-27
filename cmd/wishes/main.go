package main

import "wishes/base"

const (
	appName = "wishes"
	host    = "0.0.0.0"
	port    = "10000"
	renderHost = "https://wishes-from-sonika.onrender.com"
)

func main() {
	// start the server
	app := base.GetApplicationInstance(appName, host, port, renderHost)
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
