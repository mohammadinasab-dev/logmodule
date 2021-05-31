package main

import (
	"fmt"
)

type lop struct {
	name string
}

func main() {
	// clog.Log.Println("hello log world")
	// ll := clog.Log.WithFields(logrus.Fields{
	// 	"extra_field_one": "extra_value_one",
	// })

	// ll.Info("lop lop")
	// clog.Log.Info(lop{
	// 	name: "lop lop",
	// })

	// clog.Log.Info(lop{
	// 	name: "polpol",
	// }, "lop lop")

	type fieldkey string

	// var (
	// 	Action   fieldkey
	// 	response fieldkey
	// )

	type fieldmap map[fieldkey]interface{}

	// mapoffield := fieldmap{
	// 	Action:   "name",
	// 	response: "true",
	// }

	foods := map[string]interface{}{
		"bacon": "delicious",
		"eggs": struct {
			source string
			price  float64
		}{"chicken", 1.75},
		"steak": true,
	}

	//map[struct]

	//printcustom(mapoffield)
	printcustom(foods)
	printcustom(map[string]interface{}{
		"holo": "loplop",
	})

}

var foods map[string]interface{}

func printcustom(foods map[string]interface{}) {
	for k, v := range foods {
		fmt.Println(k, v)
	}
}
