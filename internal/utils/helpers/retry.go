package helpers

import (
	"time"

	"github.com/percolate/retry"
)

func GetDefaultRetry() retry.Re {
	return retry.Re{Max: 15, Delay: time.Second * 1}
}
