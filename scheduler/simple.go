package scheduler

import (
	"example.com/go-http-demo/crawler/engine"
)

type SimpleScheduler struct {
	workerChan chan engine.Request
}

func (s *SimpleScheduler) WorkerChan() chan engine.Request {
	return s.workerChan
}

func (s *SimpleScheduler) WorkerReady(chan engine.Request) {

}

func (s *SimpleScheduler) Run()  {
	s.workerChan = make(chan engine.Request)
}

func (s *SimpleScheduler) Submit (r engine.Request) {
	// 不希望卡死,添加goruntine来添加
	go func() { s.workerChan <- r }()
}