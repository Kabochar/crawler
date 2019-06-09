package main

import (
	"crawler/engine"
	"crawler/scheduler"
	"crawler/zhenai/parser"
)

func main() {

	e := engine.ConcurrendEngine{
		Scheduler:   &scheduler.QueuedScheduler{}, // 这里调用并发调度器
		WorkerCount: 100,
	}
	//e.Run(engine.Request{
	//	Url:       "http://www.zhenai.com/zhenghun/",
	//	ParseFunc: parser.ParseCityList,
	//})

	// 捕捉底部更多城市链接
	e.Run(engine.Request{
		Url:       "http://www.zhenai.com/zhenghun/shanghai",
		ParseFunc: parser.ParseCityList,
	})
}
