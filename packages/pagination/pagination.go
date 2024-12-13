// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package pagination

import (
	"context"
	"fmt"
	"net/http"
	"strconv"

	"github.com/stainless-sdks/dodo-payments-go/internal/apijson"
	"github.com/stainless-sdks/dodo-payments-go/internal/requestconfig"
)

type PageNumberPage[T any] struct {
	Items []T                `json:"items"`
	JSON  pageNumberPageJSON `json:"-"`
	cfg   *requestconfig.RequestConfig
	res   *http.Response
}

// pageNumberPageJSON contains the JSON metadata for the struct [PageNumberPage[T]]
type pageNumberPageJSON struct {
	Items       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PageNumberPage[T]) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r pageNumberPageJSON) RawJSON() string {
	return r.raw
}

// GetNextPage returns the next page as defined by this pagination style. When
// there is no next page, this function will return a 'nil' for the page value, but
// will not return an error
func (r *PageNumberPage[T]) GetNextPage() (res *PageNumberPage[T], err error) {
	u := r.cfg.Request.URL
	currentPage, err := strconv.Atoi(u.Query().Get("page_number"))
	if err != nil {
		currentPage = 1
	}
	cfg := r.cfg.Clone(context.Background())
	query := cfg.Request.URL.Query()
	query.Set("page_number", fmt.Sprintf("%d", currentPage+1))
	cfg.Request.URL.RawQuery = query.Encode()
	var raw *http.Response
	cfg.ResponseInto = &raw
	cfg.ResponseBodyInto = &res
	err = cfg.Execute()
	if err != nil {
		return nil, err
	}
	res.SetPageConfig(cfg, raw)
	return res, nil
}

func (r *PageNumberPage[T]) SetPageConfig(cfg *requestconfig.RequestConfig, res *http.Response) {
	if r == nil {
		r = &PageNumberPage[T]{}
	}
	r.cfg = cfg
	r.res = res
}

type PageNumberPageAutoPager[T any] struct {
	page *PageNumberPage[T]
	cur  T
	idx  int
	run  int
	err  error
}

func NewPageNumberPageAutoPager[T any](page *PageNumberPage[T], err error) *PageNumberPageAutoPager[T] {
	return &PageNumberPageAutoPager[T]{
		page: page,
		err:  err,
	}
}

func (r *PageNumberPageAutoPager[T]) Next() bool {
	if r.page == nil || len(r.page.Items) == 0 {
		return false
	}
	if r.idx >= len(r.page.Items) {
		r.idx = 0
		r.page, r.err = r.page.GetNextPage()
		if r.err != nil || r.page == nil || len(r.page.Items) == 0 {
			return false
		}
	}
	r.cur = r.page.Items[r.idx]
	r.run += 1
	r.idx += 1
	return true
}

func (r *PageNumberPageAutoPager[T]) Current() T {
	return r.cur
}

func (r *PageNumberPageAutoPager[T]) Err() error {
	return r.err
}

func (r *PageNumberPageAutoPager[T]) Index() int {
	return r.run
}
