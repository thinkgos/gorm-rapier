package assist

import (
	"testing"

	"gorm.io/gorm"
)

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
			db: newDb().Model(xDict.X_Model()).
				Where(NewColumns(xDict.Id).In(nil)).
				Find(&dummy),
			wantVars: nil,
			want:     "SELECT * FROM `dict` WHERE ",
		},
		{
			name: "in sub query",
			db: newDb().Model(xDict.X_Model()).
				Where(NewColumns(xDict.Id).In(newDb().Model(xDict.X_Model()).Scopes(SelectExpr(xDict.Id)).Where(xDict.Score.Gt(100)))).
				Find(&dummy),
			wantVars: []any{float64(100)},
			want:     "SELECT * FROM `dict` WHERE `dict`.`id` IN (SELECT `dict`.`id` FROM `dict` WHERE `dict`.`score` > ?)",
		},
		{
			name: "in sub query - (multiple fields)",
			db: newDb().Model(xDict.X_Model()).
				Where(NewColumns(xDict.Id, xDict.Name).In(newDb().Model(xDict.X_Model()).Scopes(SelectExpr(xDict.Id, xDict.Name)).Where(xDict.Score.Gt(100)))).
				Find(&dummy),
			wantVars: []any{float64(100)},
			want:     "SELECT * FROM `dict` WHERE (`dict`.`id`,`dict`.`name`) IN (SELECT `dict`.`id`,`dict`.`name` FROM `dict` WHERE `dict`.`score` > ?)",
		},
		{
			name: "in sub query - (no field)",
			db: newDb().Model(xDict.X_Model()).
				Where(NewColumns().In(newDb().Model(xDict.X_Model()).Scopes(SelectExpr(xDict.Id)).Where(xDict.Score.Gt(100)))).
				Find(&dummy),
			wantVars: nil,
			want:     "SELECT * FROM `dict` WHERE ",
		},
		{
			name: "in value",
			db: newDb().Model(xDict.X_Model()).
				Where(NewColumns(xDict.Id).In(Values([]any{1, 100}))).
				Find(&dummy),
			wantVars: []any{1, 100},
			want:     "SELECT * FROM `dict` WHERE `dict`.`id` IN (?,?)",
		},
		{
			name: "in value - multiple field",
			db: newDb().Model(xDict.X_Model()).
				Where(NewColumns(xDict.Id, xDict.Score).In(Values([][]any{{100, 200}, {1, 2}}))).
				Find(&dummy),
			wantVars: []any{100, 200, 1, 2},
			want:     "SELECT * FROM `dict` WHERE (`dict`.`id`, `dict`.`score`) IN ((?,?),(?,?))",
		},
		{
			name: "in value - (no field)",
			db: newDb().Model(xDict.X_Model()).
				Where(NewColumns().In(Values([][]any{{100, 200}, {1, 2}}))).
				Find(&dummy),
			wantVars: nil,
			want:     "SELECT * FROM `dict` WHERE ",
		},
		{
			name: "not in invalid query",
			db: newDb().Model(xDict.X_Model()).
				Where(NewColumns(xDict.Id).NotIn(nil)).
				Find(&dummy),
			wantVars: nil,
			want:     "SELECT * FROM `dict` WHERE ",
		},
		{
			name: "not in sub query",
			db: newDb().Model(xDict.X_Model()).
				Where(NewColumns(xDict.Id).NotIn(newDb().Model(xDict.X_Model()).Scopes(SelectExpr(xDict.Id)).Where(xDict.Score.Gt(100)))).
				Find(&dummy),
			wantVars: []any{float64(100)},
			want:     "SELECT * FROM `dict` WHERE `dict`.`id` NOT IN (SELECT `dict`.`id` FROM `dict` WHERE `dict`.`score` > ?)",
		},
		{
			name: "not in sub query(no fields)",
			db: newDb().Model(xDict.X_Model()).
				Where(NewColumns().NotIn(newDb().Model(xDict.X_Model()).Scopes(SelectExpr(xDict.Id)).Where(xDict.Score.Gt(100)))).
				Find(&dummy),
			wantVars: nil,
			want:     "SELECT * FROM `dict` WHERE ",
		},
		{
			name: "not in value",
			db: newDb().Model(xDict.X_Model()).
				Where(NewColumns(xDict.Id, xDict.Score).NotIn(Values([][]any{{100, 200}, {1, 2}}))).
				Find(&dummy),
			wantVars: []any{100, 200, 1, 2},
			want:     "SELECT * FROM `dict` WHERE (`dict`.`id`, `dict`.`score`) NOT IN ((?,?),(?,?))",
		},
		{
			name: "not in value(no fields)",
			db: newDb().Model(xDict.X_Model()).
				Where(NewColumns().NotIn(Values([][]any{{100, 200}, {1, 2}}))).
				Find(&dummy),
			wantVars: nil,
			want:     "SELECT * FROM `dict` WHERE ",
		},
		{
			name: "eq",
			db: newDb().Model(xDict.X_Model()).
				Where(
					xDict.Id.EqSubQuery(
						newDb().Model(xDict.X_Model()).Scopes(SelectExpr(xDict.Id.Max())),
					),
				).
				Find(&dummy),
			wantVars: nil,
			want:     "SELECT * FROM `dict` WHERE `dict`.`id` = (SELECT MAX(`dict`.`id`) FROM `dict`)",
		},
		{
			name: "neq",
			db: newDb().Model(xDict.X_Model()).
				Where(xDict.Id.NeqSubQuery(newDb().Model(xDict.X_Model()).Scopes(SelectExpr(xDict.Id.Max())))).
				Find(&dummy),
			wantVars: nil,
			want:     "SELECT * FROM `dict` WHERE `dict`.`id` <> (SELECT MAX(`dict`.`id`) FROM `dict`)",
		},
		{
			name: "gt",
			db: newDb().Model(xDict.X_Model()).
				Where(xDict.Id.GtSubQuery(newDb().Model(xDict.X_Model()).Scopes(SelectExpr(xDict.Id.Min())))).
				Find(&dummy),
			wantVars: nil,
			want:     "SELECT * FROM `dict` WHERE `dict`.`id` > (SELECT MIN(`dict`.`id`) FROM `dict`)",
		},
		{
			name: "gte",
			db: newDb().Model(xDict.X_Model()).
				Where(xDict.Id.GteSubQuery(newDb().Model(xDict.X_Model()).Scopes(SelectExpr(xDict.Id.Min())))).
				Find(&dummy),
			wantVars: nil,
			want:     "SELECT * FROM `dict` WHERE `dict`.`id` >= (SELECT MIN(`dict`.`id`) FROM `dict`)",
		},
		{
			name: "lt",
			db: newDb().Model(xDict.X_Model()).
				Where(xDict.Id.LtSubQuery(newDb().Model(xDict.X_Model()).Scopes(SelectExpr(xDict.Id.Min())))).
				Find(&dummy),
			wantVars: nil,
			want:     "SELECT * FROM `dict` WHERE `dict`.`id` < (SELECT MIN(`dict`.`id`) FROM `dict`)",
		},
		{
			name: "lte",
			db: newDb().Model(xDict.X_Model()).
				Where(xDict.Id.LteSubQuery(newDb().Model(xDict.X_Model()).Scopes(SelectExpr(xDict.Id.Min())))).
				Find(&dummy),
			wantVars: nil,
			want:     "SELECT * FROM `dict` WHERE `dict`.`id` <= (SELECT MIN(`dict`.`id`) FROM `dict`)",
		},
		{
			name: "in",
			db: newDb().Model(xDict.X_Model()).
				Where(xDict.Id.InSubQuery(newDb().Model(xDict.X_Model()).Scopes(SelectExpr(xDict.Id)).Where(xDict.Score.Gt(100)))).
				Find(&dummy),
			wantVars: []any{float64(100)},
			want:     "SELECT * FROM `dict` WHERE `dict`.`id` IN (SELECT `dict`.`id` FROM `dict` WHERE `dict`.`score` > ?)",
		},
		{
			name: "not in",
			db: newDb().Model(xDict.X_Model()).
				Where(xDict.Id.NotInSubQuery(newDb().Model(xDict.X_Model()).Scopes(SelectExpr(xDict.Id)).Where(xDict.Score.Gt(100)))).
				Find(&dummy),
			wantVars: []any{float64(100)},
			want:     "SELECT * FROM `dict` WHERE `dict`.`id` NOT IN (SELECT `dict`.`id` FROM `dict` WHERE `dict`.`score` > ?)",
		},
		{
			name: "find_in_set",
			db: newDb().Model(xDict.X_Model()).
				Where(xDict.Id.FindInSetSubQuery(newDb().Model(xDict.X_Model()).Scopes(SelectExpr(xDict.Id.Min())))).
				Find(&dummy),
			wantVars: nil,
			want:     "SELECT * FROM `dict` WHERE FIND_IN_SET(`dict`.`id`, (SELECT MIN(`dict`.`id`) FROM `dict`))",
		},
		{
			name: "exist",
			db: newDb().Model(xDict.X_Model()).
				Where(Exist(newDb().Model(xDict.X_Model()).Scopes(SelectExpr(xDict.Id.Min())))).
				Find(&dummy),
			wantVars: nil,
			want:     "SELECT * FROM `dict` WHERE EXISTS(SELECT MIN(`dict`.`id`) FROM `dict`)",
		},
		{
			name: "not exist",
			db: newDb().Model(xDict.X_Model()).
				Where(NotExist(newDb().Model(xDict.X_Model()).Scopes(SelectExpr(xDict.Id.Min())))).
				Find(&dummy),
			wantVars: nil,
			want:     "SELECT * FROM `dict` WHERE NOT EXISTS(SELECT MIN(`dict`.`id`) FROM `dict`)",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			CheckBuildExprSql(t, tt.db, tt.want, tt.wantVars)
		})
	}
}
