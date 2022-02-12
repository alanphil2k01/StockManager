package server

import (
	"context"
	"embed"
	"io/fs"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/alanphil2k01/SSMC/pkg/config"
	"github.com/alanphil2k01/SSMC/pkg/db"
	"github.com/alanphil2k01/SSMC/pkg/routes"
	"github.com/gorilla/mux"
)

var (
	//go:embed static
	content embed.FS
	address string
	port    string
	srv     *http.Server
)

func init() {
	address = config.GetAddress()
	srv = createServer()
}

func createServer(mws ...mux.MiddlewareFunc) *http.Server {
	router := mux.NewRouter()
	fsys, _ := fs.Sub(content, "static")
	for _, mw := range mws {
		router.Use(mw)
	}
	routes.RegisterRoutes(router)
	router.PathPrefix("/").Handler(http.StripPrefix("/", http.FileServer(http.FS(fsys))))

	return &http.Server{
		Handler:           router,
		Addr:              address,
		ReadTimeout:       1 * time.Second,
		ReadHeaderTimeout: 1 * time.Second,
		WriteTimeout:      1 * time.Second,
		IdleTimeout:       1 * time.Second,
	}
}

func RunServer() (<-chan error, error) {
	errC := make(chan error, 1)

	ctx, stop := signal.NotifyContext(context.Background(),
		os.Interrupt,
		syscall.SIGTERM,
		syscall.SIGQUIT,
	)

	go func() {
		<-ctx.Done()

		log.Println("Shutdown signal received")

		ctxTimeout, cancel := context.WithTimeout(context.Background(), 5*time.Second)

		defer func() {
			db.Close()
			stop()
			cancel()
			close(errC)
		}()

		srv.SetKeepAlivesEnabled(false)

		if err := srv.Shutdown(ctxTimeout); err != nil {
			errC <- err
		}

		log.Println("Shutdown complete")
		os.Exit(0)
	}()

	go func() {
		log.Println("Listening and serving on ", address)
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			errC <- err
		}
	}()

	go func() {
		for {
			time.Sleep(24 * time.Hour)
			db.RemoveExpired()
		}
	}()

	return errC, nil
}
