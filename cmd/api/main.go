package main

import (
	"net/http"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

var hits *prometheus.CounterVec

func init() {
	options := prometheus.CounterOpts{Name: "api_hits", Help: "Number of hits per endpoint"}
	labels := []string{"type"}
	hits = prometheus.NewCounterVec(options, labels)
	prometheus.MustRegister(hits)
}

func main() {
	increment := func(w http.ResponseWriter, r *http.Request) {
		hits.WithLabelValues("api").Inc()
		w.Write([]byte("This is the API endpoint."))
	}
	http.HandleFunc("/inc", increment)

	reset := func(w http.ResponseWriter, r *http.Request) {
		hits.Reset()
		w.Write([]byte("metric reset"))
	}
	http.HandleFunc("/reset", reset)

	http.Handle("/metrics", promhttp.Handler())
	http.ListenAndServe(":8080", nil)
}
