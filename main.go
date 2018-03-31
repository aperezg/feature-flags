package main

import (
	"flag"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/aperezg/feature-flags/project"
	"github.com/aperezg/feature-flags/store/inmem"
	"github.com/go-kit/kit/log"
)

const (
	defaultPort = "8089"
)

func main() {
	var (
		httpAddr = ":" + envString("FF_HTTP_PORT", defaultPort)
		inmemory = flag.Bool("inmem", false, "use in-memory repositories")
	)

	flag.Parse()

	var logger log.Logger
	logger = log.NewLogfmtLogger(log.NewSyncWriter(os.Stderr))
	logger = log.With(logger, "ts", log.DefaultTimestampUTC)

	var (
		projects project.Repository
	)

	if *inmemory {
		projects = inmem.NewProjectRepository()
	}

	var ps project.Service
	ps = project.NewService(projects)

	httpLogger := log.With(logger, "component", "http")

	mux := http.NewServeMux()

	mux.Handle("/", project.MakeHandler(ps, httpLogger))
	http.Handle("/", accessControl(mux))

	errs := make(chan error, 2)
	go func() {
		logger.Log("transport", "http", "address", httpAddr, "msg", "listening")
		errs <- http.ListenAndServe(httpAddr, nil)
	}()
	go func() {
		c := make(chan os.Signal, 1)
		signal.Notify(c, syscall.SIGINT)
		errs <- fmt.Errorf("%s", <-c)
	}()

	logger.Log("terminated", <-errs)

}

func accessControl(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Origin, Content-Type")

		if r.Method == "OPTIONS" {
			return
		}

		h.ServeHTTP(w, r)
	})
}

func envString(env, fallback string) string {
	e := os.Getenv(env)
	if e == "" {
		return fallback
	}
	return e
}
