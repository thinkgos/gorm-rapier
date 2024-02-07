package rapier

import (
	"testing"

	"gorm.io/gorm"
)

func Test_Joins(t *testing.T) {
	var dummy Dict

	xDi := xDictItem.As("di")
	_ = xDi
	tests := []struct {
		name     string
		db       *gorm.DB
		wantVars []any
		want     string
	}{
		{
			name: "inner join - empty conds",
			db: newDb().Model(&Dict{}).
				Scopes(
					InnerJoinsExpr(&xDictItem),
				).
				Take(&dummy),
			wantVars: []any{1},
			want:     "SELECT * FROM `dict` LIMIT ?",
		},
		{
			name: "cross join",
			db: newDb().Model(&Dict{}).
				Scopes(
					CrossJoinsExpr(&xDictItem, xDictItem.DictId.EqCol(xDict.Id), xDictItem.IsEnabled.Eq(true)),
				).
				Take(&dummy),
			wantVars: []any{true, 1},
			want:     "SELECT `dict`.`id`,`dict`.`pid`,`dict`.`name`,`dict`.`score`,`dict`.`is_pin`,`dict`.`sort`,`dict`.`created_at` FROM `dict` CROSS JOIN `dict_item` ON `dict_item`.`dict_id` = `dict`.`id` AND `dict_item`.`is_enabled` = ? LIMIT ?",
		},
		{
			name: "inner join",
			db: newDb().Model(&Dict{}).
				Scopes(
					InnerJoinsExpr(&xDictItem, xDictItem.DictId.EqCol(xDict.Id), xDictItem.IsEnabled.Eq(true)),
				).
				Take(&dummy),
			wantVars: []any{true, 1},
			want:     "SELECT `dict`.`id`,`dict`.`pid`,`dict`.`name`,`dict`.`score`,`dict`.`is_pin`,`dict`.`sort`,`dict`.`created_at` FROM `dict` INNER JOIN `dict_item` ON `dict_item`.`dict_id` = `dict`.`id` AND `dict_item`.`is_enabled` = ? LIMIT ?",
		},
		{
			name: "left join",
			db: newDb().Model(&Dict{}).
				Scopes(
					LeftJoinsExpr(&xDictItem, xDictItem.DictId.EqCol(xDict.Id), xDictItem.IsEnabled.Eq(true)),
				).
				Take(&dummy),
			wantVars: []any{true, 1},
			want:     "SELECT `dict`.`id`,`dict`.`pid`,`dict`.`name`,`dict`.`score`,`dict`.`is_pin`,`dict`.`sort`,`dict`.`created_at` FROM `dict` LEFT JOIN `dict_item` ON `dict_item`.`dict_id` = `dict`.`id` AND `dict_item`.`is_enabled` = ? LIMIT ?",
		},
		{
			name: "right join",
			db: newDb().Model(&Dict{}).
				Scopes(
					RightJoinsExpr(&xDictItem, xDictItem.DictId.EqCol(xDict.Id), xDictItem.IsEnabled.Eq(true)),
				).
				Take(&dummy),
			wantVars: []any{true, 1},
			want:     "SELECT `dict`.`id`,`dict`.`pid`,`dict`.`name`,`dict`.`score`,`dict`.`is_pin`,`dict`.`sort`,`dict`.`created_at` FROM `dict` RIGHT JOIN `dict_item` ON `dict_item`.`dict_id` = `dict`.`id` AND `dict_item`.`is_enabled` = ? LIMIT ?",
		},
		{
			name: "inner join - multiple",
			db: newDb().Model(&Dict{}).
				Scopes(
					InnerJoinsExpr(&xDictItem, xDictItem.DictId.EqCol(xDict.Id)),
					InnerJoinsXExpr(&xDi, xDi.X_Alias(), xDi.IsEnabled.Eq(true)),
				).
				Take(&dummy),
			wantVars: []any{true, 1},
			want:     "SELECT `dict`.`id`,`dict`.`pid`,`dict`.`name`,`dict`.`score`,`dict`.`is_pin`,`dict`.`sort`,`dict`.`created_at` FROM `dict` INNER JOIN `dict_item` ON `dict_item`.`dict_id` = `dict`.`id` INNER JOIN `dict_item` `di` ON `di`.`is_enabled` = ? LIMIT ?",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			CheckBuildExprSql(t, tt.db, tt.want, tt.wantVars)
		})
	}
}
