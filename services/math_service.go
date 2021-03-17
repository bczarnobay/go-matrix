package services

import (
	"sync"
)

func Sum(m [][]int) int {
	var sum int
	var wg sync.WaitGroup
	var lock sync.Mutex

	for _, row := range m {
		wg.Add(1)
		go func(row []int, wg *sync.WaitGroup) {
			defer wg.Done()
			for _, num := range row {
				lock.Lock()
				sum += num
				lock.Unlock()
			}
		}(row, &wg)
	}
	wg.Wait()

	return sum
}

func Multiply(m [][]int) int {
	result := 1

	var wg sync.WaitGroup
	var lock sync.Mutex

	for _, row := range m {
		wg.Add(1)
		go func(row []int, wg *sync.WaitGroup) {
			defer wg.Done()

			for _, num := range row {
				lock.Lock()
				result *= num
				lock.Unlock()
			}
		}(row, &wg)
	}
	wg.Wait()
	return result
}
