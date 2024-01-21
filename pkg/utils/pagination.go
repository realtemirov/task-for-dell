package utils

import (
	"fmt"
	"math"
	"strconv"

	"github.com/labstack/echo/v4"
)

const (
	DEFAULT_SIZE int = 10
	MAX_SIZE     int = 50
)

type Query struct {
	Limit  int    `json:"limit,omitempty"`
	Page   int    `json:"page,omitempty"`
	Search string `json:"search,omitempty"`
	Sort   string `json:"sort,omitempty"`
}

// SetLimit
func (q *Query) SetLimit(sizeQuery string) error {
	if sizeQuery == "" {
		q.Limit = DEFAULT_SIZE
		return nil
	}
	n, err := strconv.Atoi(sizeQuery)
	if err != nil {
		return err
	}
	if n > MAX_SIZE || n < 0 {
		n = MAX_SIZE
	}
	q.Limit = n

	return nil
}

// SetPage
func (q *Query) SetPage(pageQuery string) error {
	if pageQuery == "" {
		q.Limit = 0
		return nil
	}
	n, err := strconv.Atoi(pageQuery)
	if err != nil {
		return err
	}
	q.Page = n

	return nil
}

// SetSort
func (q *Query) SetSort(sortQuery string) {
	q.Sort = sortQuery
}

// GetOffset
func (q *Query) GetOffset() int {
	if q.Page == 0 {
		return 0
	}
	return (q.Page - 1) * q.Limit
}

// GetLimit
func (q *Query) GetLimit() int {
	return q.Limit
}

// GetSort
func (q *Query) GetSort() string {
	if q.Sort == "" || q.Sort == "ASC" || q.Sort == "asc" {
		return "ASC"
	} else if q.Sort == "DESC" || q.Sort == "desc" {
		return "DESC"
	} else {
		q.Sort = "DESC"
	}

	return q.Sort
}

// GetPage
func (q *Query) GetPage() int {
	return q.Page
}

func (q *Query) GetQueryString() string {
	return fmt.Sprintf("page=%v&limit=%v&sort=%s", q.GetPage(), q.GetLimit(), q.GetSort())
}

// GetPaginationFromCtx returns the query from the context
func GetPaginationFromCtx(c echo.Context) (*Query, error) {
	q := &Query{}
	if err := q.SetPage(c.QueryParam("page")); err != nil {
		return nil, err
	}
	if err := q.SetLimit(c.QueryParam("limit")); err != nil {
		return nil, err
	}
	q.SetSort(c.QueryParam("sort"))
	q.Search = c.QueryParam("search")

	return q, nil
}

// GetTotalPages calculates the total number of pages using totalCount and pageLimit
func GetTotalPages(totalCount int, pageLimit int) int {
	d := float64(totalCount) / float64(pageLimit)
	return int(math.Ceil(d))
}

// GetHasMore returns whether there is a next page
func GetHasMore(currentPage int, totalCount int, pageLimit int) bool {
	return currentPage < totalCount/pageLimit
}
