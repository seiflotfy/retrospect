package main

import (
	"fmt"
	"math/rand"
	"time"

	retrospect "github.com/seiflotfy/retrospect"
)

var obs = retrospect.New("Hello")

func print(o retrospect.Result) {
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
		j = i * 2
	}
	summary := obs.Summary()
	fmt.Printf("\n%s func summary\n------------------\nLast exec duration %v\nAvg exec duration: %v\nMin exec duration: %v\nMax exec duration: %v\nNumber of exec: %v\n", summary.Namespace(), summary.Last(), summary.Average(), summary.Min(), summary.Max(), summary.Count())
}
