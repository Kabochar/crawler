package parser

import (
	"crawler/engine"
	"regexp"
)

const cityListRe = `<a href="(http://www.zhenai.com/zhenghun/[0-9a-z]+)"[^>]*>([^<]+)</a>`

// 解析城市列表
func ParseCityList(contents []byte) engine.ParseResult {
	re := regexp.MustCompile(cityListRe)
	submatch := re.FindAllSubmatch(contents, -1)
	result := engine.ParseResult{}
	limit := 10 // 限制搜索次数
	for _, item := range submatch {
		result.Items = append(result.Items, string(item[2]))
		result.Requests = append(result.Requests, engine.Request{
			Url:       string(item[1]),
			ParseFunc: ParseCity,
		})
		if limit--; limit == 0 {
			break
		} // 控制次数
	}
	return result
}
