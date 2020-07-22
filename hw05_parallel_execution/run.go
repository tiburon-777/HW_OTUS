package hw05_parallel_execution //nolint:golint,stylecheck

import (
	"errors"
	"sync"
)

var ErrErrorsLimitExceeded = errors.New("errors limit exceeded")

type Errors struct {
	count int
	mx    sync.RWMutex
}
type Task func() error

func Run(tasks []Task, n int, m int) error {
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
		go consumer(ch, &wg, &errs, ignore)
	}

	for _, task := range tasks {
		errs.mx.RLock()
		if !ignore && errs.count <= 0 {
			errs.mx.RUnlock()
			break
		}
		errs.mx.RUnlock()
		ch <- task
	}
	close(ch)
	wg.Wait()

	if errs.count <= 0 && !ignore {
		return ErrErrorsLimitExceeded
	}
	return nil
}

func consumer(ch <-chan Task, wg *sync.WaitGroup, errs *Errors, ignore bool) {
	defer wg.Done()
	for task := range ch {
		errs.mx.RLock()
		if !ignore && errs.count <= 0 {
			errs.mx.RUnlock()
			return
		}
		errs.mx.RUnlock()
		err := task()
		if !ignore && err != nil {
			errs.mx.Lock()
			errs.count--
			errs.mx.Unlock()
		}
	}
}
