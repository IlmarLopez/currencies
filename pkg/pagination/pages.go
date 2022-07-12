// Package pagination is a helper for pagination requests and responses.
package pagination

import (
	"net/http"
	"strconv"
)

var (
	DefaultPageSize = 100        // default page size
	MaxPageSize     = 1000       // max page size
	PageVar         = "page"     // page query param
	PageSizeVar     = "per_page" // page size query param
)

// Pages represents pagination information
type Pages struct {
	Page       int         `json:"page"`
	PerPage    int         `json:"per_page"`
	PageCount  int         `json:"page_count"`
	TotalCount int         `json:"total_count"`
	Items      interface{} `json:"items"`
}

// New returns a new Pages object
func New(page, perPage, total int) *Pages {
	if perPage <= 0 {
		perPage = DefaultPageSize
	}
	if perPage > MaxPageSize {
		perPage = MaxPageSize
	}
	pageCount := -1
	if total >= 0 {
		pageCount = (total + perPage - 1) / perPage
		if page > pageCount {
			page = pageCount
		}
	}
	if page < 1 {
		page = 1
	}

	return &Pages{
		Page:       page,
		PerPage:    perPage,
		TotalCount: total,
		PageCount:  pageCount,
	}
}

// NewFromRequest returns a new Pages object from a request
func NewFromRequest(req *http.Request, count int) *Pages {
	page := parseInt(req.URL.Query().Get(PageVar), 1)
	perPage := parseInt(req.URL.Query().Get(PageSizeVar), DefaultPageSize)
	return New(page, perPage, count)
}

// parseInt parses an integer from a string.
func parseInt(value string, defaultValue int) int {
	if value == "" {
		return defaultValue
	}
	if result, err := strconv.Atoi(value); err == nil {
		return result
	}
	return defaultValue
}

// Offset returns the offset of the current page
func (p *Pages) Offset() int {
	return (p.Page - 1) * p.PerPage
}

// Limit returns the limit of the current page
func (p *Pages) Limit() int {
	return p.PerPage
}
