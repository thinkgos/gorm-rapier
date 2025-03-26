package rapier

import (
	"fmt"
	"strings"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

func SubQuery(db *gorm.DB) Field {
	return Field{
		expr{
			e: clause.Expr{
				SQL:  "(?)",
				Vars: []any{db},
			},
		},
	}
}

// Exist equivalent EXISTS(subQuery)
func Exist(subQuery *gorm.DB) Expr {
	return expr{e: clause.Expr{
		SQL:  "EXISTS(?)",
		Vars: []any{subQuery},
	}}
}

// NotExist equivalent NOT EXISTS(subQuery)
func NotExist(subQuery *gorm.DB) Expr {
	return expr{e: clause.Expr{
		SQL:  "NOT EXISTS(?)",
		Vars: []any{subQuery},
	}}
}

// Columns columns array
type Columns []Expr

// NewColumns new columns instance.
func NewColumns(cols ...Expr) Columns { return cols }

// SetSubQuery set with subQuery
func (cs Columns) Set(subQuery *gorm.DB) AssignExpr {
	if len(cs) == 0 {
		return expr{e: clause.Set{}}
	}
	cols := make([]string, len(cs))
	for i, v := range cs {
		cols[i] = v.BuildColumn(subQuery.Statement)
	}
	name := cols[0]
	if len(cols) > 1 {
		name = "(" + strings.Join(cols, ",") + ")"
	}
	return expr{
		e: clause.Set{
			{
				Column: clause.Column{
					Name: name,
					Raw:  true,
				},
				Value: gorm.Expr("(?)", subQuery),
			},
		},
	}
}

// IN return contains subQuery or value
// when len(columns) == 1, equal to columns[0] IN (subQuery/value)
// when len(columns) > 1, equal to (columns[0], columns[1], ...) IN (subQuery/value)
func (cs Columns) In(subQueryOrValue any) Expr {
	switch v := subQueryOrValue.(type) {
	case *gorm.DB:
		return cs.containsSubQuery("IN", v)
	case nil:
		return EmptyExpr()
	default:
		return cs.containsValues("IN", clause.Expr{
			SQL:                "?",
			Vars:               []any{subQueryOrValue},
			WithoutParentheses: true,
		})
	}
}

// IN return contains subQuery or value
// when len(columns) == 1, equal to NOT columns[0] IN (subQuery/value)
// when len(columns) > 1, equal to NOT (columns[0], columns[1], ...) IN (subQuery/value)
func (cs Columns) NotIn(subQueryOrValue any) Expr {
	switch v := subQueryOrValue.(type) {
	case *gorm.DB:
		return cs.containsSubQuery("NOT IN", v)
	case nil:
		return EmptyExpr()
	default:
		return cs.containsValues("NOT IN", clause.Expr{
			SQL:                "?",
			Vars:               []any{subQueryOrValue},
			WithoutParentheses: true,
		})
	}
}

func (columns Columns) containsSubQuery(op string, subQuery *gorm.DB) Expr {
	switch len(columns) {
	case 0:
		return EmptyExpr()
	case 1:
		return expr{e: clause.Expr{
			SQL:  fmt.Sprintf("? %s (?)", op),
			Vars: []any{columns[0].RawExpr(), subQuery},
		}}
	default: // len(columns) > 0
		placeholders := make([]string, len(columns))
		cols := make([]any, len(columns))
		for i, c := range columns {
			placeholders[i], cols[i] = "?", c.RawExpr()
		}
		return expr{e: clause.Expr{
			SQL:  fmt.Sprintf("(%s) %s (?)", strings.Join(placeholders, ","), op),
			Vars: append(cols, subQuery),
		}}
	}
}

func (columns Columns) containsValues(op string, value clause.Expr) Expr {
	switch len(columns) {
	case 0:
		return EmptyExpr()
	case 1:
		return expr{e: clause.Expr{
			SQL:  fmt.Sprintf("? %s (?)", op),
			Vars: []any{columns[0].RawExpr(), value},
		}}
	default: // len(columns) > 0
		vars := make([]string, len(columns))
		queryCols := make([]any, len(columns))
		for i, c := range columns {
			vars[i], queryCols[i] = "?", c.RawExpr()
		}
		return expr{e: clause.Expr{
			SQL:  fmt.Sprintf("(%s) %s (?)", strings.Join(vars, ", "), op),
			Vars: append(queryCols, value),
		}}
	}
}
