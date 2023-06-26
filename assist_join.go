package assist

import (
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

// CrossJoinsExpr cross joins condition
func CrossJoinsExpr(tableName string, conds ...Expr) Condition {
	return CrossJoinsXExpr(tableName, "", conds...)
}

// CrossJoinsXExpr cross joins condition
func CrossJoinsXExpr(tableName, alias string, conds ...Expr) Condition {
	return joinsExpr(clause.CrossJoin, tableName, alias, conds...)
}

// InnerJoinsExpr inner joins condition
func InnerJoinsExpr(tableName string, conds ...Expr) Condition {
	return InnerJoinsXExpr(tableName, "", conds...)
}

// InnerJoinsXExpr inner joins condition
func InnerJoinsXExpr(tableName, alias string, conds ...Expr) Condition {
	return joinsExpr(clause.InnerJoin, tableName, alias, conds...)
}

// LeftJoinsExpr left join condition
func LeftJoinsExpr(tableName string, conds ...Expr) Condition {
	return LeftJoinsXExpr(tableName, "", conds...)
}

// LeftJoinsXExpr left join condition
func LeftJoinsXExpr(tableName, alias string, conds ...Expr) Condition {
	return joinsExpr(clause.LeftJoin, tableName, alias, conds...)
}

// RightJoinsExpr right join condition
func RightJoinsExpr(tableName string, conds ...Expr) Condition {
	return RightJoinsXExpr(tableName, "", conds...)
}

// RightJoinsXExpr right join condition
func RightJoinsXExpr(tableName, alias string, conds ...Expr) Condition {
	return joinsExpr(clause.RightJoin, tableName, alias, conds...)
}

func joinsExpr(joinType clause.JoinType, tableName, alias string, conds ...Expr) Condition {
	return func(db *gorm.DB) *gorm.DB {
		if len(conds) == 0 {
			return db
		}

		clauseFrom := getClauseFrom(db)
		clauseFrom.Joins = append(clauseFrom.Joins, clause.Join{
			Type:  joinType,
			Table: clause.Table{Name: tableName, Alias: alias},
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
