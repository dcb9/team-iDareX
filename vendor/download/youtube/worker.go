package youtube

import (
	"fmt"
)

type Worker struct{
	ID int
	Work chan WorkRequest
	WorkerQueue chan chan WorkRequest
	QuitChan chan bool
}

func NewWroker(id int, workerQueue chan chan WorkRequest) Worker {
	// Create, and return the worker.
	
	worker := Worker{
		ID: id,
		Work: make(chan WorkRequest),
		WorkerQueue: workerQueue,
		QuitChan: make(chan bool)}
	
	return worker
}

func (w Worker) Start() {
	go func() {
		for {
			// Add ourselves into the worker queue.
			w.WorkerQueue <- w.Work
			
			select {
			case work := <-w.Work:
				// Receive a work request.
				fmt.Printf("worker%d: Received work request\n", w.ID)
				fmt.Printf("worker %d: Url: %s\n", w.ID, work.Url)
			case <-w.QuitChan:
				// We have been asked to stop.
				fmt.Printf("worker %d stopping\n", w.ID)
			}
		}
	}()
}

// Stop tells the worker to stop listening for work requests.
// 
// Note that the worker will only stop *after* it has finished its work.
func (w Worker) Stop() {
	go func() {
		w.QuitChan <- true
	}()
}