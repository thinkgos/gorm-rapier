package rapier

import (
	"context"
	"testing"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

func Test_Condition_Stand(t *testing.T) {
	t.Run("executor", func(t *testing.T) {
		err := refDict.New_Executor(newDb()).
			Debug().
			Scopes(
				NewConditions().
					Unscoped().
					Clauses(clause.Locking{Strength: "UPDATE"}).
					Select([]string{"id", "pid", "name"}).
					Distinct().
					Omit("sort").
					OmitExpr(refDict.Sort).
					Where("id = ?", 1).
					Or("pid = ?", 0).
					Not("is_ping = ?", false).
					Order("created_at").
					Group("name").
					Having("").
					InnerJoins(refDict.Alias()).
					Joins(refDict.Alias()).
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
	t.Run("preload", func(t *testing.T) {
		var td TestDict

		_ = refDict.New_Executor(newDbWithLog()).
			Scopes(NewConditions().Preload("DictItem").Build()...).
			Take(&td)
	})
}

func Test_Condition_Expr(t *testing.T) {
	var dummy Dict

	xDd := refDict.As("dd")
	xDt := refDictItem

	tests := []struct {
		name     string
		db       *gorm.DB
		wantVars []any
		want     string
	}{
		{
			name: "Expr: Configure",
			db: refDict.New_Executor(newDb()).
				Scopes(
					NewConditions().
						Configure(func(c *Conditions) *Conditions {
							return c.SelectExpr(refDict.Id)
						}).
						Build()...,
				).
				IntoDB().
				Take(&dummy),
			wantVars: []any{1},
			want:     "SELECT `dict`.`id` FROM `dict` LIMIT ?",
		},
		{
			name: "Expr: select *",
			db: refDict.New_Executor(newDb()).
				Scopes(
					NewConditions().
						SelectExpr().
						Build()...,
				).
				IntoDB().
				Take(&dummy),
			wantVars: []any{1},
			want:     "SELECT * FROM `dict` LIMIT ?",
		},
		{
			name: "Expr: select field",
			db: refDict.New_Executor(newDb()).
				Scopes(
					NewConditions().
						SelectExpr(
							refDict.Id,
							refDict.CreatedAt.UnixTimestamp().As("created_at"),
							refDict.CreatedAt.UnixTimestamp().IfNull(0).As("created_at1"),
						).
						Build()...,
				).
				IntoDB().
				Take(&dummy),
			wantVars: []any{int64(0), 1},
			want:     "SELECT `dict`.`id`,UNIX_TIMESTAMP(`dict`.`created_at`) AS `created_at`,IFNULL(UNIX_TIMESTAMP(`dict`.`created_at`),?) AS `created_at1` FROM `dict` LIMIT ?",
		},
		{
			name: "Expr: select * using distinct",
			db: refDict.New_Executor(newDb()).
				Scopes(
					NewConditions().
						DistinctExpr(refDict.Id).
						Build()...,
				).
				IntoDB().
				Take(&dummy),
			wantVars: []any{1},
			want:     "SELECT DISTINCT `dict`.`id` FROM `dict` LIMIT ?",
		},
		{
			name: "Expr: order",
			db: refDict.New_Executor(newDb()).
				Scopes(
					NewConditions().
						OrderExpr(refDict.Score).
						Build()...,
				).
				IntoDB().
				Take(&dummy),
			wantVars: []any{1},
			want:     "SELECT * FROM `dict` ORDER BY `dict`.`score` LIMIT ?",
		},
		{
			name: "Expr: group",
			db: refDict.New_Executor(newDb()).
				Scopes(
					NewConditions().
						GroupExpr(refDict.Name).
						Build()...,
				).
				IntoDB().
				Take(&dummy),
			wantVars: []any{1},
			want:     "SELECT * FROM `dict` GROUP BY `dict`.`name` LIMIT ?",
		},
		{
			name: "Expr: cross join",
			db: refDict.New_Executor(newDb()).
				Scopes(
					NewConditions().
						CrossJoinsExpr(
							&xDt,
							xDt.DictId.EqCol(refDict.Id),
						).
						Build()...,
				).
				IntoDB().
				Take(&dummy),
			wantVars: []any{1},
			want:     "SELECT `dict`.`id`,`dict`.`pid`,`dict`.`name`,`dict`.`score`,`dict`.`is_pin`,`dict`.`sort`,`dict`.`created_at` FROM `dict` CROSS JOIN `dict_item` ON `dict_item`.`dict_id` = `dict`.`id` LIMIT ?",
		},
		{
			name: "Expr: cross join with alias",
			db: refDict.New_Executor(newDb()).
				Scopes(
					NewConditions().
						CrossJoinsExpr(
							NewJoinTable(&xDd, xDd.Alias()),
							xDd.Id.EqCol(refDict.Pid),
							xDd.IsPin.Eq(true),
						).
						Build()...,
				).
				IntoDB().
				Take(&dummy),
			wantVars: []any{true, 1},
			want:     "SELECT `dict`.`id`,`dict`.`pid`,`dict`.`name`,`dict`.`score`,`dict`.`is_pin`,`dict`.`sort`,`dict`.`created_at` FROM `dict` CROSS JOIN `dict` `dd` ON `dd`.`id` = `dict`.`pid` AND `dd`.`is_pin` = ? LIMIT ?",
		},
		{
			name: "Expr: inner join",
			db: refDict.New_Executor(newDb()).
				Scopes(
					NewConditions().
						InnerJoinsExpr(&xDt, xDt.DictId.EqCol(refDict.Id)).
						Build()...,
				).
				IntoDB().
				Take(&dummy),
			wantVars: []any{1},
			want:     "SELECT `dict`.`id`,`dict`.`pid`,`dict`.`name`,`dict`.`score`,`dict`.`is_pin`,`dict`.`sort`,`dict`.`created_at` FROM `dict` INNER JOIN `dict_item` ON `dict_item`.`dict_id` = `dict`.`id` LIMIT ?",
		},
		{
			name: "Expr: inner join with alias",
			db: refDict.New_Executor(newDb()).
				Scopes(
					NewConditions().
						InnerJoinsExpr(NewJoinTable(&xDd, xDd.Alias()), xDd.Id.EqCol(refDict.Pid), xDd.IsPin.Eq(true)).
						Build()...,
				).
				IntoDB().
				Take(&dummy),
			wantVars: []any{true, 1},
			want:     "SELECT `dict`.`id`,`dict`.`pid`,`dict`.`name`,`dict`.`score`,`dict`.`is_pin`,`dict`.`sort`,`dict`.`created_at` FROM `dict` INNER JOIN `dict` `dd` ON `dd`.`id` = `dict`.`pid` AND `dd`.`is_pin` = ? LIMIT ?",
		},
		{
			name: "Expr: left join",
			db: refDict.New_Executor(newDb()).
				Scopes(
					NewConditions().
						LeftJoinsExpr(&xDt, xDt.DictId.EqCol(refDict.Id)).
						Build()...,
				).
				IntoDB().
				Take(&dummy),
			wantVars: []any{1},
			want:     "SELECT `dict`.`id`,`dict`.`pid`,`dict`.`name`,`dict`.`score`,`dict`.`is_pin`,`dict`.`sort`,`dict`.`created_at` FROM `dict` LEFT JOIN `dict_item` ON `dict_item`.`dict_id` = `dict`.`id` LIMIT ?",
		},
		{
			name: "Expr: left join with alias",
			db: refDict.New_Executor(newDb()).
				Scopes(
					NewConditions().
						LeftJoinsExpr(NewJoinTable(&xDd, xDd.Alias()), xDd.Id.EqCol(refDict.Pid), xDd.IsPin.Eq(true)).
						Build()...,
				).
				IntoDB().
				Take(&dummy),
			wantVars: []any{true, 1},
			want:     "SELECT `dict`.`id`,`dict`.`pid`,`dict`.`name`,`dict`.`score`,`dict`.`is_pin`,`dict`.`sort`,`dict`.`created_at` FROM `dict` LEFT JOIN `dict` `dd` ON `dd`.`id` = `dict`.`pid` AND `dd`.`is_pin` = ? LIMIT ?",
		},
		{
			name: "Expr: right join",
			db: refDict.New_Executor(newDb()).
				Scopes(
					NewConditions().
						RightJoinsExpr(&xDt, xDt.DictId.EqCol(refDict.Id)).
						Build()...,
				).
				IntoDB().
				Take(&dummy),
			wantVars: []any{1},
			want:     "SELECT `dict`.`id`,`dict`.`pid`,`dict`.`name`,`dict`.`score`,`dict`.`is_pin`,`dict`.`sort`,`dict`.`created_at` FROM `dict` RIGHT JOIN `dict_item` ON `dict_item`.`dict_id` = `dict`.`id` LIMIT ?",
		},
		{
			name: "Expr: right join with alias",
			db: refDict.New_Executor(newDb()).
				Scopes(
					NewConditions().
						RightJoinsExpr(NewJoinTable(&xDd, xDd.Alias()), xDd.Id.EqCol(refDict.Pid), xDd.IsPin.Eq(true)).
						Build()...,
				).
				IntoDB().
				Take(&dummy),
			wantVars: []any{true, 1},
			want:     "SELECT `dict`.`id`,`dict`.`pid`,`dict`.`name`,`dict`.`score`,`dict`.`is_pin`,`dict`.`sort`,`dict`.`created_at` FROM `dict` RIGHT JOIN `dict` `dd` ON `dd`.`id` = `dict`.`pid` AND `dd`.`is_pin` = ? LIMIT ?",
		},
		{
			name: "clause: for update",
			db: refDict.New_Executor(newDb()).
				Scopes(
					NewConditions().
						LockingUpdate().
						Build()...,
				).
				IntoDB().
				Take(&dummy),
			wantVars: []any{1},
			want:     "SELECT * FROM `dict` LIMIT ? FOR UPDATE",
		},
		{
			name: "clause: for share",
			db: refDict.New_Executor(newDb()).
				Scopes(
					NewConditions().
						LockingShare().
						Build()...,
				).
				IntoDB().
				Take(&dummy),
			wantVars: []any{1},
			want:     "SELECT * FROM `dict` LIMIT ? FOR SHARE",
		},
		{
			name: "clause: pagination",
			db: refDict.New_Executor(newDb()).
				Scopes(
					NewConditions().
						Pagination(2, 10).
						Build()...,
				).
				IntoDB().
				Find(&dummy),
			wantVars: []any{10, 10},
			want:     "SELECT * FROM `dict` LIMIT ? OFFSET ?",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ReviewBuildDb(t, tt.db, tt.want, tt.wantVars)
		})
	}
}
