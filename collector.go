package main

import (
	"github.com/prometheus/client_golang/prometheus"
)

const namespace = "my_app"

type Exporter struct {
	ingressApplication *prometheus.Desc
	ingressTechnical   *prometheus.Desc
	ingressThirdParty  *prometheus.Desc
}

func NewExporter() *Exporter {
	return &Exporter{
		ingressApplication: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, "", "ingress_application"),
			"Some new metric",
			[]string{"hostname"},
			nil,
		),
		ingressTechnical: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, "", "ingress_technical"),
			"Some new metric",
			[]string{"hostname"},
			nil,
		),
		ingressThirdParty: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, "", "ingress_third_party"),
			"Some new metric",
			[]string{"hostname"},
			nil,
		),
	}
}

func (e *Exporter) Describe(ch chan<- *prometheus.Desc) {
	ch <- e.ingressApplication
	ch <- e.ingressTechnical
	ch <- e.ingressThirdParty
}

func (e *Exporter) Collect(ch chan<- prometheus.Metric) {
	var urls = []string{"google1.com", "google2.com", "google3.com", "google4.com", "google5.com", "google6.com"}
	for _, element := range urls[0:2node] {
		ch <- prometheus.MustNewConstMetric(
			e.ingressApplication,
			prometheus.GaugeValue,
			1,
			element,
		)
	}
	for _, element := range urls[2:4] {
		ch <- prometheus.MustNewConstMetric(
			e.ingressTechnical,
			prometheus.GaugeValue,
			1,
			element,
		)
	}
	for _, element := range urls[4:6] {
		ch <- prometheus.MustNewConstMetric(
			e.ingressThirdParty,
			prometheus.GaugeValue,
			1,
			element,
		)
	}
}
