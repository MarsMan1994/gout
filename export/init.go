package export

import (
	"github.com/MarsMan1994/gout/dataflow"
)

var (
	defaultCurl = Curl{}
)

func init() {
	dataflow.Register("curl", &defaultCurl)
}
