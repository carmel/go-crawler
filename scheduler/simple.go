package scheduler

import "go-crawler/engine"

type SimpleScheduler struct {
	workerChan chan engine.Request
}

func (s *SimpleScheduler) Submit(r engine.Request) {
	// 这里有个循环卡死的问题，改用goroutine来做
	// 每个Request会建一个goroutine，这个goroutine往统一的channel(in chan)里分发Request
	go func() {
		s.workerChan <- r
	}()
}

func (s *SimpleScheduler) ConfigureMasterWorkerChan(c chan engine.Request) {
	s.workerChan = c
}
