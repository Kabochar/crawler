package parser

import (
	"crawler/engine"
	"regexp"
)

const cityListRe = `<a href="(http://www.zhenai.com/zhenghun/[0-9a-z]+)"[^>]*>([^<]+)</a>`

// 解析城市列表
func ParseCityList(contents []byte, _ string) engine.ParseResult {
	re := regexp.MustCompile(cityListRe)
	submatch := re.FindAllSubmatch(contents, -1)
	result := engine.ParseResult{}
	limit := 470 // 限制搜索次数
	for _, item := range submatch {
		result.Requests = append(result.Requests, engine.Request{
			Url:        string(item[1]),
			ParserFunc: ParseCity,
		})
		if limit--; limit == 0 {
			break
		} // 控制次数
	}
	return result
}
