package model

type SearchResult struct {
	Hits  int64 // 结果数量
	Start int   // 起始位置
	Items []interface{}
	//Items []engine.Item
	Query    string // 请求
	PrevFrom int    // 上一页
	NextFrom int    // 下一页
}
