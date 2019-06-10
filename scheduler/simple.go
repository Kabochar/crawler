package scheduler

import "crawler/engine"

type SimpleScheduler struct {
	workerChan chan engine.Request
}

func (s *SimpleScheduler) WorkerReady(w chan engine.Request) {
}

func (s *SimpleScheduler) WorkerChan() chan engine.Request {
	return s.workerChan
}

func (s *SimpleScheduler) Run() {
	s.workerChan = make(chan engine.Request)
}

func (s *SimpleScheduler) Submit(request engine.Request) {
	// send request down to worker chan 为每一个Request创建goroutine
	go func() {
		s.workerChan <- request
	}()
}
