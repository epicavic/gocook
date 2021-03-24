package main

import (
	"fmt"
	"log"
	"net/http"

	"main/metrics"
)

func main() {
	http.HandleFunc("/counter", metrics.CounterHandler)
	http.HandleFunc("/timer", metrics.TimerHandler)
	http.HandleFunc("/report", metrics.ReportHandler)

	fmt.Println("Listening on localhost:8080")
	log.Fatal(http.ListenAndServe("localhost:8080", nil))
}

/*
$ go run main.go
Listening on localhost:8080

$ curl -w'\n' http://localhost:8080/counter
success

$ curl -w'\n' http://localhost:8080/timer
success

$ curl -s http://localhost:8080/report | jq .
{
  "counterhandler.counter": {
    "count": 1
  },
  "reporthandler.writemetrics": {
    "15m.rate": 0.19044224090584513,
    "1m.rate": 0.12163203119558474,
    "5m.rate": 0.17378490931305843,
    "75%": 329318,
    "95%": 354375,
    "99%": 354375,
    "99.9%": 354375,
    "count": 5,
    "max": 354375,
    "mean": 229748.6,
    "mean.rate": 0.07144913274663778,
    "median": 167291,
    "min": 156314,
    "stddev": 82918.70961031654
  },
  "timerhandler.timer": {
    "15m.rate": 0.002079556776971638,
    "1m.rate": 0.013138818043506624,
    "5m.rate": 0.0054736267710136476,
    "75%": 123310,
    "95%": 123310,
    "99%": 123310,
    "99.9%": 123310,
    "count": 1,
    "max": 123310,
    "mean": 64437,
    "mean.rate": 0.024850849200358544,
    "median": 64437,
    "min": 5564,
    "stddev": 58873
  }
}
*/
