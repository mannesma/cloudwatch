package main

import (
   "fmt"
	"github.com/mannesma/cloudwatch"
	"os"
)

func main() {
   request := cloudwatch.ParseArgs()
   fmt.Printf("request = %v\n", request)
   fmt.Printf("metrics = %v\n", *request.Metrics)
   request.Do()

   os.Exit(0)

   // *** Send to cloudwatch ***
   // {
   //   "Namespace": "",
   //   "MetricName": "",
   //   "Dimensions": [
   //      {
   //          "Name": "",
   //          "Value": ""
   //      }
   //   ],
   //   "StartTime": null,
   //   "EndTime": null,
   //   "Period": 0,
   //   "Statistics": [
   //      ""
   //   ],
   //   "Unit": ""
   // },
}
