package assist

import (
	"fmt"
	"strings"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

func containsSubQuery(columns []Expr, subQuery *gorm.DB) Expr {
	switch len(columns) {
	case 0:
		return EmptyExpr()
	case 1:
		return expr{e: clause.Expr{
			SQL:  "? IN (?)",
			Vars: []any{columns[0].RawExpr(), subQuery},
		}}
	default: // len(columns) > 0
		placeholders := make([]string, len(columns))
		cols := make([]any, len(columns))
		for i, c := range columns {
			placeholders[i], cols[i] = "?", c.RawExpr()
		}
		return expr{e: clause.Expr{
			SQL:  fmt.Sprintf("(%s) IN (?)", strings.Join(placeholders, ",")),
			Vars: append(cols, subQuery),
		}}
	}
}

func containsValues(columns []Expr, value Value) Expr {
	switch len(columns) {
	case 0:
		return EmptyExpr()
	case 1:
		return expr{e: clause.Expr{
			SQL:  "? IN (?)",
			Vars: []any{columns[0].RawExpr(), clause.Expr(value)},
		}}
	default: // len(columns) > 0
		vars := make([]string, len(columns))
		queryCols := make([]any, len(columns))
		for i, c := range columns {
			vars[i], queryCols[i] = "?", c.RawExpr()
		}
		return expr{e: clause.Expr{
			SQL:  fmt.Sprintf("(%s) IN (?)", strings.Join(vars, ", ")),
			Vars: append(queryCols, clause.Expr(value)),
		}}
	}
}

// compareOperator compare operator
type compareOperator string

const (
	// eqOp =
	eqOp compareOperator = " = "
	// neqOp <>
	neqOp compareOperator = " <> "
	// gtOp >
	gtOp compareOperator = " > "
	// gteOp >=
	gteOp compareOperator = " >= "
	// ltOp <
	ltOp compareOperator = " < "
	// lteOp <=
	lteOp compareOperator = " <= "
)

// compareSubQuery compare with sub query
func compareSubQuery(op compareOperator, column Expr, subQuery *gorm.DB) Expr {
	return expr{e: clause.Expr{
		SQL:  fmt.Sprint("?", op, "(?)"),
		Vars: []any{column.RawExpr(), subQuery},
	}}
}

// Columns columns array
type Columns []Expr

// NewColumns new columns instance.
func NewColumns(cols ...Expr) Columns { return cols }

// IN return contains subQuery or value
// when len(columns) == 1, equal to columns[0] IN (subQuery/value)
// when len(columns) > 1, equal to (columns[0], columns[1], ...) IN (subQuery/value)
func (cs Columns) In(subQueryOrValue any) Expr {
	switch v := subQueryOrValue.(type) {
	case Value:
		return containsValues(cs, v)
	case *gorm.DB:
		return containsSubQuery(cs, v)
	default:
		return EmptyExpr()
	}
}

// IN return contains subQuery or value
// when len(columns) == 1, equal to NOT columns[0] IN (subQuery/value)
// when len(columns) > 1, equal to NOT (columns[0], columns[1], ...) IN (subQuery/value)
func (cs Columns) NotIn(subQueryOrValue any) Expr {
	return Not(cs.In(subQueryOrValue))
}

// Eq  equivalent column = (subQuery)
func (cs Columns) Eq(subQuery *gorm.DB) Expr {
	if len(cs) == 0 {
		return EmptyExpr()
	}
	return compareSubQuery(eqOp, cs[0], subQuery)
}

// Neq equivalent column <> (subQuery)
func (cs Columns) Neq(subQuery *gorm.DB) Expr {
	if len(cs) == 0 {
		return EmptyExpr()
	}
	return compareSubQuery(neqOp, cs[0], subQuery)
}

// Gt  equivalent column > (subQuery)
func (cs Columns) Gt(subQuery *gorm.DB) Expr {
	if len(cs) == 0 {
		return EmptyExpr()
	}
	return compareSubQuery(gtOp, cs[0], subQuery)
}

// Gte  equivalent column >= (subQuery)
func (cs Columns) Gte(subQuery *gorm.DB) Expr {
	if len(cs) == 0 {
		return EmptyExpr()
	}
	return compareSubQuery(gteOp, cs[0], subQuery)
}

// Lt  equivalent column < (subQuery)
func (cs Columns) Lt(subQuery *gorm.DB) Expr {
	if len(cs) == 0 {
		return EmptyExpr()
	}
	return compareSubQuery(ltOp, cs[0], subQuery)
}

// Lte equivalent column <= (subQuery)
func (cs Columns) Lte(subQuery *gorm.DB) Expr {
	if len(cs) == 0 {
		return EmptyExpr()
	}
	return compareSubQuery(lteOp, cs[0], subQuery)
}

// FindInSet FIND_IN_SET(column, (subQuery))
func (cs Columns) FindInSet(subQuery *gorm.DB) Expr {
	if len(cs) == 0 {
		return EmptyExpr()
	}
	return expr{e: clause.Expr{
		SQL:  "FIND_IN_SET(?,(?))",
		Vars: []any{cs[0].RawExpr(), subQuery},
	}}
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
