package scheduler

import (
	"example.com/go-http-demo/crawler/engine"
)

type QueuedScheduler struct {
	requestChan chan engine.Request
	workerChan chan chan engine.Request
}


func (s *QueuedScheduler) WorkerChan() chan engine.Request {
	return make(chan engine.Request)
}

func (s *QueuedScheduler) Submit(r engine.Request) {
	s.requestChan <- r
}

func (s *QueuedScheduler) WorerReady(w chan engine.Request) {
	s.workerChan <- w
}

func (s *QueuedScheduler) ConfigureMasterWorkerChan(c chan engine.Request) {

}

func (s *QueuedScheduler) Run() {
	s.workerChan = make(chan chan engine.Request)
	s.requestChan = make(chan engine.Request)
	go func() {
		var requestQ []engine.Request
		var workerQ []chan engine.Request
		for {
			var activeRequest engine.Request
			var activeWorker chan engine.Request
			// 如果两个都有值则
			if len(requestQ) > 0 && len(workerQ) > 0 {
				activeWorker = workerQ[0]
				activeRequest = requestQ[0]

			}
			// 通过使用select来同时获取r、w，有可能r先返回，有可能w先返回
			select {
			case r := <- s.requestChan:
				// send r to a worker
				requestQ = append(requestQ, r)
			case w := <- s.workerChan:
				// send next_request to w
				workerQ = append(workerQ, w)
			case activeWorker <- activeRequest:
				workerQ = workerQ[1:]
				requestQ = requestQ[1:]
			}
		}
	}()
}