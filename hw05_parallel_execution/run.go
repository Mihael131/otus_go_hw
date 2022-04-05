package hw05parallelexecution

import (
	"errors"
	"sync"
)

var ErrErrorsLimitExceeded = errors.New("errors limit exceeded")

type Task func() error

type taskWorker struct {
	wg        sync.WaitGroup
	mu        sync.Mutex
	errs      int
	limitErrs int
	limitGors int
}

func (w *taskWorker) addError() {
	w.mu.Lock()
	defer w.mu.Unlock()
	w.errs++
}

func (w *taskWorker) checkErrors() bool {
	w.wg.Wait()
	return w.errs >= w.limitErrs
}

func (w *taskWorker) checkGoroutines(i int) bool {
	return i%w.limitGors == 0
}

func (w *taskWorker) run(task Task) {
	w.wg.Add(1)
	go func() {
		defer w.wg.Done()
		if err := task(); err != nil {
			w.addError()
		}
	}()
}

func Run(tasks []Task, n, m int) error {
	worker := taskWorker{limitGors: n, limitErrs: m}
	for i, task := range tasks {
		worker.run(task)
		if worker.checkGoroutines(i) && worker.checkErrors() {
			return ErrErrorsLimitExceeded
		}
	}
	if worker.checkErrors() {
		return ErrErrorsLimitExceeded
	}
	return nil
}
