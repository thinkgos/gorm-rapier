package assist

import (
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

// CrossJoins cross joins condition
func CrossJoins(tableName string, conds ...Expr) Condition {
	return CrossJoinsX(tableName, "", conds...)
}

// CrossJoinsX cross joins condition
func CrossJoinsX(tableName, alias string, conds ...Expr) Condition {
	return joins(clause.CrossJoin, tableName, alias, conds...)
}

// InnerJoins inner joins condition
func InnerJoins(tableName string, conds ...Expr) Condition {
	return InnerJoinsX(tableName, "", conds...)
}

// InnerJoinsX inner joins condition
func InnerJoinsX(tableName, alias string, conds ...Expr) Condition {
	return joins(clause.InnerJoin, tableName, alias, conds...)
}

// LeftJoins left join condition
func LeftJoins(tableName string, conds ...Expr) Condition {
	return LeftJoinsX(tableName, "", conds...)
}

// LeftJoinsX left join condition
func LeftJoinsX(tableName, alias string, conds ...Expr) Condition {
	return joins(clause.LeftJoin, tableName, alias, conds...)
}

// RightJoins right join condition
func RightJoins(tableName string, conds ...Expr) Condition {
	return RightJoinsX(tableName, "", conds...)
}

// RightJoinsX right join condition
func RightJoinsX(tableName, alias string, conds ...Expr) Condition {
	return joins(clause.RightJoin, tableName, alias, conds...)
}

func joins(joinType clause.JoinType, tableName, alias string, conds ...Expr) Condition {
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
