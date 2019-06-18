package main

import (
	"fmt"
	"math"
	"math/rand"
	"sync"
	"time"
)

const timesEat int = 3
const numberPhilos int = 5
const concurrentEaters int = 2

var currentEaters []int

type ChopS struct{ sync.Mutex }

type Philo struct {
	idNumber        int
	leftCS, rightCS *ChopS
}

func host(c chan) {
	for i := 0; i < numberPhilos; i++ {
		
	}
}

func isNeighbor(a, b int) bool {
	return math.Abs(float64(a-b)) == 1 || (b == numberPhilos && a == 0) || (a == numberPhilos && b == 0)
}

func getPermissionFromHost(id int) bool {

}

func (p Philo) eat(wg *sync.WaitGroup, c chan) {
	if getPermissionFromHost(p.idNumber) {
		x := rand.Intn(2)
		if x == 1 {
			p.leftCS.Lock()
			// fmt.Printf("Left chopstick locked for %d\n", p.idNumber)
			p.rightCS.Lock()
			// fmt.Printf("Right chopstick locked for %d\n", p.idNumber)
		} else {
			p.rightCS.Lock()
			// fmt.Printf("Right chopstick locked for %d\n", p.idNumber)
			p.leftCS.Lock()
			// fmt.Printf("Left chopstick locked for %d\n", p.idNumber)
		}
		idNumber := p.idNumber + 1
		fmt.Printf("starting to eat %d\n", idNumber)
		fmt.Printf("finishing eating %d\n", idNumber)

		p.rightCS.Unlock()
		// fmt.Printf("Right chopstick unlocked for %d\n", p.idNumber)
		p.leftCS.Unlock()
		// fmt.Printf("Left chopstick unlocked for %d\n", p.idNumber)
		wg.Done()
	}
}

func main() {
	var wg sync.WaitGroup
	rand.Seed(time.Now().UnixNano())
	c := make(chan int)
	currentEaters = make([]int, 0)
	CSticks := make([]*ChopS, numberPhilos)
	for i := 0; i < numberPhilos; i++ {
		CSticks[i] = new(ChopS)
	}
	philos := make([]*Philo, numberPhilos)
	for i := 0; i < numberPhilos; i++ {
		philos[i] = &Philo{i, CSticks[i], CSticks[(i+1)%numberPhilos]}
	}
	wg.Add(numberPhilos*timesEat + 1)
	for i := 0; i < numberPhilos; i++ {
		for j := 0; j < timesEat; j++ {
			go philos[i].eat(&wg, c)
		}
	}
	go host(c)
	wg.Wait()
}
