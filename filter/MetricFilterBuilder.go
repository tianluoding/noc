package filter

import (
	"time"

	"github.com/tianluoding/noc"
	"github.com/tianluoding/noc/logger"
)

func MetricFilter(next noc.HandlerFunc) noc.HandlerFunc {
	return func(ctx *noc.Context) error {
		startTime := time.Now().UnixNano()
		next(ctx)
		endTime := time.Now().UnixNano()
		logger.Logger.Infof("cost time: %d", endTime-startTime)
		return nil
	}
}
