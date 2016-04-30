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

func (ts TimeSlice) Len() int {
	return len(ts)
}

func (ts TimeSlice) Less(i, j int) bool {
	return ts[i].Before(ts[j])
}

func (ts TimeSlice) Swap(i, j int) {
	save := ts[i]
	ts[i] = ts[j]
	ts[j] = save
}

func (ts TimeSlice) Sort() {
	sort.Sort(ts)
}

func (ts TimeSlice) AppendFromSet(inset TimeSet) TimeSlice {
	for key, _ := range inset {
		ts = append(ts, key)
	}
	return ts
}

func (ts TimeSlice) Print(delimiter string, format string) {
	for _, tm := range ts {
		fmt.Printf("%s%s", delimiter, tm.Format(format))
	}
	fmt.Printf("\n")
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
	times = times.AppendFromSet(d.Times)

	times.Sort()
	times.Print(",", time.RFC3339)
	fmt.Printf("\n")

	for m, time_map := range d.Values {
		fmt.Printf("metric = %s\n", m)
		for tm, val := range time_map {
			fmt.Printf("tm = %v, val = %f\n", tm, val)
		}
	}
}

func (d *Datapoints) PrintCSV(delimiter string) {
	// Format to print out
	// Time,t1,t2,...,tn
	// DS1,v11,v12,...,v1n
	// DS2,v21,v22,...,v2n
	var times TimeSlice
	times = times.AppendFromSet(d.Times)
	times.Sort()

	// Print timestamps row first
	fmt.Printf("Time")
	times.Print(",", time.RFC3339)

	// Print each dataseries now
	for m, time_map := range d.Values {
		fmt.Printf("%s", m)
		for _, tm := range times {
			if val, ok := time_map[tm]; ok {
				fmt.Printf("%s%f", delimiter, val) 
			} else {
				fmt.Printf("%s", delimiter)
			}
		}
		fmt.Printf("\n")
	}
}
