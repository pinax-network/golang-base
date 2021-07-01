package log

import (
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
)

const SUBSYTEM = "log"

var (
	promLogsFatal = promauto.NewCounter(prometheus.CounterOpts{
		Subsystem: SUBSYTEM,
		Name:      "fatal_logs",
		Help:      "Counter for logs of FATAL level.",
	})
	promLogsPanic = promauto.NewCounter(prometheus.CounterOpts{
		Subsystem: SUBSYTEM,
		Name:      "panic_logs",
		Help:      "Counter for logs of PANIC level.",
	})
	promLogsError = promauto.NewCounter(prometheus.CounterOpts{
		Subsystem: SUBSYTEM,
		Name:      "error_logs",
		Help:      "Counter for logs of ERROR level.",
	})
	promLogsWarn = promauto.NewCounter(prometheus.CounterOpts{
		Subsystem: SUBSYTEM,
		Name:      "warn_logs",
		Help:      "Counter for logs of WARNING level.",
	})
)

func incFatalCounter() {
	promLogsFatal.Inc()
}

func incPanicCounter() {
	promLogsPanic.Inc()
}

func incErrorCounter() {
	promLogsError.Inc()
}

func incWarnCounter() {
	promLogsWarn.Inc()
}
