package rapier

import (
	"testing"

	"gorm.io/gorm"
)

func Test_Joins(t *testing.T) {
	var dummy Dict

	xDi := refDictItem.As("di")
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
					InnerJoinsExpr(&refDictItem),
				).
				Take(&dummy),
			wantVars: []any{1},
			want:     "SELECT * FROM `dict` LIMIT ?",
		},
		{
			name: "cross join",
			db: newDb().Model(&Dict{}).
				Scopes(
					CrossJoinsExpr(&refDictItem, refDictItem.DictId.EqCol(refDict.Id), refDictItem.IsEnabled.Eq(true)),
				).
				Take(&dummy),
			wantVars: []any{true, 1},
			want:     "SELECT `dict`.`id`,`dict`.`pid`,`dict`.`name`,`dict`.`score`,`dict`.`is_pin`,`dict`.`sort`,`dict`.`created_at` FROM `dict` CROSS JOIN `dict_item` ON `dict_item`.`dict_id` = `dict`.`id` AND `dict_item`.`is_enabled` = ? LIMIT ?",
		},
		{
			name: "inner join",
			db: newDb().Model(&Dict{}).
				Scopes(
					InnerJoinsExpr(&refDictItem, refDictItem.DictId.EqCol(refDict.Id), refDictItem.IsEnabled.Eq(true)),
				).
				Take(&dummy),
			wantVars: []any{true, 1},
			want:     "SELECT `dict`.`id`,`dict`.`pid`,`dict`.`name`,`dict`.`score`,`dict`.`is_pin`,`dict`.`sort`,`dict`.`created_at` FROM `dict` INNER JOIN `dict_item` ON `dict_item`.`dict_id` = `dict`.`id` AND `dict_item`.`is_enabled` = ? LIMIT ?",
		},
		{
			name: "left join",
			db: newDb().Model(&Dict{}).
				Scopes(
					LeftJoinsExpr(&refDictItem, refDictItem.DictId.EqCol(refDict.Id), refDictItem.IsEnabled.Eq(true)),
				).
				Take(&dummy),
			wantVars: []any{true, 1},
			want:     "SELECT `dict`.`id`,`dict`.`pid`,`dict`.`name`,`dict`.`score`,`dict`.`is_pin`,`dict`.`sort`,`dict`.`created_at` FROM `dict` LEFT JOIN `dict_item` ON `dict_item`.`dict_id` = `dict`.`id` AND `dict_item`.`is_enabled` = ? LIMIT ?",
		},
		{
			name: "right join",
			db: newDb().Model(&Dict{}).
				Scopes(
					RightJoinsExpr(&refDictItem, refDictItem.DictId.EqCol(refDict.Id), refDictItem.IsEnabled.Eq(true)),
				).
				Take(&dummy),
			wantVars: []any{true, 1},
			want:     "SELECT `dict`.`id`,`dict`.`pid`,`dict`.`name`,`dict`.`score`,`dict`.`is_pin`,`dict`.`sort`,`dict`.`created_at` FROM `dict` RIGHT JOIN `dict_item` ON `dict_item`.`dict_id` = `dict`.`id` AND `dict_item`.`is_enabled` = ? LIMIT ?",
		},
		{
			name: "inner join - multiple",
			db: newDb().Model(&Dict{}).
				Scopes(
					InnerJoinsExpr(&refDictItem, refDictItem.DictId.EqCol(refDict.Id)),
					InnerJoinsXExpr(&xDi, xDi.Ref_Alias(), xDi.IsEnabled.Eq(true)),
				).
				Take(&dummy),
			wantVars: []any{true, 1},
			want:     "SELECT `dict`.`id`,`dict`.`pid`,`dict`.`name`,`dict`.`score`,`dict`.`is_pin`,`dict`.`sort`,`dict`.`created_at` FROM `dict` INNER JOIN `dict_item` ON `dict_item`.`dict_id` = `dict`.`id` INNER JOIN `dict_item` `di` ON `di`.`is_enabled` = ? LIMIT ?",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ReviewBuildDb(t, tt.db, tt.want, tt.wantVars)
		})
	}
}

func Test_SubJoins(t *testing.T) {
	var dummy Dict

	xDi := refDictItem.As("di")
	tests := []struct {
		name     string
		db       *gorm.DB
		wantVars []any
		want     string
	}{
		{
			name: "inner join - sub join",
			db: refDict.New_Executor(newDb()).
				SelectExpr(
					refDict.Id,
					xDi.Name,
				).
				InnerJoinsExpr(
					NewJoinTableSubQuery(
						refDictItem.New_Executor(newDb()).
							SelectExpr(refDictItem.DictId, refDictItem.Name).
							Where(refDictItem.Id.Eq(10)).
							IntoDB(),
						"di",
					),
					xDi.DictId.EqCol(refDict.Id),
				).
				IntoDB().
				Take(&dummy),
			wantVars: []any{int64(10), 1},
			want:     "SELECT `dict`.`id`,`di`.`name` FROM `dict` INNER JOIN (SELECT `dict_item`.`dict_id`,`dict_item`.`name` FROM `dict_items` WHERE `dict_item`.`id` = ?) AS `di` ON `di`.`dict_id` = `dict`.`id` LIMIT ?",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ReviewBuildDb(t, tt.db, tt.want, tt.wantVars)
		})
	}
}
