package engine

type ParserFunc interface {
	Parse(contents []byte, url string) ParseResult
	Serialize() (name string, args interface{})
}

type Request struct {
	Url string
	Parser Parser
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

type NilParser struct {
	
}

type FuncParser struct {
	parser ParserFunc
	name string
}

func (f *FuncParser) Parse(contents []byte, url string) ParseResult {
	f.Parse(contents, url)
}

func (f *FuncParser) Serialize() (name string, args interface{}) {
	return f.name, nil
}

func NewFuncParser(p ParserFunc, name string) *FuncParser {
	return &FuncParser{
		parser: p,
		name:   name,
	}
}