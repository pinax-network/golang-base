package twilio

import (
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
)

const SUBSYTEM = "twilio"

var (
	promTwilioErrors = promauto.NewCounterVec(
		prometheus.CounterOpts{
			Subsystem: SUBSYTEM,
			Name:      "errors",
			Help:      "Total number of Twilio errors by request type (REQUEST, VERIFY).",
		},
		[]string{"type"},
	)
	promTwilioRequests = promauto.NewCounterVec(
		prometheus.CounterOpts{
			Subsystem: SUBSYTEM,
			Name:      "requests",
			Help:      "Total number of successful Twilio requests by type (REQUEST, VERIFY).",
		},
		[]string{"type"},
	)
)

func incRequestErrorCounter() {
	promTwilioErrors.WithLabelValues("request").Inc()
}

func incVerifyErrorCounter() {
	promTwilioErrors.WithLabelValues("verify").Inc()
}

func incRequestCounter() {
	promTwilioRequests.WithLabelValues("request").Inc()
}

func incVerifyCounter() {
	promTwilioRequests.WithLabelValues("verify").Inc()
}
