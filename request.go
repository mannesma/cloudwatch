package cloudwatch

import (
	"encoding/json"
	"flag"
	"fmt"
   "github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/cloudwatch"
	"io/ioutil"
	"time"
)

type Metric struct {
   MetricName *string                   `json:"MetricName"`
   Namespace *string                    `json:"Namespace"`
   Statistics []*string                 `json:"Statistics"`
   Dimensions []*cloudwatch.Dimension   `json:"Dimensions,omitempty"`
   Unit *string                         `json:"Unit,omitempty"`
}

type MetricList []Metric

type Request struct {
   Period int64
   StartTime *time.Time
   EndTime *time.Time
   Metrics *MetricList
}

func parseDate(instr string, default_offset time.Duration) (*time.Time, error) {
   // default_offset is used with curr time if instr is blank
   retDate := time.Now()
   if instr != "" {
      // Try to see if it is a duration first
      offset, err := time.ParseDuration(instr)
      if err == nil {
         retDate = retDate.Add(offset)
      } else {
         // Now try RFC 3339 format
         retDate, err = time.Parse(time.RFC3339, instr)
         if err != nil {
            return nil, err
         }
      }
   } else {
      // Just add default_offset to current time
      retDate = retDate.Add(default_offset)
   }

   return &retDate, nil
}

func parseMetrics(metricFile string) *MetricList {
   var ml MetricList
   // Read up the file
   data, err := ioutil.ReadFile(metricFile)
   if err != nil {
      fmt.Printf("Error: Failed to open file %s: %s\n", metricFile, err)
      return nil
   }
   err = json.Unmarshal(data, &ml)
   if err != nil {
      fmt.Printf("Error: Failed to unmarshal json: %s\n", err)
      return nil
   }

   return &ml
}

func ParseArgs() *Request {
   var err error
   r := &Request{}
   flag.Int64Var(&r.Period, "period", 60, "Number of seconds in each period")
   startTime := flag.String("start_time", "", "Start time in RFC 3339 format")
   endTime := flag.String("end_time", "", "End time in RFC 3339 format")

   metricFile := flag.String("metric_file", "metrics.json", "The filename of the desired metrics to be queried")

   flag.Parse()

   r.StartTime, err = parseDate(*startTime, -300*time.Second)
   if err != nil {
      fmt.Printf("Error: Failed to parse start time: %s\n", err)
      return nil
   }
   r.EndTime, err = parseDate(*endTime, 0*time.Second)
   if err != nil {
      fmt.Printf("Error: Failed to parse end time: %s\n", err)
      return nil
   }

   fmt.Printf("metricFile = %s\n", *metricFile)
   r.Metrics = parseMetrics(*metricFile)

   return r
}

func DoRequest(request *Request) {
   // Set up common items in cloudwatch input struct
   params := &cloudwatch.GetMetricStatisticsInput{
      StartTime: aws.Time(*request.StartTime), // Required
      Period:    aws.Int64(request.Period),    // Required
      EndTime:   aws.Time(*request.EndTime),   // Required
   }

   // Make a new cloudwatch session
   svc := cloudwatch.New(session.New())

   for _, m := range *request.Metrics {
      params.Namespace = m.Namespace
      params.MetricName = m.MetricName
      params.Statistics = m.Statistics
      params.Dimensions = m.Dimensions
      params.Unit = m.Unit
      resp, err := svc.GetMetricStatistics(params)
      if err != nil {
         fmt.Printf("Error: Failed to get metric data: %s\n", err)
         continue
      }
      // TODO: Add to csv structure later
      fmt.Println(resp)
   }

}

