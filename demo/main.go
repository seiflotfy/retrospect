package main

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/seiflotfy/observer"
)

var obs = observer.New("Hello")

func print(id string, elapsed time.Duration, count uint64) {
	fmt.Println(id, elapsed, count)
}

func demo() {
	defer obs.Measure("demo", print)()
	<-time.After(time.Millisecond * time.Duration(rand.Int63n(1e3)))
}

func main() {
	id := "demo"
	for i := 0; i < 10; i++ {
		demo()
	}
	report := obs.Get(id)
	fmt.Printf("%s func report\nLast exec duration %v\nAvg exec duration: %v\nMin exec duration: %v\nMax exec duration: %v\nNumber of exec: %v\n",
		report.ID(), report.Last(), report.Average(), report.Min(), report.Max(), report.Count())
}
