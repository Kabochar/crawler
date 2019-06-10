package parser

import (
	"crawler/engine"
	"regexp"
)

var profileRe = regexp.MustCompile(`<a href="(http://album.zhenai.com/u/[0-9]+)"[^>]*>([^<]+)</a>`)

// 用户性别正则，因为在用户详情页没有性别信息，所以在用户性别在用户列表页面获取
var sexRe = regexp.MustCompile(`<td width="180"><span class="grayL">性别：</span>([^<]+)</td>`)

// 底部城市页
var cityUrlsRe = regexp.MustCompile(`href="http://www.zhenai.com/zhenghun/[^"]+]`)

// 城市页面用户解析器
func ParseCity(bytes []byte, _ string) engine.ParseResult {
	cityMatch := profileRe.FindAllSubmatch(bytes, -1)
	gendermatch := sexRe.FindAllSubmatch(bytes, -1)

	result := engine.ParseResult{}

	for k, item := range cityMatch {
		result.Requests = append(result.Requests, engine.Request{
			Url:        string(item[1]),
			ParserFunc: ProfileParser(string(item[2]), string(gendermatch[k][1])),
		})
	}

	// 匹配城市 底部更多
	matches := cityUrlsRe.FindAllSubmatch(bytes, -1)
	for _, m := range matches {
		result.Requests = append(result.Requests,
			engine.Request{
				Url:        string(m[1]),
				ParserFunc: ParseCity,
			})
	}

	return result
}

func ProfileParser(name string, gender string) engine.ParserFunc {
	return func(c []byte, url string) engine.ParseResult {
		return ParseProfile(c, url, name, gender)
	}
}
