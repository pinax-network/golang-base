package log

import (
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
)

const SUBSYTEM = "log"

var (
	promLogs = promauto.NewCounterVec(
		prometheus.CounterOpts{
			Subsystem: SUBSYTEM,
			Name:      "count",
			Help:      "Total number of log entries by log level.",
		},
		[]string{"level"},
	)
)

func incFatalCounter() {
	promLogs.WithLabelValues("fatal").Inc()
}

func incPanicCounter() {
	promLogs.WithLabelValues("panic").Inc()
}

func incErrorCounter() {
	promLogs.WithLabelValues("error").Inc()
}

func incWarnCounter() {
	promLogs.WithLabelValues("warn").Inc()
}
