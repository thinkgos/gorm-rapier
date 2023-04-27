package assist

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
func Pagination(page, perPage int64, maxPerPages ...int64) Condition {
	return func(db *gorm.DB) *gorm.DB {
		return db.Clauses(Paginate(page, perPage, maxPerPages...))
	}
}
