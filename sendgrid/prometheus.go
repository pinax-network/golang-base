package sendgrid

import (
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
)

const SUBSYTEM = "sendgrid"

var (
	promSendgridErrors = promauto.NewCounterVec(
		prometheus.CounterOpts{
			Subsystem: SUBSYTEM,
			Name:      "errors",
			Help:      "Total number of Sendgrid errors by request endpoint.",
		},
		[]string{"endpoint"},
	)
	promSendgridRequests = promauto.NewGaugeVec(
		prometheus.GaugeOpts{
			Subsystem: SUBSYTEM,
			Name:      "request_time",
			Help:      "Timestamp of the last successful requests by endpoint.",
		},
		[]string{"endpoint"},
	)
)

func addFailedRequest(endpoint string) {
	promSendgridErrors.WithLabelValues(endpoint).Inc()
}

func addSuccessfulRequest(endpoint string) {
	promSendgridRequests.WithLabelValues(endpoint).SetToCurrentTime()
}
