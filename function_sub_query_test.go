package rapier

import (
	"testing"

	"gorm.io/gorm"
)

func Test_Columns_SubQuery_Assign(t *testing.T) {
	tests := []struct {
		name     string
		db       *gorm.DB
		wantVars []any
		want     string
	}{
		{
			name: "set: empty columns",
			db: newDb().Model(&Dict{}).
				Where(refDict.Id.Eq(1)).
				Clauses(buildClauseSet(
					newDb(),
					[]SetExpr{
						NewColumns().
							Set(
								newDb().Model(&Dict{}).
									Scopes(SelectExpr(
										refDict.Sort,
										refDict.IsPin,
									)),
							),
					},
				)).
				Updates(map[string]any{}),
			wantVars: []any{int64(1)},
			want:     "UPDATE `dict` SET `id`=`id` WHERE `dict`.`id` = ?",
		},
		{
			name: "set: sub query",
			db: newDb().Model(&Dict{}).
				Where(refDict.Id.Eq(1)).
				Clauses(buildClauseSet(
					newDb(),
					[]SetExpr{
						NewColumns(refDict.Sort, refDict.IsPin).
							Set(
								newDb().Model(&Dict{}).
									Scopes(SelectExpr(
										refDict.Sort,
										refDict.IsPin,
									)),
							),
					},
				)).
				Updates(map[string]any{}),
			wantVars: []any{int64(1)},
			want:     "UPDATE `dict` SET (`sort`,`is_pin`)=(SELECT `dict`.`sort`,`dict`.`is_pin` FROM `dict`) WHERE `dict`.`id` = ?",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ReviewBuildDb(t, tt.db, tt.want, tt.wantVars)
		})
	}
}

func Test_Columns_SubQuery(t *testing.T) {
	var dummy []Dict

	tests := []struct {
		name     string
		db       *gorm.DB
		wantVars []any
		want     string
	}{
		{
			name: "in invalid query",
			db: newDb().Model(&Dict{}).
				Where(NewColumns(refDict.Id).In(nil)).
				Find(&dummy),
			wantVars: nil,
			want:     "SELECT * FROM `dict` WHERE ",
		},
		{
			name: "in sub query",
			db: newDb().Model(&Dict{}).
				Where(NewColumns(refDict.Id).In(newDb().Model(&Dict{}).Scopes(SelectExpr(refDict.Id)).Where(refDict.Score.Gt(100)))).
				Find(&dummy),
			wantVars: []any{float64(100)},
			want:     "SELECT * FROM `dict` WHERE `dict`.`id` IN (SELECT `dict`.`id` FROM `dict` WHERE `dict`.`score` > ?)",
		},
		{
			name: "in sub query - (multiple fields)",
			db: newDb().Model(&Dict{}).
				Where(NewColumns(refDict.Id, refDict.Name).In(newDb().Model(&Dict{}).Scopes(SelectExpr(refDict.Id, refDict.Name)).Where(refDict.Score.Gt(100)))).
				Find(&dummy),
			wantVars: []any{float64(100)},
			want:     "SELECT * FROM `dict` WHERE (`dict`.`id`,`dict`.`name`) IN (SELECT `dict`.`id`,`dict`.`name` FROM `dict` WHERE `dict`.`score` > ?)",
		},
		{
			name: "in sub query - (no field)",
			db: newDb().Model(&Dict{}).
				Where(NewColumns().In(newDb().Model(&Dict{}).Scopes(SelectExpr(refDict.Id)).Where(refDict.Score.Gt(100)))).
				Find(&dummy),
			wantVars: nil,
			want:     "SELECT * FROM `dict` WHERE ",
		},
		{
			name: "in value",
			db: newDb().Model(&Dict{}).
				Where(NewColumns(refDict.Id).In([]any{1, 100})).
				Find(&dummy),
			wantVars: []any{1, 100},
			want:     "SELECT * FROM `dict` WHERE `dict`.`id` IN (?,?)",
		},
		{
			name: "in value - multiple field",
			db: newDb().Model(&Dict{}).
				Where(NewColumns(refDict.Id, refDict.Score).In([][]any{{100, 200}, {1, 2}})).
				Find(&dummy),
			wantVars: []any{100, 200, 1, 2},
			want:     "SELECT * FROM `dict` WHERE (`dict`.`id`, `dict`.`score`) IN ((?,?),(?,?))",
		},
		{
			name: "in value - (no field)",
			db: newDb().Model(&Dict{}).
				Where(NewColumns().In([][]any{{100, 200}, {1, 2}})).
				Find(&dummy),
			wantVars: nil,
			want:     "SELECT * FROM `dict` WHERE ",
		},
		{
			name: "not in invalid query",
			db: newDb().Model(&Dict{}).
				Where(NewColumns(refDict.Id).NotIn(nil)).
				Find(&dummy),
			wantVars: nil,
			want:     "SELECT * FROM `dict` WHERE ",
		},
		{
			name: "not in sub query",
			db: newDb().Model(&Dict{}).
				Where(NewColumns(refDict.Id).NotIn(newDb().Model(&Dict{}).Scopes(SelectExpr(refDict.Id)).Where(refDict.Score.Gt(100)))).
				Find(&dummy),
			wantVars: []any{float64(100)},
			want:     "SELECT * FROM `dict` WHERE `dict`.`id` NOT IN (SELECT `dict`.`id` FROM `dict` WHERE `dict`.`score` > ?)",
		},
		{
			name: "not in sub query(no fields)",
			db: newDb().Model(&Dict{}).
				Where(NewColumns().NotIn(newDb().Model(&Dict{}).Scopes(SelectExpr(refDict.Id)).Where(refDict.Score.Gt(100)))).
				Find(&dummy),
			wantVars: nil,
			want:     "SELECT * FROM `dict` WHERE ",
		},
		{
			name: "not in value",
			db: newDb().Model(&Dict{}).
				Where(NewColumns(refDict.Id, refDict.Score).NotIn([][]any{{100, 200}, {1, 2}})).
				Find(&dummy),
			wantVars: []any{100, 200, 1, 2},
			want:     "SELECT * FROM `dict` WHERE (`dict`.`id`, `dict`.`score`) NOT IN ((?,?),(?,?))",
		},
		{
			name: "not in value(no fields)",
			db: newDb().Model(&Dict{}).
				Where(NewColumns().NotIn([][]any{{100, 200}, {1, 2}})).
				Find(&dummy),
			wantVars: nil,
			want:     "SELECT * FROM `dict` WHERE ",
		},
		{
			name: "eq",
			db: newDb().Model(&Dict{}).
				Where(
					refDict.Id.EqSubQuery(
						newDb().Model(&Dict{}).Scopes(SelectExpr(refDict.Id.Max())),
					),
				).
				Find(&dummy),
			wantVars: nil,
			want:     "SELECT * FROM `dict` WHERE `dict`.`id` = (SELECT MAX(`dict`.`id`) FROM `dict`)",
		},
		{
			name: "neq",
			db: newDb().Model(&Dict{}).
				Where(refDict.Id.NeqSubQuery(newDb().Model(&Dict{}).Scopes(SelectExpr(refDict.Id.Max())))).
				Find(&dummy),
			wantVars: nil,
			want:     "SELECT * FROM `dict` WHERE `dict`.`id` <> (SELECT MAX(`dict`.`id`) FROM `dict`)",
		},
		{
			name: "gt",
			db: newDb().Model(&Dict{}).
				Where(refDict.Id.GtSubQuery(newDb().Model(&Dict{}).Scopes(SelectExpr(refDict.Id.Min())))).
				Find(&dummy),
			wantVars: nil,
			want:     "SELECT * FROM `dict` WHERE `dict`.`id` > (SELECT MIN(`dict`.`id`) FROM `dict`)",
		},
		{
			name: "gte",
			db: newDb().Model(&Dict{}).
				Where(refDict.Id.GteSubQuery(newDb().Model(&Dict{}).Scopes(SelectExpr(refDict.Id.Min())))).
				Find(&dummy),
			wantVars: nil,
			want:     "SELECT * FROM `dict` WHERE `dict`.`id` >= (SELECT MIN(`dict`.`id`) FROM `dict`)",
		},
		{
			name: "lt",
			db: newDb().Model(&Dict{}).
				Where(refDict.Id.LtSubQuery(newDb().Model(&Dict{}).Scopes(SelectExpr(refDict.Id.Min())))).
				Find(&dummy),
			wantVars: nil,
			want:     "SELECT * FROM `dict` WHERE `dict`.`id` < (SELECT MIN(`dict`.`id`) FROM `dict`)",
		},
		{
			name: "lte",
			db: newDb().Model(&Dict{}).
				Where(refDict.Id.LteSubQuery(newDb().Model(&Dict{}).Scopes(SelectExpr(refDict.Id.Min())))).
				Find(&dummy),
			wantVars: nil,
			want:     "SELECT * FROM `dict` WHERE `dict`.`id` <= (SELECT MIN(`dict`.`id`) FROM `dict`)",
		},
		{
			name: "in",
			db: newDb().Model(&Dict{}).
				Where(refDict.Id.InSubQuery(newDb().Model(&Dict{}).Scopes(SelectExpr(refDict.Id)).Where(refDict.Score.Gt(100)))).
				Find(&dummy),
			wantVars: []any{float64(100)},
			want:     "SELECT * FROM `dict` WHERE `dict`.`id` IN (SELECT `dict`.`id` FROM `dict` WHERE `dict`.`score` > ?)",
		},
		{
			name: "not in",
			db: newDb().Model(&Dict{}).
				Where(refDict.Id.NotInSubQuery(newDb().Model(&Dict{}).Scopes(SelectExpr(refDict.Id)).Where(refDict.Score.Gt(100)))).
				Find(&dummy),
			wantVars: []any{float64(100)},
			want:     "SELECT * FROM `dict` WHERE `dict`.`id` NOT IN (SELECT `dict`.`id` FROM `dict` WHERE `dict`.`score` > ?)",
		},
		{
			name: "find_in_set",
			db: newDb().Model(&Dict{}).
				Where(refDict.Id.FindInSetSubQuery(newDb().Model(&Dict{}).Scopes(SelectExpr(refDict.Id.Min())))).
				Find(&dummy),
			wantVars: nil,
			want:     "SELECT * FROM `dict` WHERE FIND_IN_SET(`dict`.`id`, (SELECT MIN(`dict`.`id`) FROM `dict`))",
		},
		{
			name: "exist",
			db: newDb().Model(&Dict{}).
				Where(Exist(newDb().Model(&Dict{}).Scopes(SelectExpr(refDict.Id.Min())))).
				Find(&dummy),
			wantVars: nil,
			want:     "SELECT * FROM `dict` WHERE EXISTS(SELECT MIN(`dict`.`id`) FROM `dict`)",
		},
		{
			name: "not exist",
			db: newDb().Model(&Dict{}).
				Where(NotExist(newDb().Model(&Dict{}).Scopes(SelectExpr(refDict.Id.Min())))).
				Find(&dummy),
			wantVars: nil,
			want:     "SELECT * FROM `dict` WHERE NOT EXISTS(SELECT MIN(`dict`.`id`) FROM `dict`)",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ReviewBuildDb(t, tt.db, tt.want, tt.wantVars)
		})
	}
}
