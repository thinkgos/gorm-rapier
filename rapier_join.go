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

func joinsExpr(joinType clause.JoinType, table schema.Tabler, conds ...Expr) Condition {
	return func(db *gorm.DB) *gorm.DB {
		if len(conds) == 0 {
			return db
		}
		tableName := table.TableName()
		join := clause.Join{
			Type:  joinType,
			Table: clause.Table{Name: tableName},
			ON:    clause.Where{Exprs: IntoExpression(conds...)},
		}
		// if table implement Alias interface, then we can use alias.
		if al, ok := table.(Alias); ok {
			if alias := al.Alias(); alias != "" && alias != tableName {
				join.Table.Alias = alias
			}
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
				if jt.alias != "" && jt.alias != tableName {
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
// if table implement `Alias` interface too, you can directly use it, not need this api.
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
