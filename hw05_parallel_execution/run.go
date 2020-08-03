package hw05_parallel_execution //nolint:golint,stylecheck
import "sync"

func Run(tasks []Task, n int, m int) error {
	if m == -1 {
		m = len(tasks)
	}
	task, errs, wg, done := make(chan Task), make(chan error, len(tasks)), sync.WaitGroup{}, make(chan struct{})
	defer func() {
		close(done)
		wg.Wait()
	}()
	wg.Add(n)
	for i := 1; i <= n; i++ {
		go func(task chan Task, errs chan error, wg *sync.WaitGroup, done chan struct{}) {
			for {
				select {
				case <-done:
					wg.Done()
					return
				case t := <-task:
					if err := t(); err != nil {
						errs <- err
					}
				}
			}
		}(task, errs, &wg, done)
	}
	for _, t := range tasks {
		task <- t
		if len(errs) >= m {
			return ErrErrorsLimitExceeded
		}
	}
	return nil
}
