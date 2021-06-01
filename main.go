package main

import (
	log "mohammadinasab-dev/logmodule/logger"
)

type lop struct {
	name string
}

func main() {
	foods := map[string]interface{}{
		"baghali": "delicious",
		"eggs": struct {
			Source string
			Price  float64
		}{
			Source: "chicken",
			Price:  1.75,
		},
		"steak": true,
	}

	//clog.Log.INFO("message from INFO", foods)
	log.Info("first info message", foods)
	log.Debug("debug", foods)
	log.Debug("second log message", foods)

}

// fruite := struct {
// 	Name string
// 	Type string
// }{
// 	Name: "Sammy",
// 	Type: "Shark",
// }

// fruitesMap := map[string]interface{}{
// 	"type": "summer",
// 	"fruite": struct {
// 		Name string
// 		Type string
// 	}{
// 		Name: "holo",
// 		Type: "waterful",
// 	},
// }
// //
// cl := clog.Log.WithFields(logrus.Fields{
// 	"fields": fruite,
// })
// cl.Info("logrus.fields with struct as value")

// pl := clog.Log.WithFields(fruitesMap)
// pl.Info("withe a map as input")
