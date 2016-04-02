package cloudwatch

import (
	"fmt"
	// "github.com/aws/aws-sdk-go/service/cloudwatch"
	"sort"
	"time"
)

type PointKey struct {
	Timestamp time.Time
	FullName string // Namespace.MetricName.Dimensions.Statistic
}

type TimeSlice []time.Time

func (t TimeSlice) Len() int {
	return len(t)
}

func (t TimeSlice) Less(i, j int) bool {
	return t[i].Before(t[j])
}

func (t TimeSlice) Swap(i, j int) {
	save := t[i]
	t[i] = t[j]
	t[j] = save
}

func (t TimeSlice) Sort() {
	sort.Sort(t)
}

type TimeMap map[time.Time]float64
type DataseriesMap map[string]TimeMap
type TimeSet map[time.Time]bool

// Timestamps
type Datapoints struct {
	Values DataseriesMap
	Times TimeSet
}

func MakeDatapoints() *Datapoints {
	d := &Datapoints{}
	d.Values = make(DataseriesMap)
	d.Times = make(TimeSet)
	return d
}

func (d *Datapoints) AddPoint(time time.Time, 
										full_name string,
										value float64) error {
	dataseries, ok := d.Values[full_name]
	if ! ok {
		d.Values[full_name] = make(TimeMap)
		dataseries = d.Values[full_name]
	}
	dataseries[time] = value
	d.Times[time] = true

	return nil
}

func (d *Datapoints) PrintPoints() {
	var times TimeSlice
	fmt.Printf("Before sort\n")
	for t := range d.Times {
		times = append(times, t)
		fmt.Printf("%v\n", t)
	}

	times.Sort()
	fmt.Printf("After sort\n")
	for _, t := range times {
		fmt.Printf("%v\n", t)
	}


	for m, time_map := range d.Values {
		fmt.Printf("metric = %s\n", m)
		for tm, val := range time_map {
			fmt.Printf("tm = %v, val = %f\n", tm, val)
		}
	}
}
