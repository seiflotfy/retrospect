package main

import (
	"fmt"
	"math/rand"
	"time"

	hindsight "github.com/seiflotfy/hindsight"
)

var obs = hindsight.New("Hello")

func print(o hindsight.Result) {
	v := o.Payload().(*int)
	fmt.Println(o.Elapsed(), o.Count(), *v)
}

func demo(i *int) {
	defer obs.Observe(print, i)()
	<-time.After(time.Millisecond * time.Duration(rand.Int63n(1e3)))
}

func main() {
	for i := 0; i < 10; i++ {
		j := i
		demo(&j)
		j = 0
	}
	summary := obs.Summary()
	fmt.Printf("%s func summary\nLast exec duration %v\nAvg exec duration: %v\nMin exec duration: %v\nMax exec duration: %v\nNumber of exec: %v\n", summary.Namespace(), summary.Last(), summary.Average(), summary.Min(), summary.Max(), summary.Count())
}
