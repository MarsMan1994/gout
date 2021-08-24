package filter

import (
	"github.com/MarsMan1994/gout/dataflow"
	"math/rand"
	"time"
)

var (
	defaultBench = Bench{}
	defaultRetry = Retry{}
)

func init() {
	dataflow.Register("bench", &defaultBench)
	dataflow.Register("retry", &defaultRetry)

	rand.Seed(time.Now().UnixNano())
}
