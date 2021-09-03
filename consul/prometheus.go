package consul

import (
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
)

const SUBSYTEM = "consul"

var (
	promConsulErrors = promauto.NewCounterVec(
		prometheus.CounterOpts{
			Subsystem: SUBSYTEM,
			Name:      "errors",
			Help:      "Total number of Consul errors by request type (PUT, GET).",
		},
		[]string{"type"},
	)
	promConsulRequests = promauto.NewCounterVec(
		prometheus.CounterOpts{
			Subsystem: SUBSYTEM,
			Name:      "requests",
			Help:      "Total number of successful Consul requests by type (PUT, GET).",
		},
		[]string{"type"},
	)
)

func incPutErrorCounter() {
	promConsulErrors.WithLabelValues("put").Inc()
}

func incGetErrorCounter() {
	promConsulErrors.WithLabelValues("get").Inc()
}

func incPutCounter() {
	promConsulRequests.WithLabelValues("put").Inc()
}

func incGetCounter() {
	promConsulRequests.WithLabelValues("get").Inc()
}
