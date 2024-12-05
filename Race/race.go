/*
Write two goroutines which have a race condition when executed concurrently. Explain what the race condition is and how it can occur.

A race condition happens when two or more goroutines access shared data concurrently, and at least one of them is modifying the data.
If the operations are not properly synchronized, the outcome may depend on the order in which the goroutines execute, which is unpredictable.
*/

package main

import (
	"fmt"
	"sync"
)

var wait_gr sync.WaitGroup

var return_text string

func RaceConditionsExample(text string) {
	defer wait_gr.Done()
	return_text = text + ", not all those who wander are lost"
}

func main() {
	wait_gr.Add(2)
	go RaceConditionsExample("Frodo")
	go RaceConditionsExample("Bilbo")
	wait_gr.Wait()

	fmt.Println(return_text)

	/*
		In this case, both goroutines modify the global variable return_text without any synchronization (like mutexes or channels).
		When one goroutine reads string and writes (modifies) it back, the other goroutine might be reading or writing the value at the same time.
		This leads to inconsistent or unexpected results.

		If the goroutines are interleaved in such a way that both read the same value of global var before either writes back their new value,
		we lose one update and therefore instead of returning Bilbo, it can return Frodo.
	*/
}

/*
==================
WARNING: DATA RACE
Write at 0x0001401fb2a0 by goroutine 7:
  main.RaceConditionsExample()
      ./race.go:14 +0xba
  main.main.gowrap1()
      ./Race/race.go:19 +0x35

Previous write at 0x0001401fb2a0 by goroutine 8:
  main.RaceConditionsExample()
      ./Race/race.go:14 +0xba
  main.main.gowrap2()
      ./Race/race.go:20 +0x35

Goroutine 7 (running) created at:
  main.main()
      ./Race/race.go:19 +0x38

Goroutine 8 (finished) created at:
  main.main()
      ./Race/race.go:20 +0x44
==================
Frodo, not all those who wander are lost
Found 1 data race(s)
exit status 66

*/
