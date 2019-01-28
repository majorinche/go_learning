package main

import (
	"fmt"
)

type sensorData struct {
	ID    int64
	Name  string
	Type  string
	State string
	Value float64
	Unit  string
	Event string
}

func main() {
	var result []sensorData // struct can be a list
	var data sensorData     // struct can also be a single value, really intresting
	data.ID = 111
	data.Name = "majorin"
	result = append(result, data) // so we don't need to define every variable value
	data.ID = 222
	data.Name = "jinguo"
	result = append(result, data)
	fmt.Println("aaa: %s\n", result[0]) // i just wannt print out firt element out
}
