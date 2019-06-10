package engine

// 转换函数
type ParserFunc func(contents []byte, url string) ParseResult

// 请求结构
type Request struct {
	Url        string     // 请求地址
	ParserFunc ParserFunc // 解析函数
}

// 解析结果结构
type ParseResult struct {
	Requests []Request // 解析出的请求
	Items    []Item    // 解析出的内容
}

type Item struct {
	Url     string
	Type    string
	Id      string
	Payload interface{}
}

// 定义一个真实的函数，什么也不做
func NilParse([]byte) ParseResult {
	return ParseResult{}
}
