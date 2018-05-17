package main

import (
	//"flag"
	"fmt"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

var (
	delay           float64 //delay seconds before health turned abnormal.
	listen          string  // listen address
	count           float64
	unhealthElapsed = prometheus.NewCounter(
		prometheus.CounterOpts{
			Name: "unhealth_elapsed",
			Help: "Seconds of abnormal state elapsed",
		},
	)
)

func init() {
	prometheus.MustRegister(unhealthElapsed)
}

func main() {
	// find vars from env
	delay, err := strconv.ParseFloat(os.Getenv("DELAY_SECONDS"), 64)
	if err != nil {
		delay = 60
	}

	listen = os.Getenv("LISTEN")
	if listen == "" {
		listen = "0.0.0.0:8899"
	}

	//flag.Float64Var(&delay, "delay", 60, "wait seconds before health turned abnormal")

	//flag.StringVar(&listen, "listen", "0.0.0.0:8899", "listening address")

	//// parse flags
	//flag.Parse()

	handler := http.NewServeMux()
	handler.Handle("/health", health())
	handler.Handle("/metrics", metrics())

	quit := make(chan struct{})

	go counter(quit)

	srv := &http.Server{
		Addr:    listen,
		Handler: handler,
	}

	fmt.Println("listening on", listen, "for health check and metrics")
	if err := srv.ListenAndServe(); err != nil {
		fmt.Println("server quit")
	}

	close(quit)
}

func health() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if count < delay {
			w.WriteHeader(http.StatusOK)
			w.Write([]byte("ok"))
			return
		}

		w.WriteHeader(http.StatusInternalServerError)
	})
}

func metrics() http.Handler {
	return promhttp.Handler()
}

func counter(quit <-chan struct{}) {
	ticker := time.NewTicker(time.Second)
	defer ticker.Stop()

	for {
		select {
		case <-ticker.C:
			count++
			if count > delay {
				unhealthElapsed.Inc()
			}
		case <-quit:
			return
		}
	}
}
