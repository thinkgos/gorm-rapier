package assist

import "gorm.io/gorm/clause"

// Option field option
type Option func(clause.Column) clause.Column

var columnRaw Option = func(col clause.Column) clause.Column {
	col.Raw = false
	return col
}

func intoClauseColumn(table, column string, opts ...Option) clause.Column {
	col := clause.Column{Table: table, Name: column}
	for _, opt := range opts {
		col = opt(col)
	}
	return columnRaw(col)
}
