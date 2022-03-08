package main

import (
	"flag"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/go-kit/log"

	"behometest/infrastructure/persistence"
	"behometest/interfaces/wallet"
)

func main() {
	var (
		addr       = envString("PORT", "8080")
		dbhost     = envString("DB_HOST", "localhost")
		dbport     = envString("DB_PORT", "5432")
		dbuser     = envString("DB_USER", "coins")
		dbname     = envString("DB_NAME", "coins")
		dbpassword = envString("DB_PASSWORD", "mysecretpassword")

		httpAddr = flag.String("http.addr", ":"+addr, "HTTP listen address")
	)

	flag.Parse()

	var logger log.Logger
	logger = log.NewLogfmtLogger(log.NewSyncWriter(os.Stderr))
	logger = log.With(logger, "ts", log.DefaultTimestampUTC)

	r, err := persistence.NewRepositories(dbhost, dbport, dbuser, dbname, dbpassword)
	if err != nil {
		panic(err)
	}
	r.Automigrate()

	walletService := wallet.NewService(r.AccountRepository, r.PaymentRepository)
	httpLogger := log.With(logger, "component", "http")

	mux := http.NewServeMux()

	mux.Handle("/v1/", wallet.MakeHandler(walletService, httpLogger))
	http.Handle("/", accessControl(mux))

	errs := make(chan error, 2)
	go func() {
		logger.Log("transport", "http", "address", *httpAddr, "msg", "listening")
		errs <- http.ListenAndServe(*httpAddr, nil)
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
