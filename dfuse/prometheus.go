package dfuse

import (
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
	"time"
)

const SUBSYTEM = "dfuse"

var (
	promHeadBlockTime = promauto.NewGaugeVec(
		prometheus.GaugeOpts{
			Subsystem: SUBSYTEM,
			Name:      "last_block_time",
			Help:      "Last read block time from dfuse.",
		},
		[]string{"connector"},
	)

	promHeadBlockNumber = promauto.NewGaugeVec(
		prometheus.GaugeOpts{
			Subsystem: SUBSYTEM,
			Name:      "last_block_number",
			Help:      "Last read block number from dfuse.",
		},
		[]string{"connector"},
	)

	promSuccessfulHeadBlockTime = promauto.NewGaugeVec(
		prometheus.GaugeOpts{
			Subsystem: SUBSYTEM,
			Name:      "last_successful_block_time",
			Help:      "Last read block time from dfuse in which no errors occurred. If this falls behind last_block_time it means there have been errors.",
		},
		[]string{"connector"},
	)

	promSuccessfulHeadBlockNumber = promauto.NewGaugeVec(
		prometheus.GaugeOpts{
			Subsystem: SUBSYTEM,
			Name:      "last_successful_block_number",
			Help:      "Last read block number from dfuse in which no errors occurred. If this falls behind last_block_time it means there have been errors.",
		},
		[]string{"connector"},
	)

	promDfuseErrorCounter = promauto.NewCounterVec(
		prometheus.CounterOpts{
			Subsystem: SUBSYTEM,
			Name:      "handler_errors",
			Help:      "Counter for errors coming from the dfuse handlers.",
		},
		[]string{"connector"},
	)
)

func reportLastHeadBlockTime(connnector string, time time.Time) {
	promHeadBlockTime.WithLabelValues(connnector).Set(float64(time.Unix()))
}

func reportLastHeadBlockNumber(connnector string, number int) {
	promHeadBlockNumber.WithLabelValues(connnector).Set(float64(number))
}

func reportLastSuccessfulBlockTime(connnector string, time time.Time) {
	promSuccessfulHeadBlockTime.WithLabelValues(connnector).Set(float64(time.Unix()))
}

func reportLastSuccessfulBlockNumber(connnector string, number int) {
	promSuccessfulHeadBlockNumber.WithLabelValues(connnector).Set(float64(number))
}

func increaseDfuseErrorCounter(connector string) {
	promDfuseErrorCounter.WithLabelValues(connector).Inc()
}
