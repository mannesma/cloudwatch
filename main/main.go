package main

import (
	"github.com/mannesma/cloudwatch"
	"os"
)

func main() {
   request := cloudwatch.ParseArgs()
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
