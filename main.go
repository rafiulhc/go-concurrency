package main

import "sync"

func main() {
	w := sync.WaitGroup{}
	myArray := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	w.Add(len(myArray))
	go func() {


		for _, index := range myArray {
			defer w.Done()
			println("Go routine", index)
		}

	}()
	w.Wait()
	print("Main routine")
}