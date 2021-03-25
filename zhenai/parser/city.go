package parser

import (
	"example.com/go-http-demo/crawler/engine"
	"regexp"
)

const cityRe = `<th><a href="(http://album.zhenai.com/u/[0-9]+)"[^>]*>[^<]+</a></th>`

func ParseCity(contents []byte) engine.ParseResult {
	re := regexp.MustCompile(cityRe)
	matches := re.FindAllSubmatch(contents, -1)

	result := engine.ParseReuslt{}
	for _, m := range matches {
		name := string(m[2])
		result.Items = append(result.Items, "user " + name)
		result.Requests = append(result.Requests, engine.Request{
			Url: string(m[1]),
			ParseFunc: func(c []byte) engine.ParseResult {
				return ParseProfile(c, name)
			},
		})
	}
	return result
}
