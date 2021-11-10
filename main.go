package main

import (
	"log"
	"net/http"
	"time"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

var (
	opsProcessed = promauto.NewCounter(prometheus.CounterOpts{
		Name: "myapp_processed_ops_total",
		Help: "The total number of processed events",
	})
	gaugeMetric = promauto.NewGauge(prometheus.GaugeOpts{
		Name: "myapp_new_metric",
		Help: "Some new metric",
	})
)

func recordMetrics() {
	go func() {
		for {
			opsProcessed.Inc()
			gaugeMetric.Inc()
			time.Sleep(2 * time.Second)
		}
	}()
}

func main() {
	recordMetrics()

	const metricsPath = "/metrics"
	const portNumber = "2112"

	http.Handle(metricsPath, promhttp.Handler())
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(`
			<html>
			<body>
				<h1>Custom exporter</h1>
				<p><a href='` + metricsPath + `'>Metrics</a></p>
			</body>
			</html>
		`))
	})
	log.Printf("server listening on :%q", portNumber)
	if err := http.ListenAndServe(":"+portNumber, nil); err != nil {
		log.Fatalf("cannot start exporter: %s", err)
	}
}
