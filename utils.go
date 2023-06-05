package assist

import (
	"reflect"
	"strings"

	"golang.org/x/exp/constraints"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

func buildSelectValue(stmt *gorm.Statement, exprs ...Expr) (query string, args []any) {
	if len(exprs) == 0 {
		return "", nil
	}

	var queryItems []string
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

func intoAnySlice[T any](values ...T) []any {
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

// clause.IN{} expression
func intoInExpr(column, value any) clause.Expression {
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
