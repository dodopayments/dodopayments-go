// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package pagination

import (
	"context"
	"fmt"
	"net/http"
	"strconv"

	"github.com/dodopayments/dodopayments-go/internal/apijson"
	"github.com/dodopayments/dodopayments-go/internal/requestconfig"
	"github.com/dodopayments/dodopayments-go/option"
)

type DefaultPageNumberPagination[T any] struct {
	Items []T                             `json:"items"`
	JSON  defaultPageNumberPaginationJSON `json:"-"`
	cfg   *requestconfig.RequestConfig
	res   *http.Response
}

// defaultPageNumberPaginationJSON contains the JSON metadata for the struct
// [DefaultPageNumberPagination[T]]
type defaultPageNumberPaginationJSON struct {
	Items       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DefaultPageNumberPagination[T]) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r defaultPageNumberPaginationJSON) RawJSON() string {
	return r.raw
}

// GetNextPage returns the next page as defined by this pagination style. When
// there is no next page, this function will return a 'nil' for the page value, but
// will not return an error
func (r *DefaultPageNumberPagination[T]) GetNextPage() (res *DefaultPageNumberPagination[T], err error) {
	if len(r.Items) == 0 {
		return nil, nil
	}
	u := r.cfg.Request.URL
	currentPage, err := strconv.ParseInt(u.Query().Get("page_number"), 10, 64)
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

func (r *DefaultPageNumberPagination[T]) SetPageConfig(cfg *requestconfig.RequestConfig, res *http.Response) {
	if r == nil {
		r = &DefaultPageNumberPagination[T]{}
	}
	r.cfg = cfg
	r.res = res
}

type DefaultPageNumberPaginationAutoPager[T any] struct {
	page *DefaultPageNumberPagination[T]
	cur  T
	idx  int
	run  int
	err  error
}

func NewDefaultPageNumberPaginationAutoPager[T any](page *DefaultPageNumberPagination[T], err error) *DefaultPageNumberPaginationAutoPager[T] {
	return &DefaultPageNumberPaginationAutoPager[T]{
		page: page,
		err:  err,
	}
}

func (r *DefaultPageNumberPaginationAutoPager[T]) Next() bool {
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

func (r *DefaultPageNumberPaginationAutoPager[T]) Current() T {
	return r.cur
}

func (r *DefaultPageNumberPaginationAutoPager[T]) Err() error {
	return r.err
}

func (r *DefaultPageNumberPaginationAutoPager[T]) Index() int {
	return r.run
}

type CursorPagePagination[T any] struct {
	Data     []T                      `json:"data"`
	Iterator string                   `json:"iterator"`
	Done     bool                     `json:"done"`
	JSON     cursorPagePaginationJSON `json:"-"`
	cfg      *requestconfig.RequestConfig
	res      *http.Response
}

// cursorPagePaginationJSON contains the JSON metadata for the struct
// [CursorPagePagination[T]]
type cursorPagePaginationJSON struct {
	Data        apijson.Field
	Iterator    apijson.Field
	Done        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CursorPagePagination[T]) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r cursorPagePaginationJSON) RawJSON() string {
	return r.raw
}

// GetNextPage returns the next page as defined by this pagination style. When
// there is no next page, this function will return a 'nil' for the page value, but
// will not return an error
func (r *CursorPagePagination[T]) GetNextPage() (res *CursorPagePagination[T], err error) {
	if len(r.Data) == 0 {
		return nil, nil
	}

	if !r.JSON.Done.IsMissing() && r.Done == false {
		return nil, nil
	}
	next := r.Iterator
	if len(next) == 0 {
		return nil, nil
	}
	cfg := r.cfg.Clone(r.cfg.Context)
	err = cfg.Apply(option.WithQuery("iterator", next))
	if err != nil {
		return nil, err
	}
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

func (r *CursorPagePagination[T]) SetPageConfig(cfg *requestconfig.RequestConfig, res *http.Response) {
	if r == nil {
		r = &CursorPagePagination[T]{}
	}
	r.cfg = cfg
	r.res = res
}

type CursorPagePaginationAutoPager[T any] struct {
	page *CursorPagePagination[T]
	cur  T
	idx  int
	run  int
	err  error
}

func NewCursorPagePaginationAutoPager[T any](page *CursorPagePagination[T], err error) *CursorPagePaginationAutoPager[T] {
	return &CursorPagePaginationAutoPager[T]{
		page: page,
		err:  err,
	}
}

func (r *CursorPagePaginationAutoPager[T]) Next() bool {
	if r.page == nil || len(r.page.Data) == 0 {
		return false
	}
	if r.idx >= len(r.page.Data) {
		r.idx = 0
		r.page, r.err = r.page.GetNextPage()
		if r.err != nil || r.page == nil || len(r.page.Data) == 0 {
			return false
		}
	}
	r.cur = r.page.Data[r.idx]
	r.run += 1
	r.idx += 1
	return true
}

func (r *CursorPagePaginationAutoPager[T]) Current() T {
	return r.cur
}

func (r *CursorPagePaginationAutoPager[T]) Err() error {
	return r.err
}

func (r *CursorPagePaginationAutoPager[T]) Index() int {
	return r.run
}
