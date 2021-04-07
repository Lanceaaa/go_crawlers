package parser

import (
	"example.com/go-http-demo/crawler/engine"
	"fmt"
	"regexp"
)

const cityListRe = `<a href="(http://www.zhenai.com/zhenghun/[0-9a-z]+)"[^>]*>[^<]+</a>`

/**
 * 打印城市列表
 */
func ParseCityList(contents []byte) engine.ParseResult {
	// 正则匹配城市和url
	re := regexp.MustCompile(cityListRe)
	matches := re.FindAllSubmatch(contents, -1)

	result := engine.ParseResult{}
	for _, m := range matches {
		// 把城市对应的名字和url保存起来
		//result.Items = append(result.Items, "City " + string(m[2]))
		result.Requests = append(result.Requests, engine.Request{
			Url: string(m[1]),
			ParseFunc: ParseCity,
		})
	}
	fmt.Printf("Matches found: %d\n", len(matches))
	return result
}
