package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

type Data struct {
	Value int
}

func main() {
	ch := make(chan Data)

	var wg sync.WaitGroup

	wg.Add(2)

	go func() {
		defer wg.Done()
		reciveData(ch)
	}()

	go func() {
		defer wg.Done()
		processData(ch)
	}()

	wg.Wait()

	close(ch)
}

// Funcktion to generate data
func getData() Data {
	return Data{Value: rand.Intn(100)}
}

// Funkction to recive Data
func reciveData(ch chan<- Data) {
	for i := 0; i < 5; i++ {
		data := getData()
		fmt.Printf("Recived data: %d\n", data.Value)
		ch <- data
		time.Sleep(time.Second)
	}
}

// Function for process data
func processData(ch <-chan Data) {
	for data := range ch {
		fmt.Printf("Processed data: %d\n", data.Value)
	}
}

	