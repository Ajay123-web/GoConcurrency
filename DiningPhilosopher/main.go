package main

import (
	"fmt"
	"sync"
	"time"
)

var philosophers = []string{
	"Aristotle",
	"Pascal",
	"Locke",
	"Plato",
	"Aryabhatt",
}

const hunger = 1


var wg sync.WaitGroup
var sleepTime = 0 * time.Second

var forks = [5]*sync.Mutex{}

var finishOrder = []string{}

func diningProblem(philosopher string , i int) {
	defer wg.Done()
	n := len(philosophers)

	leftFork := i
	rightFork := (i + 1) % n

	if i == n - 1 {
		temp := rightFork
		rightFork = leftFork
		leftFork = temp
	}

	for j := hunger; j >= 0; j-- {
		fmt.Println(philosopher, "is hungry")
        time.Sleep(sleepTime)

		forks[leftFork].Lock()
        fmt.Println(philosopher , "picked up left fork")
		forks[rightFork].Lock()
        fmt.Println(philosopher , "picked up right fork")
        
        fmt.Println(philosopher, "is eating")
        time.Sleep(sleepTime)

        forks[leftFork].Unlock()
        fmt.Println(philosopher, "put down left fork")
        forks[rightFork].Unlock()
        fmt.Println(philosopher, "put down right fork")
        time.Sleep(sleepTime)
	}
	finishOrder = append(finishOrder , philosopher)
}

func main() {

	n := len(philosophers)
	wg.Add(n)

	for i := 0; i < n; i++ {
		forks[i] = &sync.Mutex{}
	}

	for index , philosopher := range philosophers {
		go diningProblem(philosopher , index)
	}

	wg.Wait()

	for _ , philosopher := range finishOrder {
		fmt.Printf("%s " , philosopher)
	}
	fmt.Println()

}