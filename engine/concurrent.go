package engine

type ConcurrentEngine struct {
	Scheduler Scheduler
	WorkerCount int
	ItemChan chan Item
	RequestProcessor Processor
}

type Processor func (r Request) (ParseResult, error)

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
		e.createWorker(e.Scheduler.WorkerChan(), out, e.Scheduler)
	}

	for _, r := range seeds {
		// url去重
		if isDuplicate(r.Url) {
			continue
		}
		e.Scheduler.Submit(r)
	}

	for {
		// 接收out
		result := <- out
		for _, item := range result.Items {
			go func() { e.ItemChan <- item }()
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

func (e *ConcurrentEngine) createWorker(in chan Request, out chan ParseResult, ready ReadyNotifier) {
	go func() {
		for {
			// tell scheduler i'm ready
			ready.WorkerReady(in)
			request := <- in
			parseResult, err := e.RequestProcessor(request)
			//parseResult, err := worker(request)
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
