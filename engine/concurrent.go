package engine

// 并发引擎
type ConcurrendEngine struct {
	Scheduler   Scheduler
	WorkerCount int
	ItemChan    chan Item
}

// 任务调度器
type Scheduler interface {
	ReadyNotifier
	Submit(request Request) // 提交任务
	WorkerChan() chan Request
	Run()
}

type ReadyNotifier interface {
	WorkerReady(w chan Request)
}

func (e *ConcurrendEngine) Run(seeds ...Request) {

	out := make(chan ParseResult)
	e.Scheduler.Run()

	// 创建 goruntine
	for i := 0; i < e.WorkerCount; i++ {
		createWorker(e.Scheduler.WorkerChan(), out, e.Scheduler)
	}

	// engine把请求任务提交给 Scheduler
	for _, request := range seeds {
		if isDuplicated(request.Url) {
			continue
		}
		e.Scheduler.Submit(request)
	}

	for {
		// 接受 Worker 的解析结果
		result := <-out
		for _, item := range result.Items {
			// 使用 goroutine 传递内容
			go func() { e.ItemChan <- item }()
		}

		// 然后把 Worker 解析出的 Request 送给 Scheduler
		for _, request := range result.Requests {
			if isDuplicated(request.Url) {
				continue
			}
			e.Scheduler.Submit(request)
		}
	}
}

func createWorker(in chan Request, out chan ParseResult, ready ReadyNotifier) {
	go func() {
		for {
			ready.WorkerReady(in) // 告诉调度器任务空闲
			request := <-in
			result, err := worker(request)
			if err != nil {
				continue
			}
			out <- result
		}
	}()
}

// 去重判断。仅限当次运行
var visitedUrls = make(map[string]bool)

func isDuplicated(url string) bool {
	if visitedUrls[url] {
		return true
	}

	visitedUrls[url] = true

	return false
}
