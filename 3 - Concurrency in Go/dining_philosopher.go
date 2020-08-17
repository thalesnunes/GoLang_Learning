package main

import (
	"fmt"
	"sync"
)

type Philo struct {
	number          int
	leftCS, rightCS *ChopS
}

type ChopS struct{ sync.Mutex }

func (p *Philo) eat(buffChan *chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 0; i < 3; i++ {
		p.leftCS.Lock()
		p.rightCS.Lock()
		<-*buffChan
		fmt.Printf("starting to eat %d\n", p.number)
		p.leftCS.Unlock()
		p.rightCS.Unlock()
		fmt.Printf("finishing eating %d\n", p.number)
	}
}

func Initializer() []*Philo {
	CSticks := make([]*ChopS, 5)
	philos := make([]*Philo, 5)
	for n := 0; n < 5; n++ {
		CSticks[n] = new(ChopS)
	}
	for n := 0; n < 5; n++ {
		philos[n] = &Philo{n + 1, CSticks[n], CSticks[(n+1)%5]}
	}
	return philos
}

func host(buffChan *chan int) {
	for {
		*buffChan <- 1
		*buffChan <- 2
	}
}

func main() {

	var wg sync.WaitGroup
	buffChan := make(chan int, 2)
	philos := Initializer()
	go host(&buffChan)
	wg.Add(5)
	for i := 0; i < 5; i++ {
		go philos[i].eat(&buffChan, &wg)
	}
	wg.Wait()

}
