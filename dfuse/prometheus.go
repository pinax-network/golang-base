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
			Name:      "head_block_time",
			Help:      "Last successfully read block time from dfuse.",
		},
		[]string{"connector"},
	)

	promHeadBlockNumber = promauto.NewGaugeVec(
		prometheus.GaugeOpts{
			Subsystem: SUBSYTEM,
			Name:      "head_block_number",
			Help:      "Last successfully read block number from dfuse.",
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
