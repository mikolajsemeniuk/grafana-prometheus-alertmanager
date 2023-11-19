package main

import (
	"fmt"
	"time"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/push"
)

var hits *prometheus.CounterVec
var pushgateway = "http://localhost:9091"
var pusher *push.Pusher

func init() {
	options := prometheus.CounterOpts{Name: "api_hits", Help: "Number of hits per endpoint"}
	labels := []string{"type"}
	hits = prometheus.NewCounterVec(options, labels)
	prometheus.MustRegister(hits)
	pusher = push.New(pushgateway, "worker").Collector(hits)
}

func main() {
	for i := 0; i < 10; i++ {
		hits.WithLabelValues("worker").Inc()

		if err := pusher.Push(); err != nil {
			fmt.Println(err)
		}

		time.Sleep(time.Millisecond * 300)
	}

	fmt.Println("done")
}
