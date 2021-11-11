package main

import (
	"log"
	"net/http"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func main() {
	const metricsPath = "/metrics"
	const portNumber = "2112"

	exporter := NewExporter()
	prometheus.MustRegister(exporter)

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
