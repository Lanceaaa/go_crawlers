package parser

import (
	"example.com/go-http-demo/crawler/engine"
	"regexp"
)

const cityRe = `<th><a href="(http://album.zhenai.com/u/[0-9]+)"[^>]*>[^<]+</a></th>`

var (
	profileRe = regexp.MustCompile(`<th><a href="(http://album.zhenai.com/u/[0-9]+)"[^>]*>[^<]+</a></th>`)
	cityUrlRe = regexp.MustCompile(`href=http://www.zhenai.com/zhenghun/[^"]+"`)
)

func ParseCity(contents []byte, _ string) engine.ParseResult {
	matches := profileRe.FindAllSubmatch(contents, -1)

	result := engine.ParseReuslt{}
	for _, m := range matches {
		result.Requests = append(result.Requests, engine.Request{
			Url: string(m[1]),
			ParseFunc: NewProfileParser(string(m[2])),
		})
	}

	cityUrlRe.FindAllSubmatch(contents, -1)
	for _, m := range matches {
		result.Requests = append(result.Requests, engine.Request{
			Url: string(m[1]),
			ParseFunc: engine.NewFuncParser(ParseCity, "ParseCity"),
		})
	}

	return result
}
