package audit

import (
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
)

const SUBSYTEM = "audit"

var (
	promCreateResource = promauto.NewCounter(prometheus.CounterOpts{
		Subsystem: SUBSYTEM,
		Name:      "created_resources",
		Help:      "Counter for created resources.",
	})
	promUpdateResource = promauto.NewCounter(prometheus.CounterOpts{
		Subsystem: SUBSYTEM,
		Name:      "updated_resources",
		Help:      "Counter for updated resources.",
	})
	promDeleteResource = promauto.NewCounter(prometheus.CounterOpts{
		Subsystem: SUBSYTEM,
		Name:      "deleted_resources",
		Help:      "Counter for deleted resources.",
	})
)

func incCreateResourceCounter() {
	promCreateResource.Inc()
}

func incUpdateResourceCounter() {
	promUpdateResource.Inc()
}

func incDeleteResourceCounter() {
	promDeleteResource.Inc()
}
