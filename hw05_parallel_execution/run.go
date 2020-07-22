package hw05_parallel_execution //nolint:golint,stylecheck

func Run(tasks []Task, n int, m int) error {
	if m == -1 {
		m = len(tasks)
	}
	pool := make(chan int, n)
	errs := make(chan int)
	for _, task := range tasks {
		pool <- 1
		go func(task Task, errs chan int, pool chan int) {
			if _, ok := <-errs; !ok {
				return
			}
			if task() != nil {
				errs <- 1
			}
			<-pool
		}(task, errs, pool)
		if len(errs) >= m {
			close(errs)
			return ErrErrorsLimitExceeded
		}
	}
	close(errs)
	return nil
}
