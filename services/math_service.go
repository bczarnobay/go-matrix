package services

import (
	"strconv"
	"sync"
)

func Sum(m [][]string) int {
	var sum int
	var wg sync.WaitGroup
	var lock sync.Mutex

	for _, row := range m {
		wg.Add(1)
		go func(row []string, wg *sync.WaitGroup) {
			defer wg.Done()
			for _, num := range row {
				value, err := strconv.Atoi(num)
				if err != nil {
					return
				}
				lock.Lock()
				sum += value
				lock.Unlock()
			}
		}(row, &wg)
	}
	wg.Wait()

	return sum
}

func Multiply(m [][]string) int {
	result := 1

	var wg sync.WaitGroup
	var lock sync.Mutex

	for _, row := range m {
		wg.Add(1)
		go func(row []string, wg *sync.WaitGroup) {
			defer wg.Done()

			for _, num := range row {
				value, err := strconv.Atoi(num)
				if err != nil {
					return
				}
				lock.Lock()
				result *= value
				lock.Unlock()
			}
		}(row, &wg)
	}
	wg.Wait()
	return result
}
