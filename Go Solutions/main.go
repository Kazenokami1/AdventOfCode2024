package main

import (
	"fmt"
	"os"
	"reflect"
	"strings"
	"time"
)

func main() {
	start := time.Now()
	d := Days{}
	daysToRun := strings.Split(os.Getenv("DAYSTORUN"), ",")
	for _, day := range daysToRun {
		dMethod := reflect.ValueOf(d).MethodByName(day)
		dMethod.Call(nil)
	}
	duration := time.Since(start)
	fmt.Print("Time Since Start: ")
	fmt.Println(duration)
}
