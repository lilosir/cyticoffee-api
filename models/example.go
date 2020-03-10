package models

import (
	"fmt"
)

// Job interface
type Job interface {
	Do()
}

// Worker struct
type Worker struct {
	JobQueue chan Job
}

// CreateWorker return new worker
func CreateWorker() Worker {
	return Worker{make(chan Job)}
}

// Run from worker
func (w Worker) Run(wq chan chan Job) {
	go func() {
		wq <- w.JobQueue
		for {
			select {
			case job := <-w.JobQueue:
				job.Do()
			}
		}
	}()
}

// WorkerPool struct
type WorkerPool struct {
	workerLen   int
	JobQueue    chan Job
	WorkerQueue chan chan Job
}

// CreateWorkerPool ..
func CreateWorkerPool(workerLen int) *WorkerPool {
	return &WorkerPool{
		workerLen:   workerLen,
		JobQueue:    make(chan Job),
		WorkerQueue: make(chan chan Job, workerLen),
	}
}

// Run ..
func (wp *WorkerPool) Run() {
	fmt.Println("initialize...")

	for i := 0; i < wp.workerLen; i++ {
		worker := CreateWorker()
		worker.Run(wp.WorkerQueue)
	}

	go func() {
		for {
			select {
			case job := <-wp.JobQueue:
				worker := <-wp.WorkerQueue
				worker <- job
			}
		}
	}()
}
