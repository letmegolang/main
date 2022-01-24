package main

import (
	"fmt"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"net/http"
)

/* need to install
go get github.com/prometheus/client_golang/prometheus
go get github.com/prometheus/client_golang/prometheus/promauto
go get github.com/prometheus/client_golang/prometheus/promhttp
*/

func main() {
	// usual page
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Just index page")

		PageProcessed.Inc() // increase custom application-specific metric (index_page_counter)
		CustomGauge.Set(10)
	})

	// metrics URL
	http.Handle("/metrics", promhttp.Handler())

	http.ListenAndServe(":80", nil)
}

// define application-specific metric with name index_page_counter
// index_page_counter metrics type - Counter
// some_gauge_var - Gauge
var (
	PageProcessed = promauto.NewCounter(prometheus.CounterOpts{
		Name: "index_page_counter",
		Help: "The total number of open base URL page open events",
	})

	CustomGauge = promauto.NewGauge(prometheus.GaugeOpts{
		Name: "some_gauge_var",
		Help: "Current gauge data, f.e. concurrent active users",
	})
)
