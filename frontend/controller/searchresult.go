package controller

import (
	"context"
	"example.com/go-http-demo/crawler/frontend/view"
	"github.com/olivere/elastic/v7"
	"net/http"
	"reflect"
	"regexp"
	"strconv"
	"strings"
	"example.com/go-http-demo/crawler/frontend/model"
	"example.com/go-http-demo/crawler/engine"
)

type SearchResultHandle struct {
	view view.SearchResultView
	client *elastic.Client
}

func CreateSearchResultHandler(template string) SearchResultHandle {
	// 创建elastic客户端
	client, err := elastic.NewClient(
		// 必须在docker中关闭嗅探
		elastic.SetSniff(false))
	if err != nil {
		panic(err)
	}

	return SearchResultHandle{
		view:   view.CreateSearchResultView(template),
		client: client,
	}
}

// localhost:8888/search?q=男 已购房&from=20
func (h SearchResultHandle) ServeHTTP(w http.ResponseWriter, req *http.Request)  {
	q := strings.TrimSpace(req.FormValue("q"))
	from, err := strconv.Atoi(req.FormValue("from"))
	if err != nil {
		from = 0
	}

	page, err := h.getSearchResult(q, from)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	err = h.view.Render(w, page)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}
}

func (h SearchResultHandle) getSearchResult(q string, from int) (model.SearchResult, error) {
	var result model.SearchResult
	result.Query = q
	resp, err := h.client.
		Search("dating_profile").
		Query(elastic.NewQueryStringQuery(rewriteQueryString(q))).
		From(from).
		Do(context.Background())
	if err != nil {
		return result, err
	}

    result.Hits = resp.TotalHits()
    result.Start = from
    result.Items = resp.Each(reflect.TypeOf(engine.Item{}))
    result.PrevFrom =  result.Start - len(result.Items)
    result.NextFrom = result.Start + len(result.Items)

    return result, nil
}

func rewriteQueryString(q string) string {
	re := regexp.MustCompile(`[A-Z][a-z]*:`)
	return re.ReplaceAllString(q, "Payload:$1:")
}