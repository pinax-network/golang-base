package database

import (
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
)

const SUBSYTEM = "database"

var (
	promConnsHealthy = promauto.NewGauge(prometheus.GaugeOpts{
		Subsystem: SUBSYTEM,
		Name:      "connections_healthy",
		Help:      "Number of currently healthy connections in the MySQL connection pool.",
	})
	promConnsUnhealthy = promauto.NewGauge(prometheus.GaugeOpts{
		Subsystem: SUBSYTEM,
		Name:      "connections_unhealthy",
		Help:      "Number of currently unhealthy connections in the MySQL connection pool.",
	})
	promNoHealthyConnAvailableError = promauto.NewCounter(prometheus.CounterOpts{
		Subsystem: SUBSYTEM,
		Name:      "no_healthy_connection_errors",
		Help:      "Error counter for 'no healthy connections available'",
	})
)

func recordConnStats(numHealthy, numUnhealthy int) {
	promConnsHealthy.Set(float64(numHealthy))
	promConnsUnhealthy.Set(float64(numUnhealthy))
}

func incNoHealthyConnError() {
	promNoHealthyConnAvailableError.Inc()
}
