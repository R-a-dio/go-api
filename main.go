package main

import (
	"flag"
	"log"
	"sync/atomic"
)

var (
	apiMain       atomic.Value
	apiQueue      atomic.Value
	apiLastPlayed atomic.Value
)

func main() {
	flag.Parse()
	if flag.Arg(0) == "" {
		log.Fatalf("no Data Source Name for database connection found as argument")
	}

	if err := openDatabase(flag.Arg(0)); err != nil {
		log.Fatalf("database init error: %s", err)
	}

	startUpdater()
	if err := startServer(); err != nil {
		log.Fatalf("critical server error: %s", err)
	}
}
