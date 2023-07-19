package assist

import (
	"testing"

	"gorm.io/gorm"
)

func Test_Joins(t *testing.T) {
	var dummy Dict

	xDd := xDict.As("dd")
	xDi := xDict.As("di")

	tests := []struct {
		name     string
		db       *gorm.DB
		wantVars []any
		want     string
	}{
		{
			name: "inner join - empty conds",
			db: newDb().Model( &Dict{}).
				Scopes(
					InnerJoinsExpr(xDd.X_TableName()),
				).
				Take(&dummy),
			wantVars: nil,
			want:     "SELECT * FROM `dict` LIMIT 1",
		},
		{
			name: "cross join",
			db: newDb().Model( &Dict{}).
				Scopes(
					CrossJoinsExpr(xDd.X_TableName(), xDd.Id.EqCol(xDict.Pid), xDd.IsPin.Eq(true)),
				).
				Take(&dummy),
			wantVars: []any{true},
			want:     "SELECT `dict`.`id`,`dict`.`pid`,`dict`.`name`,`dict`.`score`,`dict`.`is_pin`,`dict`.`sort`,`dict`.`created_at` FROM `dict` CROSS JOIN `dd` ON `dd`.`id` = `dict`.`pid` AND `dd`.`is_pin` = ? LIMIT 1",
		},
		{
			name: "inner join",
			db: newDb().Model( &Dict{}).
				Scopes(
					InnerJoinsExpr(xDd.X_TableName(), xDd.Id.EqCol(xDict.Pid), xDd.IsPin.Eq(true)),
				).
				Take(&dummy),
			wantVars: []any{true},
			want:     "SELECT `dict`.`id`,`dict`.`pid`,`dict`.`name`,`dict`.`score`,`dict`.`is_pin`,`dict`.`sort`,`dict`.`created_at` FROM `dict` INNER JOIN `dd` ON `dd`.`id` = `dict`.`pid` AND `dd`.`is_pin` = ? LIMIT 1",
		},
		{
			name: "left join",
			db: newDb().Model( &Dict{}).
				Scopes(
					LeftJoinsExpr(xDd.X_TableName(), xDd.Id.EqCol(xDict.Pid), xDd.IsPin.Eq(true)),
				).
				Take(&dummy),
			wantVars: []any{true},
			want:     "SELECT `dict`.`id`,`dict`.`pid`,`dict`.`name`,`dict`.`score`,`dict`.`is_pin`,`dict`.`sort`,`dict`.`created_at` FROM `dict` LEFT JOIN `dd` ON `dd`.`id` = `dict`.`pid` AND `dd`.`is_pin` = ? LIMIT 1",
		},
		{
			name: "right join",
			db: newDb().Model( &Dict{}).
				Scopes(
					RightJoinsExpr(xDd.X_TableName(), xDd.Id.EqCol(xDict.Pid), xDd.IsPin.Eq(true)),
				).
				Take(&dummy),
			wantVars: []any{true},
			want:     "SELECT `dict`.`id`,`dict`.`pid`,`dict`.`name`,`dict`.`score`,`dict`.`is_pin`,`dict`.`sort`,`dict`.`created_at` FROM `dict` RIGHT JOIN `dd` ON `dd`.`id` = `dict`.`pid` AND `dd`.`is_pin` = ? LIMIT 1",
		},
		{
			name: "inner join - multiple",
			db: newDb().Model( &Dict{}).
				Scopes(
					InnerJoinsExpr(xDd.X_TableName(), xDd.Id.EqCol(xDict.Pid)),
					InnerJoinsExpr(xDi.X_TableName(), xDi.IsPin.Eq(true)),
				).
				Take(&dummy),
			wantVars: []any{true},
			want:     "SELECT `dict`.`id`,`dict`.`pid`,`dict`.`name`,`dict`.`score`,`dict`.`is_pin`,`dict`.`sort`,`dict`.`created_at` FROM `dict` INNER JOIN `dd` ON `dd`.`id` = `dict`.`pid` INNER JOIN `di` ON `di`.`is_pin` = ? LIMIT 1",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			CheckBuildExprSql(t, tt.db, tt.want, tt.wantVars)
		})
	}
}
