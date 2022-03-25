package dfuse

import (
	"github.com/eosnationftw/eosn-base-api/log"
	"go.uber.org/zap"
	"google.golang.org/protobuf/types/known/timestamppb"
	"time"
)

const EOSTimeFormat = "2006-01-02T15:04:05"

func ParseEOSTimeToProtobuf(eosTime string) *timestamppb.Timestamp {

	parsedTime, err := time.Parse(EOSTimeFormat, eosTime)
	if err != nil {
		log.Error("failed to parse time from data table", zap.Error(err), zap.String("time", eosTime))
		return &timestamppb.Timestamp{}
	}

	return timestamppb.New(parsedTime)
}
