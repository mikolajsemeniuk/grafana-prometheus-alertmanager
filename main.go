package main

import (
	"net/http"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

var hits *prometheus.CounterVec

func init() {
	options := prometheus.CounterOpts{Name: "api_hits", Help: "Number of hits per endpoint"}
	labels := []string{"endpoint"}
	hits = prometheus.NewCounterVec(options, labels)
	prometheus.MustRegister(hits)
}

func main() {
	handler := func(w http.ResponseWriter, r *http.Request) {
		hits.WithLabelValues("/hits").Inc()
		w.Write([]byte("This is the API endpoint."))
	}
	http.HandleFunc("/hits", handler)
	http.Handle("/metrics", promhttp.Handler())
	http.ListenAndServe(":8080", nil)
}
