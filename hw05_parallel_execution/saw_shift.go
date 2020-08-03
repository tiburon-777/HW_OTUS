package hw05_parallel_execution //nolint:golint,stylecheck

import (
	"sync"
)

// Функция, решающая задачу методом пилообразного сдвига.
func RunSawShift(tasks []Task, n int, m int) error {
	errs := Errors{}
	for i := 0; i < len(tasks); i += n {
		wg := sync.WaitGroup{}
		for g := 1; g <= n && i+g < len(tasks); g++ {
			wg.Add(1)
			go func(rt Task, errs *Errors) {
				defer wg.Done()
				err := rt()
				if err != nil {
					errs.mx.Lock()
					errs.count++
					errs.mx.Unlock()
				}
			}(tasks[i+g], &errs)
		}
		wg.Wait()
		if m > 0 && errs.count > m {
			return ErrErrorsLimitExceeded
		}
	}
	return nil
}
