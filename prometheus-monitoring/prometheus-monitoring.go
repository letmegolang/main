package main

import (
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"net/http"
)

func main() {
	// metrics URL
	http.Handle("/metrics", promhttp.Handler())
	http.ListenAndServe(":80", nil)
}
