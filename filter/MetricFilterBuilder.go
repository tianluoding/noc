package filter

import (
	"fmt"
	"time"

	"github.com/tianluoding/noc"
)

func MetricFilter(next noc.HandlerFunc) noc.HandlerFunc {
	return func(ctx *noc.Context) error {
		startTime := time.Now().UnixNano()
		next(ctx)
		endTime := time.Now().UnixNano()
		fmt.Printf("cost time: %d\n", endTime-startTime)
		return nil
	}
}
