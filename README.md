# Hindsight

***NOTE** This is a stupid simple helper for debugging and execution time measuring purposes*

Just add something like ```defer hs.Observe(callback, payload)()``` to the beginning of a function
```callback``` ==> a function with the signature ```func(o hindsight.Result)```
```payload``` ==> just an interface{} which will be part of the ```hindsight.Result```
```hindsight.Result``` will give you access to the the properties:
* **Count**: how many times the function was called
* **Payload**: the payload passed to observe
* **Elapsed**: how did it take the function to execute

## Example
```go
package main

import (
	"fmt"
	"math/rand"
	"time"

	hindsight "github.com/seiflotfy/hindsight"
)

var hs = hindsight.New("Hello")

func print(o hindsight.Result) {
	v := o.Payload().(*int)
	fmt.Println(o.Elapsed(), o.Count(), *v)
}

func demo(i *int) {
	defer hs.Observe(print, i)() // here we go
	<-time.After(time.Millisecond * time.Duration(rand.Int63n(1e3)))
}

func main() {
	for i := 0; i < 10; i++ {
		j := i
		demo(&j)
		j = i*2
	}
	summary := hs.Summary()
	fmt.Printf("%s func summary\nLast exec duration %v\nAvg exec duration: %v\nMin exec duration: %v\nMax exec duration: %v\nNumber of exec: %v\n", summary.Namespace(), summary.Last(), summary.Average(), summary.Min(), summary.Max(), summary.Count())
}
```

will yield something like 

```
410.162803ms 1 0
551.204403ms 2 2
821.197889ms 3 4
51.197652ms 4 6
937.184059ms 5 8
320.181371ms 6 10
758.303625ms 7 12
148.187807ms 8 14
216.187312ms 9 16

Hello func summary
------------------
Last exec duration 216.187312ms
Avg exec duration: 468.200769ms
Min exec duration: 51.197652ms
Max exec duration: 937.184059ms
Number of exec: 9
```
