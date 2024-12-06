/* 2024-12-06 - MCK
Implement the dining philosopher’s problem with the following constraints/modifications.

There should be 5 philosophers sharing chopsticks, with one chopstick between each adjacent pair of philosophers.

Each philosopher should eat only 3 times (not in an infinite loop as we did in lecture)

The philosophers pick up the chopsticks in any order, not lowest-numbered first (which we did in lecture).

In order to eat, a philosopher must get permission from a host which executes in its own goroutine.

The host allows no more than 2 philosophers to eat concurrently.

Each philosopher is numbered, 1 through 5.

When a philosopher starts eating (after it has obtained necessary locks) it prints “starting to eat <number>” on a line by itself, where <number> is the number of the philosopher.

When a philosopher finishes eating (before it has released its locks) it prints “finishing eating <number>” on a line by itself, where <number> is the number of the philosopher.
*/

package main

import (
	"fmt"
	"sync"
	"time"
)

var wait_gr sync.WaitGroup

// create channel with buffer of 2. "The host allows no more than 2 philosophers to eat concurrently."
var host = make(chan bool, 2)

//var mutex sync.Mutex

// Each chopstick is a mutex
type ChopS struct{ sync.Mutex }

// Each philosopher is associated with a goroutine and two chopsticks + has an Id
type Philo struct {
	leftCS, rightCS *ChopS
	pId             int
}

func (p Philo) eat() {
	for i := 0; i < 3; i++ {
		//In order to eat, a philosopher must get permission from a host which executes in its own goroutine.
		host <- true

		p.leftCS.Lock()
		p.rightCS.Lock()
		time.Sleep(500 * time.Microsecond)
		//When a philosopher starts eating (after it has obtained necessary locks) it prints “starting to eat <number>” on a line by itself, where <number> is the number of the philosopher.
		fmt.Println("starting to eat", p.pId)
		time.Sleep(1 * time.Second)
		//When a philosopher finishes eating (before it has released its locks) it prints “finishing eating <number>” on a line by itself, where <number> is the number of the philosopher.
		fmt.Println("finishing eating", p.pId)
		time.Sleep(500 * time.Microsecond)
		//fmt.Println()
		p.rightCS.Unlock()
		p.leftCS.Unlock()

		<-host
	}
	wait_gr.Done()
}

func main() {
	CSticks := make([]*ChopS, 5)
	for i := 0; i < 5; i++ {
		CSticks[i] = new(ChopS)
	}
	philos := make([]*Philo, 5)
	for i := 0; i < 5; i++ {
		//Each philosopher is numbered, 1 through 5.
		philos[i] = &Philo{leftCS: CSticks[i], rightCS: CSticks[(i+1)%5], pId: i + 1}
	}
	for i := 0; i < 5; i++ {
		wait_gr.Add(1)

		go philos[i].eat()
	}
	wait_gr.Wait()
}
