package rapier

import (
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"gorm.io/gorm/schema"
)

// CrossJoinsExpr cross joins condition
func CrossJoinsExpr(table schema.Tabler, conds ...Expr) Condition {
	return joinsExpr(clause.CrossJoin, table, conds...)
}

// InnerJoinsExpr inner joins condition
func InnerJoinsExpr(table schema.Tabler, conds ...Expr) Condition {
	return joinsExpr(clause.InnerJoin, table, conds...)
}

// LeftJoinsExpr left join condition
func LeftJoinsExpr(table schema.Tabler, conds ...Expr) Condition {
	return joinsExpr(clause.LeftJoin, table, conds...)
}

// RightJoinsExpr right join condition
func RightJoinsExpr(table schema.Tabler, conds ...Expr) Condition {
	return joinsExpr(clause.RightJoin, table, conds...)
}

// Deprecated: use other CrossJoinsExpr(NewJoinTable(table, alias), conds...).
// CrossJoinsXExpr cross joins condition
func CrossJoinsXExpr(table schema.Tabler, alias string, conds ...Expr) Condition {
	return CrossJoinsExpr(NewJoinTable(table, alias), conds...)
}

// Deprecated: use other InnerJoinsExpr(NewJoinTable(table, alias), conds...).
// InnerJoinsXExpr inner joins condition
func InnerJoinsXExpr(table schema.Tabler, alias string, conds ...Expr) Condition {
	return InnerJoinsExpr(NewJoinTable(table, alias), conds...)
}

// Deprecated: use other LeftJoinsExpr(NewJoinTable(table, alias), conds...).
// LeftJoinsXExpr left join condition
func LeftJoinsXExpr(table schema.Tabler, alias string, conds ...Expr) Condition {
	return LeftJoinsExpr(NewJoinTable(table, alias), conds...)
}

// Deprecated: use other RightJoinsExpr(NewJoinTable(table, alias), conds...).
// RightJoinsXExpr right join condition
func RightJoinsXExpr(table schema.Tabler, alias string, conds ...Expr) Condition {
	return RightJoinsExpr(NewJoinTable(table, alias), conds...)
}

func joinsExpr(joinType clause.JoinType, table schema.Tabler, conds ...Expr) Condition {
	return func(db *gorm.DB) *gorm.DB {
		if len(conds) == 0 {
			return db
		}

		join := clause.Join{
			Type:  joinType,
			Table: clause.Table{Name: table.TableName()},
			ON:    clause.Where{Exprs: IntoExpression(conds...)},
		}
		if jt, ok := table.(*JoinTable); ok {
			switch tb := jt.table.(type) {
			case *gorm.DB:
				newDb := db.Session(&gorm.Session{NewDB: true})
				join.Expression = JoinTableExpr{
					Join:      join,
					TableExpr: TableExpr(From{Alias: jt.alias, SubQuery: tb})(newDb).Statement.TableExpr,
				}
			case schema.Tabler:
				if jt.alias != "" {
					join.Table.Alias = jt.alias
				}
			default:
				// do nothing
			}
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

type JoinTable struct {
	table any // table(schema.Tabler) or table subquery(*gorm.DB)
	alias string
}

// NewJoinTable new join table as alias
func NewJoinTable(table schema.Tabler, alias string) *JoinTable {
	return &JoinTable{
		table: table,
		alias: alias,
	}
}

// NewJoinTableSubQuery new join table sub query as alias
func NewJoinTableSubQuery(subQuery *gorm.DB, alias string) *JoinTable {
	return &JoinTable{
		table: subQuery,
		alias: alias,
	}
}

// TableName implement schema.Tabler
func (jt *JoinTable) TableName() string {
	if jt.table == nil {
		return jt.alias
	}
	switch t := jt.table.(type) {
	case schema.Tabler:
		return t.TableName()
	// case *gorm.DB:
	default:
	}
	return jt.alias
}
