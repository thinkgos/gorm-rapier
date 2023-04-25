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
		return queryItems[0], intoSlice(queryItems[1:])
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
func intoSlice[T any](values ...T) []any {
	slice := make([]any, len(values))
	for i, v := range values {
		slice[i] = v
	}
	return slice
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
