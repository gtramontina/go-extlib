package hash_test

import (
	"os"
	"strconv"
)

// maxCountScale reads the `MAX_COUNT_SCALE` environment variable and parses it
// as float64 if available. It defaults to 0.1 otherwise.
// FIXME: maxCountScale is a workaround to the `-quickchecks` flag.
func maxCountScale() float64 {
	maxCountScaleEnv := os.Getenv("MAX_COUNT_SCALE")
	if len(maxCountScaleEnv) > 0 {
		maxCountScale, err := strconv.ParseFloat(maxCountScaleEnv, 64)
		if err != nil {
			panic(err)
		}
		return maxCountScale
	}

	return 0.1
}
