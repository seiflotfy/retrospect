package main

import (
	"fmt"
	"math/rand"
	"time"

	hindsight "github.com/seiflotfy/hindsight"
)

var obs = hindsight.New("Hello")

func print(o hindsight.Observation) {
	fmt.Println(o.ID(), o.Elapsed(), o.Count(), o.Payload())
}

func demo(i int) {
	_, mfunc := obs.Observe(print, i)
	defer mfunc(nil)
	<-time.After(time.Millisecond * time.Duration(rand.Int63n(1e3)))
}

func main() {
	for i := 0; i < 10; i++ {
		demo(i)
	}
	summary := obs.Summary()
	fmt.Printf("%s func summary\nLast exec duration %v\nAvg exec duration: %v\nMin exec duration: %v\nMax exec duration: %v\nNumber of exec: %v\n",
		summary.Namespace(), summary.Last(), summary.Average(), summary.Min(), summary.Max(), summary.Count())
}
