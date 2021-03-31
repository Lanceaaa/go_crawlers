package engine

import (
	"log"
)

type ConcurrentEngine struct {
	Scheduler Scheduler
	WorkerCount int
}

type Scheduler interface {
	ReadyNotifier
	Submit(Request)
	WorkerChan() chan Request
	Run()
}

type ReadyNotifier interface {
	WorkerReady(chan Request)
}

func (e *ConcurrentEngine) Run(seeds ...Request) {
	out := make(chan ParseResult)
	e.Scheduler.Run()

	for i := 0; i < e.WorkerCount; i++ {
		// 创建worker
		createWorker(e.Scheduler.WorkerChan(), out, e.Scheduler)
	}

	for _, r := range seeds {
		// url去重
		if isDuplicate(r.Url) {
			continue
		}
		e.Scheduler.Submit(r)
	}

	itemCount := 0
	for {
		// 接收out
		result := <- out
		for _, item := range result.Items {
			if _, ok := item.(model.Profile); ok {
				log.Printf("Got item #%d: %v", itemCount, item)
				itemCount++
			}
		}

		for _, request := range result.Requests {
			// url去重
			if isDuplicate(request.Url) {
				continue
			}
			e.Scheduler.Submit(request)
		}
	}
}

func createWorker(in chan Request, out chan ParseResult, ready ReadyNotifier) {
	go func() {
		for {
			// tell scheduler i'm ready
			ready.WorkerReady(in)
			request := <- in
			parseResult, err := worker(request)
			if err != nil {
				continue
			}
			// 发送
			out <- parseResult
		}
	}()
}

var visitedUrls = make(map[string]bool)
func isDuplicate(url string) bool {
	if visitedUrls[url] {
		return true
	}

	visitedUrls[url] = true
	return false
}
