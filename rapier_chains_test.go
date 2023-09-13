package rapier

import (
	"context"
	"testing"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

func Test_Condition_Stand(t *testing.T) {
	t.Run("executor", func(t *testing.T) {
		err := xDict.New_Executor(newDb()).
			Debug().
			Scopes(
				NewConditions().
					Unscoped().
					Clauses(clause.Locking{Strength: "UPDATE"}).
					Select([]string{"id", "pid", "name"}).
					Distinct().
					Omit("sort").
					OmitExpr(xDict.Sort).
					Where("id = ?", 1).
					Or("pid = ?", 0).
					Not("is_ping = ?", false).
					Order("created_at").
					Group("name").
					Having("").
					InnerJoins(xDict.X_Alias()).
					Joins(xDict.X_Alias()).
					Limit(10).
					Offset(2).
					Scopes(func(d *gorm.DB) *gorm.DB {
						return d.Where("score > ?", 10)
					}).
					Build()...,
			).
			WithContext(context.Background()).
			Session(&gorm.Session{}).
			Table("a").
			Find(&[]Dict{})
		if err != nil {
			t.Error(err)
		}
	})
}

func Test_Condition_Expr(t *testing.T) {
	var dummy Dict

	xDd := xDict.As("dd")
	xDt := xDictItem

	tests := []struct {
		name     string
		db       *gorm.DB
		wantVars []any
		want     string
	}{
		{
			name: "Expr: Configure",
			db: xDict.New_Executor(newDb()).
				Scopes(
					NewConditions().
						Configure(func(c *Conditions) *Conditions {
							return c.SelectExpr(xDict.Id)
						}).
						Build()...,
				).
				IntoDB().
				Take(&dummy),
			wantVars: nil,
			want:     "SELECT `dict`.`id` FROM `dict` LIMIT 1",
		},
		{
			name: "Expr: select *",
			db: xDict.New_Executor(newDb()).
				Scopes(
					NewConditions().
						SelectExpr().
						Build()...,
				).
				IntoDB().
				Take(&dummy),
			wantVars: nil,
			want:     "SELECT * FROM `dict` LIMIT 1",
		},
		{
			name: "Expr: select field",
			db: xDict.New_Executor(newDb()).
				Scopes(
					NewConditions().
						SelectExpr(
							xDict.Id,
							xDict.CreatedAt.UnixTimestamp().As("created_at"),
							xDict.CreatedAt.UnixTimestamp().IfNull(0).As("created_at1"),
						).
						Build()...,
				).
				IntoDB().
				Take(&dummy),
			wantVars: []any{int64(0)},
			want:     "SELECT `dict`.`id`,UNIX_TIMESTAMP(`dict`.`created_at`) AS `created_at`,IFNULL(UNIX_TIMESTAMP(`dict`.`created_at`),?) AS `created_at1` FROM `dict` LIMIT 1",
		},
		{
			name: "Expr: select * using distinct",
			db: xDict.New_Executor(newDb()).
				Scopes(
					NewConditions().
						DistinctExpr(xDict.Id).
						Build()...,
				).
				IntoDB().
				Take(&dummy),
			wantVars: nil,
			want:     "SELECT DISTINCT `dict`.`id` FROM `dict` LIMIT 1",
		},
		{
			name: "Expr: order",
			db: xDict.New_Executor(newDb()).
				Scopes(
					NewConditions().
						OrderExpr(xDict.Score).
						Build()...,
				).
				IntoDB().
				Take(&dummy),
			wantVars: nil,
			want:     "SELECT * FROM `dict` ORDER BY `dict`.`score` LIMIT 1",
		},
		{
			name: "Expr: group",
			db: xDict.New_Executor(newDb()).
				Scopes(
					NewConditions().
						GroupExpr(xDict.Name).
						Build()...,
				).
				IntoDB().
				Take(&dummy),
			wantVars: nil,
			want:     "SELECT * FROM `dict` GROUP BY `dict`.`name` LIMIT 1",
		},
		{
			name: "Expr: cross join",
			db: xDict.New_Executor(newDb()).
				Scopes(
					NewConditions().
						CrossJoinsExpr(
							&xDt,
							xDt.DictId.EqCol(xDict.Id),
						).
						Build()...,
				).
				IntoDB().
				Take(&dummy),
			wantVars: nil,
			want:     "SELECT `dict`.`id`,`dict`.`pid`,`dict`.`name`,`dict`.`score`,`dict`.`is_pin`,`dict`.`sort`,`dict`.`created_at` FROM `dict` CROSS JOIN `dict_item` ON `dict_item`.`dict_id` = `dict`.`id` LIMIT 1",
		},
		{
			name: "Expr: cross join X",
			db: xDict.New_Executor(newDb()).
				Scopes(
					NewConditions().
						CrossJoinsXExpr(
							&xDd,
							xDd.X_Alias(),
							xDd.Id.EqCol(xDict.Pid),
							xDd.IsPin.Eq(true),
						).
						Build()...,
				).
				IntoDB().
				Take(&dummy),
			wantVars: []any{true},
			want:     "SELECT `dict`.`id`,`dict`.`pid`,`dict`.`name`,`dict`.`score`,`dict`.`is_pin`,`dict`.`sort`,`dict`.`created_at` FROM `dict` CROSS JOIN `dict` `dd` ON `dd`.`id` = `dict`.`pid` AND `dd`.`is_pin` = ? LIMIT 1",
		},
		{
			name: "Expr: inner join",
			db: xDict.New_Executor(newDb()).
				Scopes(
					NewConditions().
						InnerJoinsExpr(&xDt, xDt.DictId.EqCol(xDict.Id)).
						Build()...,
				).
				IntoDB().
				Take(&dummy),
			wantVars: nil,
			want:     "SELECT `dict`.`id`,`dict`.`pid`,`dict`.`name`,`dict`.`score`,`dict`.`is_pin`,`dict`.`sort`,`dict`.`created_at` FROM `dict` INNER JOIN `dict_item` ON `dict_item`.`dict_id` = `dict`.`id` LIMIT 1",
		},
		{
			name: "Expr: inner join X",
			db: xDict.New_Executor(newDb()).
				Scopes(
					NewConditions().
						InnerJoinsXExpr(&xDd, xDd.X_Alias(), xDd.Id.EqCol(xDict.Pid), xDd.IsPin.Eq(true)).
						Build()...,
				).
				IntoDB().
				Take(&dummy),
			wantVars: []any{true},
			want:     "SELECT `dict`.`id`,`dict`.`pid`,`dict`.`name`,`dict`.`score`,`dict`.`is_pin`,`dict`.`sort`,`dict`.`created_at` FROM `dict` INNER JOIN `dict` `dd` ON `dd`.`id` = `dict`.`pid` AND `dd`.`is_pin` = ? LIMIT 1",
		},
		{
			name: "Expr: left join",
			db: xDict.New_Executor(newDb()).
				Scopes(
					NewConditions().
						LeftJoinsExpr(&xDt, xDt.DictId.EqCol(xDict.Id)).
						Build()...,
				).
				IntoDB().
				Take(&dummy),
			wantVars: nil,
			want:     "SELECT `dict`.`id`,`dict`.`pid`,`dict`.`name`,`dict`.`score`,`dict`.`is_pin`,`dict`.`sort`,`dict`.`created_at` FROM `dict` LEFT JOIN `dict_item` ON `dict_item`.`dict_id` = `dict`.`id` LIMIT 1",
		},
		{
			name: "Expr: left join X",
			db: xDict.New_Executor(newDb()).
				Scopes(
					NewConditions().
						LeftJoinsXExpr(&xDd, xDd.X_Alias(), xDd.Id.EqCol(xDict.Pid), xDd.IsPin.Eq(true)).
						Build()...,
				).
				IntoDB().
				Take(&dummy),
			wantVars: []any{true},
			want:     "SELECT `dict`.`id`,`dict`.`pid`,`dict`.`name`,`dict`.`score`,`dict`.`is_pin`,`dict`.`sort`,`dict`.`created_at` FROM `dict` LEFT JOIN `dict` `dd` ON `dd`.`id` = `dict`.`pid` AND `dd`.`is_pin` = ? LIMIT 1",
		},
		{
			name: "Expr: right join",
			db: xDict.New_Executor(newDb()).
				Scopes(
					NewConditions().
						RightJoinsExpr(&xDt, xDt.DictId.EqCol(xDict.Id)).
						Build()...,
				).
				IntoDB().
				Take(&dummy),
			wantVars: nil,
			want:     "SELECT `dict`.`id`,`dict`.`pid`,`dict`.`name`,`dict`.`score`,`dict`.`is_pin`,`dict`.`sort`,`dict`.`created_at` FROM `dict` RIGHT JOIN `dict_item` ON `dict_item`.`dict_id` = `dict`.`id` LIMIT 1",
		},
		{
			name: "Expr: right join X",
			db: xDict.New_Executor(newDb()).
				Scopes(
					NewConditions().
						RightJoinsXExpr(&xDd, xDd.X_Alias(), xDd.Id.EqCol(xDict.Pid), xDd.IsPin.Eq(true)).
						Build()...,
				).
				IntoDB().
				Take(&dummy),
			wantVars: []any{true},
			want:     "SELECT `dict`.`id`,`dict`.`pid`,`dict`.`name`,`dict`.`score`,`dict`.`is_pin`,`dict`.`sort`,`dict`.`created_at` FROM `dict` RIGHT JOIN `dict` `dd` ON `dd`.`id` = `dict`.`pid` AND `dd`.`is_pin` = ? LIMIT 1",
		},
		{
			name: "clause: for update",
			db: xDict.New_Executor(newDb()).
				Scopes(
					NewConditions().
						LockingUpdate().
						Build()...,
				).
				IntoDB().
				Take(&dummy),
			wantVars: nil,
			want:     "SELECT * FROM `dict` LIMIT 1 FOR UPDATE",
		},
		{
			name: "clause: for share",
			db: xDict.New_Executor(newDb()).
				Scopes(
					NewConditions().
						LockingShare().
						Build()...,
				).
				IntoDB().
				Take(&dummy),
			wantVars: nil,
			want:     "SELECT * FROM `dict` LIMIT 1 FOR SHARE",
		},
		{
			name: "clause: pagination",
			db: xDict.New_Executor(newDb()).
				Scopes(
					NewConditions().
						Pagination(2, 10).
						Build()...,
				).
				IntoDB().
				Find(&dummy),
			wantVars: nil,
			want:     "SELECT * FROM `dict` LIMIT 10 OFFSET 10",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			CheckBuildExprSql(t, tt.db, tt.want, tt.wantVars)
		})
	}
}
