package hw05_parallel_execution //nolint:golint,stylecheck

import (
	"errors"
	"log"
	"sync"
)

var ErrErrorsLimitExceeded = errors.New("errors limit exceeded")

type Task func() error
type Errors struct {
	count int
	mx    sync.Mutex
}

func Run(tasks []Task, n int, m int) error {
	log.Println("Tasks:", len(tasks), "| Goroutines:", n, "| Errors:", m)
	errs := Errors{}
	for i := 0; i < len(tasks); i += n {
		wg := sync.WaitGroup{}
		for g := 1; g <= n && i+g < len(tasks); g++ {
			wg.Add(1)
			go func(rt Task, errs *Errors) {
				err := rt
				if err != nil {
					errs.mx.Lock()
					errs.count++
					errs.mx.Unlock()
				}
				wg.Done()
			}(tasks[i+g], &errs)
		}
		wg.Wait()
		if errs.count > m {
			log.Println("Produced", errs.count, "errors of", m)
			return ErrErrorsLimitExceeded
		}
	}
	return nil
}
