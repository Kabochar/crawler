package parser

import (
	"crawler/engine"
	"crawler/model"
	"regexp"
	"strconv"
)

var ageRe = regexp.MustCompile(`<div class="m-btn purple" [^>]*>([\d]+)岁</div>`)
var heightRe = regexp.MustCompile(`<div class="m-btn purple" [^>]*>([\d]+)cm</div>`)
var weightRe = regexp.MustCompile(`<div class="m-btn purple" [^>]*>([\d]+)kg</div>`)

var incomeRe = regexp.MustCompile(`<div class="m-btn purple" [^>]*>月收入:([^<]+)</div>`)
var marriageRe = regexp.MustCompile(`<div class="m-btn purple" [^>]*>([^<]+)</div>`)
var addressRe = regexp.MustCompile(`<div class="m-btn purple" [^>]*>工作地:([^<]+)</div>`)
var xinzuoRe = regexp.MustCompile(`<div class="m-btn purple" [^>]*>.*?座([^<]+)</div>`)
var idUrlRe = regexp.MustCompile(`http://album.zhenai.com/u/([\d]+)`)

func ParseProfile(bytes []byte, url string, name string, gender string) engine.ParseResult {
	profile := model.Profile{}
	profile.Name = name
	profile.Gender = gender
	if age, err := strconv.Atoi(extractString(bytes, ageRe)); err == nil {
		profile.Age = age
	}
	if height, err := strconv.Atoi(extractString(bytes, heightRe)); err == nil {
		profile.Height = height
	}
	if weight, err := strconv.Atoi(extractString(bytes, weightRe)); err == nil {
		profile.Weight = weight
	}

	profile.Income = extractString(bytes, incomeRe)
	profile.Marriage = extractString(bytes, marriageRe)
	profile.Address = extractString(bytes, addressRe)
	profile.Xinzuo = extractString(bytes, xinzuoRe)

	// 解析完用户信息后，没有请求任务
	result := engine.ParseResult{
		Items: []engine.Item{
			{
				Url:     url,
				Type:    "zhenai",
				Id:      extractString([]byte(url), idUrlRe),
				Payload: profile,
			},
		},
	}
	return result
}

func extractString(contents []byte, re *regexp.Regexp) string {
	submatch := re.FindSubmatch(contents)
	if len(submatch) >= 2 {
		return string(submatch[1])
	} else {
		return ""
	}
}
