package hw05parallelexecution

import (
	"errors"
	"sync"
)

var (
	ErrErrorsLimitExceeded    = errors.New("errors limit exceeded")
	ErrErrorsNotEnoughWorkers = errors.New("errors limit exceeded")
)

type Task func() error

// Run starts tasks in n goroutines and stops its work when receiving m errors from tasks.
func Run(tasks []Task, n, m int) error {
	if n <= 0 {
		return ErrErrorsNotEnoughWorkers
	}
	if m <= 0 {
		return ErrErrorsLimitExceeded
	}
	if len(tasks) == 0 {
		return nil
	}

	wg := sync.WaitGroup{}
	mu := sync.Mutex{}
	ch := make(chan Task, len(tasks))
	errCount := 0

	for _, task := range tasks {
		ch <- task
	}
	close(ch)

	for i := 0; i < n; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for v := range ch {
				if err := v(); err != nil {
					mu.Lock()
					if errCount >= m {
						mu.Unlock()
						break
					}
					errCount++
					mu.Unlock()
				}
			}
		}()
	}

	wg.Wait()

	if errCount >= m {
		return ErrErrorsLimitExceeded
	}
	return nil
}
