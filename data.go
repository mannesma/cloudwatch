package cloudwatch

import (
	// "github.com/aws/aws-sdk-go/service/cloudwatch"
	"time"
)

type PointKey struct {
	Timestamp *time.Time
	MetricName *string
}

type Datapoints struct {
	Keys []PointKey
	Values map[PointKey]*float64
}

func MakeDatapoints() *Datapoints {
	d := &Datapoints{}
	return d
}
