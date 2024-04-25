package rapier

import (
	"reflect"
	"strings"

	"golang.org/x/exp/constraints"
	"gorm.io/gorm"
	"gorm.io/gorm/callbacks"
	"gorm.io/gorm/clause"
)

func buildSelectValue(stmt *gorm.Statement, exprs ...Expr) (query string, args []any) {
	if len(exprs) == 0 {
		return "", nil
	}

	queryItems := make([]string, 0, len(exprs))
	for _, e := range exprs {
		sql, vars := e.BuildWithArgs(stmt)
		queryItems = append(queryItems, sql)
		args = append(args, vars...)
	}
	if len(args) == 0 {
		return queryItems[0], intoAnySlice(queryItems[1:])
	}
	return strings.Join(queryItems, ","), args
}

func buildColumnsValue(db *gorm.DB, columns ...Expr) string {
	stmt := &gorm.Statement{
		DB:     db.Statement.DB,
		Table:  db.Statement.Table,
		Schema: db.Statement.Schema,
	}
	for i, column := range columns {
		if i != 0 {
			_ = stmt.WriteByte(',')
		}
		column.Build(stmt)
	}

	return stmt.SQL.String()
}

// buildClauseSet build all set
func buildClauseSet(db *gorm.DB, exprs []AssignExpr) (set clause.Set) {
	for _, expr := range exprs {
		column := clause.Column{
			Table: "", // FIXME: when need table?.
			Name:  expr.ColumnName(),
		}

		switch e := expr.SetExpr().(type) {
		case clause.Expr:
			set = append(set, clause.Assignment{
				Column: column,
				Value:  e,
			})
		case clause.Eq:
			set = append(set, clause.Assignment{
				Column: column,
				Value:  e.Value,
			})
		case clause.Set:
			set = append(set, e...)
		}
	}
	stmt := db.Session(&gorm.Session{}).Statement
	stmt.Dest = map[string]any{}
	return append(set, callbacks.ConvertToAssignments(stmt)...)
}

func buildAttrsValue(attrs []AssignExpr) []any {
	values := make([]any, 0, len(attrs))
	for _, expr := range attrs {
		switch e := expr.SetExpr().(type) {
		case clause.Eq:
			values = append(values, e)
		case clause.Set:
			for _, v := range e {
				values = append(values, clause.Eq{
					Column: v.Column,
					Value:  v.Value,
				})
			}
		}
	}
	return values
}

func buildColumnName(columns ...Expr) []string {
	vs := make([]string, 0, len(columns))
	for _, v := range columns {
		vs = append(vs, v.ColumnName())
	}
	return vs
}

// IntoExpression convert Expr to clause.Expression
func IntoExpression(conds ...Expr) []clause.Expression {
	exprs := make([]clause.Expression, len(conds))
	for i, cond := range conds {
		if cond != nil {
			exprs[i] = cond.Expression()
		}
	}
	return exprs
}

func intoAnySlice[T any](values []T) []any {
	slices := make([]any, len(values))
	for i, v := range values {
		slices[i] = v
	}
	return slices
}

func IntoSlice[T any, R any](values []T, f func(T) R) []R {
	slices := make([]R, len(values))
	for i, v := range values {
		slices[i] = f(v)
	}
	return slices
}

func IntoIntegerSlice[T constraints.Integer, R constraints.Integer](values []T) []R {
	slices := make([]R, len(values))
	for i, v := range values {
		slices[i] = R(v)
	}
	return slices
}

// intoClauseIN clause.IN{} OR EmptyExpr() expression
func intoClauseIN(column, value any) clause.Expression {
	reflectValue := reflect.Indirect(reflect.ValueOf(value))
	for reflectValue.Kind() == reflect.Ptr {
		reflectValue = reflectValue.Elem()
	}

	kind := reflectValue.Kind()
	if kind == reflect.Slice || kind == reflect.Array {
		valueLen := reflectValue.Len()
		values := make([]any, valueLen)
		for i := 0; i < valueLen; i++ {
			values[i] = reflectValue.Index(i).Interface()
		}
		if len(values) > 0 {
			return clause.IN{Column: column, Values: values}
		}
	}
	return EmptyExpr()
}
