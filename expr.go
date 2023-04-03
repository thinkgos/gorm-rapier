package assist

import (
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

// BuildOption build option
type BuildOption uint

const (
	// WithTable build column with table
	WithTable BuildOption = iota

	// WithAll build column with table and alias(if have)
	WithAll

	// WithoutQuote build column without quote
	WithoutQuote
)

type expr struct {
	// default column
	col clause.Column

	// if has expression
	e clause.Expression
	// build column options
	buildOpts []BuildOption
}

func (e expr) ColumnName() string { return e.col.Name }

func (e expr) Expression() clause.Expression {
	if e.e == nil {
		return clause.NamedExpr{SQL: "?", Vars: []any{e.col}}
	}
	return e.e
}

func (e expr) WithTable(table string) Expr {
	e.col.Table = table
	return e
}

func (e expr) BuildColumn(stmt *gorm.Statement, opts ...BuildOption) string {
	col := clause.Column{Name: e.col.Name}
	for _, opt := range append(e.buildOpts, opts...) {
		switch opt {
		case WithTable:
			col.Table = e.col.Table
		case WithAll:
			col.Table = e.col.Table
			col.Alias = e.col.Alias
		case WithoutQuote:
			col.Raw = true
		}
	}
	if col.Name == "*" {
		if col.Table != "" {
			return stmt.Quote(col.Table) + ".*"
		}
		return "*"
	}
	return stmt.Quote(col)
}

func (e expr) Build(builder clause.Builder) {
	if e.e == nil {
		if stmt, ok := builder.(*gorm.Statement); ok {
			builder.WriteString(string(e.BuildColumn(stmt, WithAll)))
			return
		}
	}

	e.e.Build(builder)
}

func (e expr) BuildWithArgs(stmt *gorm.Statement) (string, []any) {
	if e.e == nil {
		return e.BuildColumn(stmt, WithAll), nil
	}
	newStmt := &gorm.Statement{DB: stmt.DB, Table: stmt.Table, Schema: stmt.Schema}
	e.e.Build(newStmt)
	return newStmt.SQL.String(), newStmt.Vars
}

func (e expr) RawExpr() expression {
	if e.e == nil {
		return e.col
	}
	return e.e
}

func (e expr) withAppendBuildOpts(opts ...BuildOption) expr {
	e.buildOpts = append(e.buildOpts, opts...)
	return e
}
