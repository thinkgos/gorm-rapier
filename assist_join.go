package assist

import (
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

// CrossJoin cross joins condition
func CrossJoin(tableName string, conds ...Expr) Condition {
	return CrossJoinX(tableName, "", conds...)
}

// CrossJoinX cross joins condition
func CrossJoinX(tableName, alias string, conds ...Expr) Condition {
	return join(clause.CrossJoin, tableName, alias, conds...)
}

// Join same as InnerJoin.
func Join(tableName string, conds ...Expr) Condition {
	return JoinX(tableName, "", conds...)
}

// JoinX same as InnerJoinX.
func JoinX(tableName, alias string, conds ...Expr) Condition {
	return InnerJoinX(tableName, alias, conds...)
}

// InnerJoin inner joins condition
func InnerJoin(tableName string, conds ...Expr) Condition {
	return InnerJoinX(tableName, "", conds...)
}

// InnerJoinX inner joins condition
func InnerJoinX(tableName, alias string, conds ...Expr) Condition {
	return join(clause.InnerJoin, tableName, alias, conds...)
}

// LeftJoin left join condition
func LeftJoin(tableName string, conds ...Expr) Condition {
	return LeftJoinX(tableName, "", conds...)
}

// LeftJoinX left join condition
func LeftJoinX(tableName, alias string, conds ...Expr) Condition {
	return join(clause.LeftJoin, tableName, alias, conds...)
}

// RightJoin right join condition
func RightJoin(tableName string, conds ...Expr) Condition {
	return RightJoinX(tableName, "", conds...)
}

// RightJoinX right join condition
func RightJoinX(tableName, alias string, conds ...Expr) Condition {
	return join(clause.RightJoin, tableName, alias, conds...)
}

func join(joinType clause.JoinType, tableName, alias string, conds ...Expr) Condition {
	return func(db *gorm.DB) *gorm.DB {
		if len(conds) == 0 {
			return db
		}

		join := clause.Join{
			Type:  joinType,
			Table: clause.Table{Name: tableName, Alias: alias},
			ON:    clause.Where{Exprs: IntoExpression(conds...)},
		}
		clauseFrom := getClauseFrom(db)
		clauseFrom.Joins = append(clauseFrom.Joins, join)
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
