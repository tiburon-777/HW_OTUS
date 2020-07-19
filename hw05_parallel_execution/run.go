package hw05_parallel_execution //nolint:golint,stylecheck

import (
	"errors"
	"log"
	"sync"
)

var ErrErrorsLimitExceeded = errors.New("errors limit exceeded")

type Task func() error

// Run starts tasks in N goroutines and stops its work when receiving M errors from tasks
func Run(tasks []Task, N int, M int) error {
	log.Println("Tasks:", len(tasks), "| Goroutines:", N, "| Errors:", M)
	errs := 0
	for i := 0; i < len(tasks); i = i + N {
		wg := sync.WaitGroup{}
		for g := 1; g <= N && i+g < len(tasks); g++ {
			wg.Add(1)
			go func(rt Task, i int, g int, errs *int) {
				if err := rt; err != nil {
					*errs++
				}
				wg.Done()
			}(tasks[i+g], i, g, &errs)
		}
		wg.Wait()
		if errs > M {
			log.Println("Produced", errs, "errors of", M)
			return ErrErrorsLimitExceeded
		}
	}
	return nil
}
