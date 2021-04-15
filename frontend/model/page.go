package model

import (
	"example.com/go-http-demo/crawler/engine"
)

type SearchResult struct {
	Hits int64
	Start int
	Items []interface{}
}
