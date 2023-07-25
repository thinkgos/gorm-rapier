package assist

import (
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"gorm.io/gorm/schema"
)

// CrossJoinsExpr cross joins condition
func CrossJoinsExpr(table schema.Tabler, conds ...Expr) Condition {
	return CrossJoinsXExpr(table, "", conds...)
}

// CrossJoinsXExpr cross joins condition
func CrossJoinsXExpr(table schema.Tabler, alias string, conds ...Expr) Condition {
	return joinsExpr(clause.CrossJoin, table, alias, conds...)
}

// InnerJoinsExpr inner joins condition
func InnerJoinsExpr(table schema.Tabler, conds ...Expr) Condition {
	return InnerJoinsXExpr(table, "", conds...)
}

// InnerJoinsXExpr inner joins condition
func InnerJoinsXExpr(table schema.Tabler, alias string, conds ...Expr) Condition {
	return joinsExpr(clause.InnerJoin, table, alias, conds...)
}

// LeftJoinsExpr left join condition
func LeftJoinsExpr(table schema.Tabler, conds ...Expr) Condition {
	return LeftJoinsXExpr(table, "", conds...)
}

// LeftJoinsXExpr left join condition
func LeftJoinsXExpr(table schema.Tabler, alias string, conds ...Expr) Condition {
	return joinsExpr(clause.LeftJoin, table, alias, conds...)
}

// RightJoinsExpr right join condition
func RightJoinsExpr(table schema.Tabler, conds ...Expr) Condition {
	return RightJoinsXExpr(table, "", conds...)
}

// RightJoinsXExpr right join condition
func RightJoinsXExpr(table schema.Tabler, alias string, conds ...Expr) Condition {
	return joinsExpr(clause.RightJoin, table, alias, conds...)
}

func joinsExpr(joinType clause.JoinType, table schema.Tabler, alias string, conds ...Expr) Condition {
	return func(db *gorm.DB) *gorm.DB {
		if len(conds) == 0 {
			return db
		}

		clauseFrom := getClauseFrom(db)
		clauseFrom.Joins = append(clauseFrom.Joins, clause.Join{
			Type:  joinType,
			Table: clause.Table{Name: table.TableName(), Alias: alias},
			ON:    clause.Where{Exprs: IntoExpression(conds...)},
		})
		return db.Clauses(clauseFrom)
	}
}

func getClauseFrom(db *gorm.DB) *clause.From {
	if db == nil || db.Statement == nil {
		return &clause.From{}
	}
	c, ok := db.Statement.Clauses[clause.From{}.Name()]
	if !ok || c.Expression == nil {
		return &clause.From{}
	}
	from, ok := c.Expression.(clause.From)
	if !ok {
		return &clause.From{}
	}
	return &from
}
