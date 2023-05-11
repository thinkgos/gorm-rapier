package assist

import (
	"fmt"

	"gorm.io/gorm/clause"
)

// String string type field
type String Field

// NewString new string field.
func NewString(table, column string, opts ...Option) String {
	return String{
		expr: expr{
			col: intoClauseColumn(table, column, opts...),
		},
	}
}

// IfNull use IFNULL(expr,?)
func (field String) IfNull(value string) Expr {
	return field.ifNull(value)
}

// Eq equal to, use expr = ?
func (field String) Eq(value string) Expr {
	return expr{
		col:       field.col,
		e:         clause.Eq{Column: field.RawExpr(), Value: value},
		buildOpts: field.buildOpts,
	}
}

// Neq not equal to, use expr <> ?
func (field String) Neq(value string) Expr {
	return expr{
		col:       field.col,
		e:         clause.Neq{Column: field.RawExpr(), Value: value},
		buildOpts: field.buildOpts,
	}
}

// Gt greater than, use expr > ?
func (field String) Gt(value string) Expr {
	return expr{
		col:       field.col,
		e:         clause.Gt{Column: field.RawExpr(), Value: value},
		buildOpts: field.buildOpts,
	}
}

// Gte greater or equal to, use expr >= ?
func (field String) Gte(value string) Expr {
	return expr{
		col:       field.col,
		e:         clause.Gte{Column: field.RawExpr(), Value: value},
		buildOpts: field.buildOpts,
	}
}

// Lt less than, use expr < ?
func (field String) Lt(value string) Expr {
	return expr{
		col:       field.col,
		e:         clause.Lt{Column: field.RawExpr(), Value: value},
		buildOpts: field.buildOpts,
	}
}

// Lte less or equal to, use expr <= ?
func (field String) Lte(value string) Expr {
	return expr{
		col:       field.col,
		e:         clause.Lte{Column: field.RawExpr(), Value: value},
		buildOpts: field.buildOpts,
	}
}

// Between use expr BETWEEN ? AND ?
func (field String) Between(left, right string) Expr {
	return field.between([]any{left, right})
}

// NotBetween use NOT (expr BETWEEN ? AND ?)
func (field String) NotBetween(left, right string) Expr {
	return field.notBetween([]any{left, right})
}

// In use expr IN (?)
func (field String) In(values ...string) Expr {
	return expr{
		col:       field.col,
		e:         clause.IN{Column: field.RawExpr(), Values: intoAnySlice(values...)},
		buildOpts: field.buildOpts,
	}
}

// NotIn use expr NOT IN (?)
func (field String) NotIn(values ...string) Expr {
	return expr{
		col:       field.col,
		e:         clause.Not(clause.IN{Column: field.RawExpr(), Values: intoAnySlice(values...)}),
		buildOpts: field.buildOpts,
	}
}

// Like use expr LIKE ?
func (field String) Like(value string) Expr {
	return expr{
		col:       field.col,
		e:         clause.Like{Column: field.RawExpr(), Value: value},
		buildOpts: field.buildOpts,
	}
}

// FuzzyLike use expr LIKE ?, ? contain prefix % and suffix %
// e.g. expr LIKE %value%
func (field String) FuzzyLike(value string) Expr {
	return expr{
		col:       field.col,
		e:         clause.Like{Column: field.RawExpr(), Value: "%" + value + "%"},
		buildOpts: field.buildOpts,
	}
}

// LeftLike use expr LIKE ?, ? contain suffix %.
// e.g. expr LIKE value%
func (field String) LeftLike(value string) Expr {
	return expr{
		col:       field.col,
		e:         clause.Like{Column: field.RawExpr(), Value: value + "%"},
		buildOpts: field.buildOpts,
	}
}

// NotLike use expr NOT LIKE ?
func (field String) NotLike(value string) Expr {
	return expr{
		col:       field.col,
		e:         clause.Not(clause.Like{Column: field.RawExpr(), Value: value}),
		buildOpts: field.buildOpts,
	}
}

// Regexp use expr REGEXP ?
func (field String) Regexp(value string) Expr {
	return field.regexp(value)
}

// NotRegxp use NOT expr REGEXP ?
func (field String) NotRegxp(value string) Expr {
	return field.notRegexp(value)
}

// FindInSet equal to FIND_IN_SET(field_name, input_string_list)
func (field String) FindInSet(targetList string) Expr {
	return expr{
		col:       field.col,
		e:         clause.Expr{SQL: "FIND_IN_SET(?,?)", Vars: []any{field.RawExpr(), targetList}},
		buildOpts: field.buildOpts,
	}
}

// FindInSetWith equal to FIND_IN_SET(input_string, field_name)
func (field String) FindInSetWith(target string) Expr {
	return expr{
		col:       field.col,
		e:         clause.Expr{SQL: "FIND_IN_SET(?,?)", Vars: []any{target, field.RawExpr()}},
		buildOpts: field.buildOpts,
	}
}

// SubstringIndex use SUBSTRING_INDEX(expr,?,?)
// https://dev.mysql.com/doc/refman/8.0/en/functions.html#function_substring-index
func (field String) SubstringIndex(delim string, count int) String {
	return String{
		expr{
			col: field.col,
			e: clause.Expr{
				SQL:  fmt.Sprintf("SUBSTRING_INDEX(?,%q,%d)", delim, count),
				Vars: []any{field.RawExpr()},
			},
			buildOpts: field.buildOpts,
		},
	}
}

// Replace use REPLACE(expr,?,?)
func (field String) Replace(from, to string) String {
	return String{
		expr{
			col:       field.col,
			e:         clause.Expr{SQL: "REPLACE(?,?,?)", Vars: []any{field.RawExpr(), from, to}},
			buildOpts: field.buildOpts,
		},
	}
}

// Concat use CONCAT(?,?,?)
func (field String) Concat(before, after string) String {
	var e clause.Expression

	switch {
	case before != "" && after != "":
		e = &clause.Expr{SQL: "CONCAT(?,?,?)", Vars: []any{before, field.RawExpr(), after}}
	case before != "":
		e = &clause.Expr{SQL: "CONCAT(?,?)", Vars: []any{before, field.RawExpr()}}
	case after != "":
		e = &clause.Expr{SQL: "CONCAT(?,?)", Vars: []any{field.RawExpr(), after}}
	default:
		return field
	}
	return String{
		expr{
			col:       field.col,
			e:         e,
			buildOpts: field.buildOpts,
		},
	}
}

// Trim use TRIM(BOTH ? FROM ?)
func (field String) Trim(remStr string) String {
	return String{
		expr{
			col:       field.col,
			e:         clause.Expr{SQL: "TRIM(BOTH ? FROM ?)", Vars: []any{remStr, field.RawExpr()}},
			buildOpts: field.buildOpts,
		},
	}
}

// LTrim use TRIM(LEADING ? FROM ?)
func (field String) LTrim(remStr string) String {
	return String{
		expr{
			col:       field.col,
			e:         clause.Expr{SQL: "TRIM(LEADING ? FROM ?)", Vars: []any{remStr, field.RawExpr()}},
			buildOpts: field.buildOpts,
		},
	}
}

// RTrim use TRIM(TRAILING ? FROM ?)
func (field String) RTrim(remStr string) String {
	return String{
		expr{
			col:       field.col,
			e:         clause.Expr{SQL: "TRIM(TRAILING ? FROM ?)", Vars: []any{remStr, field.RawExpr()}},
			buildOpts: field.buildOpts,
		},
	}
}

// TrimSpace use TRIM(?)
func (field String) TrimSpace() String {
	return String{
		expr{
			col:       field.col,
			e:         clause.Expr{SQL: "TRIM(?)", Vars: []any{field.RawExpr()}},
			buildOpts: field.buildOpts,
		},
	}
}

// LTrimSpace use LTRIM(?)
func (field String) LTrimSpace() String {
	return String{
		expr{
			col:       field.col,
			e:         clause.Expr{SQL: "LTRIM(?)", Vars: []any{field.RawExpr()}},
			buildOpts: field.buildOpts,
		},
	}
}

// RTrimSpace use RTRIM(?)
func (field String) RTrimSpace() String {
	return String{
		expr{
			col:       field.col,
			e:         clause.Expr{SQL: "RTRIM(?)", Vars: []any{field.RawExpr()}},
			buildOpts: field.buildOpts,
		},
	}
}

// IntoColumns columns array with sub method
func (field String) IntoColumns() Columns {
	return NewColumns(field)
}
