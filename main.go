package main

import (
	"fmt"
	"runtime"
	"sync"
)

func concurrentFactorial(n int64, numThreads int) (int64, error) {
	chunkSize := (n + int64(numThreads) - 1) / int64(numThreads) //формула равномерного распределение между потоками
	var err error
	var wg sync.WaitGroup

	results := make(chan int64, numThreads)
	for i := int64(0); i < int64(numThreads); i++ {
		wg.Add(1)
		go func(startIndex int64) {
			defer wg.Done()
			localResult := int64(1)
			end := min(startIndex*chunkSize+chunkSize, n+1)

			for j := startIndex * chunkSize; j < end; j++ {
				if j > 0 {
					localResult *= j
				}
			}
			results <- localResult
		}(i)
	}

	wg.Wait()
	close(results)

	finalRes := int64(1)
	for res := range results {
		finalRes *= res
	}

	return finalRes, err
}

func min(a, b int64) int64 {
	if a < b {
		return a
	}
	return b
}

func main() {
	var num int64
	fmt.Scanln(&num)
	numThreads := runtime.NumCPU() // под количество ядер в моей системе

	res, err := concurrentFactorial(num, numThreads)

	if err != nil {
		fmt.Println("Ошибка:", err)
		return
	}

	fmt.Println("Факториал =", res)
}
