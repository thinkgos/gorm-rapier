package rapier

import (
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

var (
	// DefaultPerPage 默认页大小
	DefaultPerPage = int64(50)
	// DefaultPageSize 默认最大页大小
	DefaultMaxPerPage = int64(500)
)

// Paginate 分页器
// 分页索引: page >= 1
// 分页大小: perPage >= 1 && <= DefaultMaxPerPage
func Paginate(page, perPage int64, maxPerPages ...int64) clause.Expression {
	maxPerPage := DefaultMaxPerPage
	if len(maxPerPages) > 0 && maxPerPages[0] > 0 {
		maxPerPage = maxPerPages[0]
	}
	if page < 1 {
		page = 1
	}
	switch {
	case perPage < 1:
		perPage = DefaultPerPage
	case perPage > maxPerPage:
		perPage = maxPerPage
	default: // do nothing
	}
	limit, offset := int(perPage), int(perPage*(page-1))
	return clause.Limit{
		Limit:  &limit,
		Offset: offset,
	}
}

// Pagination 分页器
// 分页索引: page >= 1
// 分页大小: perPage >= 1 && <= DefaultMaxPerPage
func Pagination(page, perPage int64, maxPerPages ...int64) func(*gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		return db.Clauses(Paginate(page, perPage, maxPerPages...))
	}
}

// 限制器
// offset = perPage * (page - 1)
// limit = perPage
// if limit > 0: use limit
// if offset > 0: use offset
// if limit <= 0 and offset <= 0: use none
func Limit(page, perPage int64) func(*gorm.DB) *gorm.DB {
	offset := 0
	if page > 0 {
		offset = int(perPage * (page - 1))
	}
	limit := int(perPage)
	return func(db *gorm.DB) *gorm.DB {
		if offset > 0 || limit > 0 {
			l := clause.Limit{
				Limit:  new(int),
				Offset: offset,
			}
			if offset > 0 {
				l.Offset = offset
			}
			if limit > 0 {
				l.Limit = &limit
			}
			return db.Clauses(l)
		} else {
			return db
		}
	}
}
