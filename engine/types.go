package engine

type ParseFunc func(contents []byte, url string) ParseResult

type Request struct {
	Url string
	ParseFunc ParseFunc
}

type ParseResult struct {
	Requests []Request
	Items []Item
}

type Item struct {
	Url string // 链接
	Type string // 类型
	Id string // 主键
	Payload interface{}
}

func NilParser([]byte) ParseResult {
	return ParseResult{}
}
