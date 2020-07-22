package hw05_parallel_execution //nolint:golint,stylecheck

import (
	"sync"
)

func RunMixed(tasks []Task, n int, m int) error {
	ch := make(chan Task)
	wg := sync.WaitGroup{}
	errs := Errors{}
	errs.count = m
	var ignore bool
	if errs.count < 0 {
		ignore = true
	}

	wg.Add(n)
	for i := 0; i < n; i++ {
		go consMixed(ch, &wg, &errs, ignore)
	}

	for _, task := range tasks {
		if !ignore && errs.count <= 0 {
			break
		}
		ch <- task
	}
	close(ch)
	wg.Wait()

	if errs.count <= 0 && !ignore {
		return ErrErrorsLimitExceeded
	}
	return nil
}

func consMixed(ch <-chan Task, wg *sync.WaitGroup, errs *Errors, ignore bool) {
	defer wg.Done()
	for task := range ch {
		if !ignore && errs.count <= 0 {
			return
		}
		err := task()
		if !ignore && err != nil {
			errs.count--
		}
	}
}
