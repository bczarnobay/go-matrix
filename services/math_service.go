package services

import (
	"sync"
)

func Sum(m [][]int) int {
	if len(m) == 0 {
		return 0
	} else if len(m[0]) == 0 {
		return 0
	}

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
	if len(m) == 0 {
		return 0
	} else if len(m[0]) == 0 {
		return 0
	}

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
