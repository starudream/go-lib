package sqlite

import (
	"regexp"
	"strings"
)

type Pagination struct {
	Page     int `json:"current"`
	PageSize int `json:"pageSize"`

	TotalSize int `json:"total"`
	TotalPage int `json:"totalPage"`
}

const defaultPage, defaultPageSize int = 1, 20

func (p *Pagination) Limit() int {
	if p == nil {
		return defaultPageSize
	}
	if p.PageSize <= 0 {
		p.PageSize = defaultPageSize
	}
	return p.PageSize
}

func (p *Pagination) Offset() int {
	if p == nil {
		return 0
	}
	if p.Page <= 0 {
		p.Page = defaultPage
	}
	return (p.Page - 1) * p.Limit()
}

func (p *Pagination) Total(size int64) *Pagination {
	np := p
	if p == nil {
		np = &Pagination{Page: defaultPage, PageSize: defaultPageSize}
	}
	np.TotalSize = int(size)
	np.TotalPage = (np.TotalSize + p.Limit() - 1) / p.Limit()
	p.Offset() // init page
	return np
}

type ListParams struct {
	Pagination *Pagination `json:"pagination"`
	Sort       []string    `json:"sort,omitempty"`
}

func (ps *ListParams) GetPagination() *Pagination {
	if ps == nil {
		return nil
	}
	return ps.Pagination
}

var orderKeyRegex = regexp.MustCompile("^[+-]?[a-zA-Z_][a-zA-Z0-9_]*$")

func (ps *ListParams) Order() (ss []string) {
	if ps == nil || len(ps.Sort) == 0 {
		return
	}
	for i := 0; i < len(ps.Sort); i++ {
		field := ps.Sort[i]
		if !orderKeyRegex.MatchString(field) {
			continue
		}
		desc := false
		switch {
		case strings.HasPrefix(field, "-"):
			desc = true
			fallthrough
		case strings.HasPrefix(field, "+"):
			field = field[1:]
		}
		if desc {
			field += " DESC"
		}
		ss = append(ss, field)
	}
	return
}

type ListResult[T any] struct {
	Pagination *Pagination `json:"pagination"`
	Records    []T         `json:"records"`
}

func QueryList[T any](tx *GormDB, ps *ListParams) (_ *ListResult[T], err error) {
	var total int64
	err = tx.Count(&total).Error
	if err != nil {
		return nil, err
	}

	pagination := ps.GetPagination()
	if total == 0 {
		return &ListResult[T]{Pagination: pagination.Total(total)}, nil
	}

	var records []T
	err = tx.Order(strings.Join(ps.Order(), ",")).Offset(pagination.Offset()).Limit(pagination.Limit()).Find(&records).Error
	if err != nil {
		return nil, err
	}

	return &ListResult[T]{Pagination: pagination.Total(total), Records: records}, nil
}
