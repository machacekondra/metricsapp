package main

import (
	"math/rand"
	"net/http"
	"time"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func recordMetrics() {
	go func() {
		for {
			metricsCPU.WithLabelValues("Cloud Pak for Data", "11827733", "IBM", "5cb3d853-a575-4346-ba85-db55c64b5dfd").Set((float64)(rand.Intn(20)))
			time.Sleep(2 * time.Second)
		}
	}()
}

var (
	metricsCPU = promauto.NewGaugeVec(
		prometheus.GaugeOpts{
			Name: "restricted_use_cores:sum",
			Help: "Restricted CPU usage",
		},
		[]string{
			"product",
			"product_identifier",
			"vendor",
			"vendor_identifier",
		},
	)
)

func main() {
	recordMetrics()

	http.Handle("/metrics", promhttp.Handler())
	http.ListenAndServe(":8081", nil)
}
