package main

import (
	clog "mohammadinasab-dev/logmodule/logger"
)

type lop struct {
	name string
}

func main() {
	foods := map[string]interface{}{
		"bacon": "delicious",
		"eggs": struct {
			source string
			price  float64
		}{"chicken", 1.75},
		"steak": true,
	}
	// clog.Log.Println("hello log world")
	// clog.Log.WithField("loplp", "holo")
	ll := clog.Log.WithFields(foods)
	ll.Info("message from Info")
	clog.Log.INFO("message from INFO", foods)
}
