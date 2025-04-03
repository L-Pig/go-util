package github.com/l-pig/go-util

import (
	"errors"
	"gorm.io/gorm"
	"math"
)

type PageRequest struct {
	PageIndex int `json:"page_index,default=1" form:"page_index,,default=1" query:"page_index,default=1"`
	PageSize  int `json:"page_size,default=10" form:"page_size,default=10" query:"page_size,default=10"`
}

type PageResponse[T any] struct {
	PageIndex int   `json:"pageIndex"`
	PageSize  int   `json:"pageSize"`
	TotalPage int   `json:"totalPage"`
	Total     int64 `json:"total"`
	List      []T   `json:"data"`
}

type PageModel[T any] struct {
	limit     int
	offset    int
	pageIndex int
	pageSize  int
	totalPage int
	query     *gorm.DB
	scopes    []func(*gorm.DB) *gorm.DB
	total     int64
	data      []T
	Error     error
}

func StartPage[T any](pageIndex, pageSize int, query *gorm.DB) *PageModel[T] {
	pageModel := PageModel[T]{
		pageIndex: pageIndex,
		pageSize:  pageSize,
		query:     query,
		limit:     pageSize,
		offset:    (pageIndex - 1) * pageSize,
	}

	var count int64

	pageModel.query.Count(&count)

	pageModel.total = count
	pageModel.totalPage = calcPage(count, pageSize)

	pageModel.list()
	return &pageModel
}

// calcPage 计算总页数
// 使用向上取整，避免最后一页不足 pageSize 时无法显示
func calcPage(total int64, pageSize int) int {
	return int(math.Ceil(float64(total) / float64(pageSize)))
}

// list 查询
func (p *PageModel[T]) list() {
	if p.query == nil {
		p.Error = errors.New("query is nil")
		return
	}
	p.Error = p.query.Limit(p.limit).Offset(p.offset).Find(&p.data).Error
}

func (p *PageModel[T]) Result() PageResponse[T] {
	return PageResponse[T]{
		PageIndex: p.pageIndex,
		PageSize:  p.pageSize,
		TotalPage: p.totalPage,
		Total:     p.total,
		List:      p.data,
	}
}
