package workerpool

import "log"

type WorkerPool interface {
	Run()
	AddTask(task func())
}

type workerPool struct {
	maxWorker   int
	queuedTaskC chan func()
}

func NewWorkerPool(maxWorker int) WorkerPool {
	wp := &workerPool{
		maxWorker:   maxWorker,
		queuedTaskC: make(chan func(), maxWorker),
	}

	return wp
}

func (wp *workerPool) Run() {
	wp.run()
}

func (wp *workerPool) AddTask(task func()) {
	wp.queuedTaskC <- task
}

func (wp *workerPool) GetTotalQueuedTask() int {
	return len(wp.queuedTaskC)
}

func (wp *workerPool) run() {
	for i := 0; i < wp.maxWorker; i++ {
		wID := i + 1
		log.Printf("[WorkerPool] Worker %d has been spawned", wID)

		go func(workerId int) {
			for task := range wp.queuedTaskC {
				log.Printf("[WorkerPool] Worker %d get task", wID)
				task()
				log.Printf("[WorkerPool] Worker %d finish task", wID)
			}
		}(wID)
	}
}
