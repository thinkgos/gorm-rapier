package assist

import (
	"reflect"
	"strings"

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

// Indirect returns the value that v reflect.Type.
func Indirect(value interface{}) reflect.Type {
	mt := reflect.TypeOf(value)
	if mt.Kind() == reflect.Pointer {
		mt = mt.Elem()
	}
	return mt
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
