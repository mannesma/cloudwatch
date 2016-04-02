package cloudwatch

import (
	// "fmt"
	"testing"
	"time"
)

func TestMe(t *testing.T) {
	d := MakeDatapoints() 

	ts1 := time.Now()
	ts2 := ts1.Add(60*time.Second)
	ts3 := ts2.Add(60*time.Second)
	ts4 := ts3.Add(60*time.Second)
	ts5 := ts4.Add(60*time.Second)

	var metric1 string = "AWS/EC2.CPUUtilization.InstanceId:i-12345678.Average"
	var metric2 string = "AWS/EC2.CPUUtilization.InstanceId:i-23456789.Average"
	var metric3 string = "AWS/EC2.CPUUtilization.InstanceId:i-23456789.Sum"

	d.AddPoint(ts5, metric1, 7.14)
	d.AddPoint(ts1, metric1, 3.14)
	d.AddPoint(ts2, metric1, 4.14)
	d.AddPoint(ts3, metric1, 5.14)
	d.AddPoint(ts4, metric1, 6.14)

	d.AddPoint(ts2, metric2, 12.12)
	d.AddPoint(ts2, metric3, 13.12)

	d.AddPoint(ts3, metric2, 12.13)
	d.AddPoint(ts3, metric3, 13.13)

	d.AddPoint(ts5, metric2, 12.15)
	d.AddPoint(ts5, metric3, 13.15)

	d.PrintPoints()
}
