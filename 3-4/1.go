package main

import (
	"fmt"
	"math"
	"math/rand"
	"sync"
	"time"
)

const (
	eatingTime       = 300
	timesEat         = 3
	numberPhilos     = 5
	concurrentEaters = 2
)

var maxMeals int
var currentEaters []int

// ChopS Chopstick
type ChopS struct{ sync.Mutex }

var currentEatersLock sync.Mutex

// Philo Philosopher
type Philo struct {
	idNumber        int
	leftCS, rightCS *ChopS
}

func contains(s []int, e int) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}

func isNeighbor(a, b int) bool {
	return math.Abs(float64(a-b)) == 1 || (b == numberPhilos && a == 0) || (a == numberPhilos && b == 0)
}

func isAllowedToEat(currentEaters []int, id int) bool {
	if len(currentEaters) < concurrentEaters {
		if contains(currentEaters, id) {
			return false
		}
		for _, eaterID := range currentEaters {
			if isNeighbor(eaterID, id) {
				return false
			}
		}
		return true
	}
	return false
}

func host(wg *sync.WaitGroup, currentEaters *[]int, requestChan chan int, responseChan chan bool) {
	var meals int
	for {
		if meals == maxMeals {
			wg.Done()
		}
		id := <-requestChan
		answer := isAllowedToEat(*currentEaters, id)
		if answer {
			currentEatersLock.Lock()
			*currentEaters = append(*currentEaters, id)
			currentEatersLock.Unlock()
			meals++
		}
		responseChan <- answer
	}
}

func removeByValue(slice []int, value int) []int {
	output := []int{}
	for _, v := range slice {
		if v != value {
			output = append(output, v)
		}
	}
	return output
}

func (p Philo) eat(wg *sync.WaitGroup, currentEaters *[]int, requestChan chan int, responseChan chan bool) {
	for {
		requestChan <- p.idNumber
		permitted := <-responseChan
		if permitted {
			x := rand.Intn(2)
			if x == 1 {
				p.leftCS.Lock()
				p.rightCS.Lock()
			} else {
				p.rightCS.Lock()
				p.leftCS.Lock()
			}
			idNumber := p.idNumber + 1
			fmt.Printf("starting to eat %d\n", idNumber)
			time.Sleep(eatingTime * time.Millisecond)
			fmt.Printf("finishing eating %d\n", idNumber)

			p.rightCS.Unlock()
			p.leftCS.Unlock()
			currentEatersLock.Lock()
			*currentEaters = removeByValue(*currentEaters, p.idNumber)
			currentEatersLock.Unlock()
			wg.Done()
			return
		}
	}
}

func main() {
	var wg sync.WaitGroup
	rand.Seed(time.Now().UnixNano())
	maxMeals = numberPhilos * timesEat
	requestChan := make(chan int)
	responseChan := make(chan bool)
	currentEaters = make([]int, 0)
	CSticks := make([]*ChopS, numberPhilos)
	for i := 0; i < numberPhilos; i++ {
		CSticks[i] = new(ChopS)
	}
	philos := make([]*Philo, numberPhilos)
	for i := 0; i < numberPhilos; i++ {
		philos[i] = &Philo{i, CSticks[i], CSticks[(i+1)%numberPhilos]}
	}
	wg.Add(maxMeals + 1)
	for i := 0; i < numberPhilos; i++ {
		for j := 0; j < timesEat; j++ {
			go philos[i].eat(&wg, &currentEaters, requestChan, responseChan)
		}
	}
	go host(&wg, &currentEaters, requestChan, responseChan)
	wg.Wait()
}
