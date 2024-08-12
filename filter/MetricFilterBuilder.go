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
		logger.Logger.Infof("%v: %v, cost time: %d", ctx.R.Method, ctx.R.URL.Path, endTime-startTime)
		return nil
	}
}
