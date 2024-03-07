package rapier

import (
	"testing"

	"gorm.io/gorm"
)

func Test_TableImplSchemaTabler(t *testing.T) {
	tableName := "tt"
	tb := Table(tableName)
	if gotTableName := tb.TableName(); gotTableName != tableName {
		t.Errorf("want: %s, got: %s", tableName, gotTableName)
	}
}

func Test_Table(t *testing.T) {
	var dummy Dict

	tests := []struct {
		name     string
		db       *gorm.DB
		wantVars []any
		want     string
	}{
		{
			name: "empty table",
			db: newDb().Model(&Dict{}).
				Scopes(
					TableExpr(),
				).
				Take(&dummy),
			wantVars: []any{1},
			want:     "SELECT * FROM `dict` LIMIT ?",
		},
		{
			name: "single table",
			db: newDb().Model(&Dict{}).
				Scopes(
					TableExpr(
						From{
							"a",
							newDb().
								Model(&Dict{}),
						},
					),
				).
				Take(&dummy),
			wantVars: []any{1},
			want:     "SELECT * FROM (SELECT * FROM `dict`) AS `a` LIMIT ?",
		},
		{
			name: "multi table",
			db: newDb().Model(&Dict{}).
				Scopes(
					TableExpr(
						From{
							"a",
							newDb().
								Model(&Dict{}),
						},
						From{
							"b",
							newDb().
								Model(&Dict{}),
						},
					),
				).
				Take(&dummy),
			wantVars: []any{1},
			want:     "SELECT * FROM (SELECT * FROM `dict`) AS `a`, (SELECT * FROM `dict`) AS `b` LIMIT ?",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ReviewBuildDb(t, tt.db, tt.want, tt.wantVars)
		})
	}
}

func Test_Select(t *testing.T) {
	var dummy Dict

	tests := []struct {
		name     string
		db       *gorm.DB
		wantVars []any
		want     string
	}{
		{
			name: "select *",
			db: newDb().Model(&Dict{}).
				Scopes(
					SelectExpr(),
				).
				Take(&dummy),
			wantVars: []any{1},
			want:     "SELECT * FROM `dict` LIMIT ?",
		},
		{
			name: "select field",
			db: newDb().Model(&Dict{}).
				Scopes(
					SelectExpr(
						refDict.Id,
						refDict.CreatedAt.UnixTimestamp().As("created_at"),
						refDict.CreatedAt.UnixTimestamp().IfNull(0).As("created_at1"),
					),
				).
				Take(&dummy),
			wantVars: []any{int64(0), 1},
			want:     "SELECT `dict`.`id`,UNIX_TIMESTAMP(`dict`.`created_at`) AS `created_at`,IFNULL(UNIX_TIMESTAMP(`dict`.`created_at`),?) AS `created_at1` FROM `dict` LIMIT ?",
		},
		{
			name: "select field where",
			db: newDb().Model(&Dict{}).
				Scopes(
					SelectExpr(refDict.Id, refDict.Score),
				).
				Where(refDict.Name.Eq(""), refDict.IsPin.Is(true)).
				Take(&dummy),
			wantVars: []any{"", true, 1},
			want:     "SELECT `dict`.`id`,`dict`.`score` FROM `dict` WHERE `dict`.`name` = ? AND `dict`.`is_pin` = ? LIMIT ?",
		},
		{
			name: "select 1",
			db: newDb().Model(&Dict{}).
				Scopes(
					SelectExpr(One),
				).
				Take(&dummy),
			wantVars: []any{1},
			want:     "SELECT 1 FROM `dict` LIMIT ?",
		},
		{
			name: "select COUNT(1)",
			db: newDb().Model(&Dict{}).
				Scopes(
					SelectExpr(One.Count()),
				).
				Take(&dummy),
			wantVars: []any{1},
			want:     "SELECT COUNT(1) FROM `dict` LIMIT ?",
		},
		{
			name: "select COUNT(*)",
			db: newDb().Model(&Dict{}).
				Scopes(
					SelectExpr(Star.Count()),
				).
				Take(&dummy),
			wantVars: []any{1},
			want:     "SELECT COUNT(*) FROM `dict` LIMIT ?",
		},
		{
			name: "select AVG(field)",
			db: newDb().Model(&Dict{}).
				Scopes(
					SelectExpr(refDict.Score.Avg()),
				).
				Take(&dummy),
			wantVars: []any{1},
			want:     "SELECT AVG(`dict`.`score`) FROM `dict` LIMIT ?",
		},
		{
			name: "update with select field",
			db: newDb().Model(&Dict{}).
				Scopes(
					SelectExpr(
						refDict.Score,
						refDict.IsPin,
					),
				).
				Where(refDict.Id.Eq(100)).
				Updates(&Dict{
					Score: 100,
					IsPin: true,
				}),
			wantVars: []any{float64(100), true, int64(100)},
			want:     "UPDATE `dict` SET `score`=?,`is_pin`=? WHERE `dict`.`id` = ?",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ReviewBuildDb(t, tt.db, tt.want, tt.wantVars)
		})
	}
}

func Test_Omit(t *testing.T) {
	var dummy Dict

	tests := []struct {
		name     string
		db       *gorm.DB
		wantVars []any
		want     string
	}{
		{
			name: "select *",
			db: newDb().Model(&Dict{}).
				Scopes(
					OmitExpr(),
				).
				Take(&dummy),
			wantVars: []any{1},
			want:     "SELECT * FROM `dict` LIMIT ?",
		},
		{
			name: "omit field",
			db: newDb().Model(&Dict{}).
				Scopes(
					OmitExpr(
						refDict.CreatedAt,
					),
				).
				Take(&dummy),
			wantVars: []any{1},
			want:     "SELECT `dict`.`id`,`dict`.`pid`,`dict`.`name`,`dict`.`score`,`dict`.`is_pin`,`dict`.`sort` FROM `dict` LIMIT ?",
		},
		{
			name: "omit more fields",
			db: newDb().Model(&Dict{}).
				Scopes(
					OmitExpr(
						refDict.Score,
						refDict.CreatedAt,
					),
				).
				Take(&dummy),
			wantVars: []any{1},
			want:     "SELECT `dict`.`id`,`dict`.`pid`,`dict`.`name`,`dict`.`is_pin`,`dict`.`sort` FROM `dict` LIMIT ?",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ReviewBuildDb(t, tt.db, tt.want, tt.wantVars)
		})
	}
}

func Test_Distinct(t *testing.T) {
	tests := []struct {
		name     string
		db       *gorm.DB
		wantVars []any
		want     string
	}{
		{
			name: "select * using distinct",
			db: newDb().Model(&Dict{}).
				Scopes(
					DistinctExpr(),
					SelectExpr(refDict.Id),
				).
				Take(&Dict{}),
			wantVars: []any{1},
			want:     "SELECT DISTINCT `dict`.`id` FROM `dict` LIMIT ?",
		},
		{
			name: "distinct field",
			db: newDb().Model(&Dict{}).
				Scopes(DistinctExpr(refDict.Id)).
				Take(&Dict{}),
			wantVars: []any{1},
			want:     "SELECT DISTINCT `dict`.`id` FROM `dict` LIMIT ?",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ReviewBuildDb(t, tt.db, tt.want, tt.wantVars)
		})
	}
}

func Test_Order(t *testing.T) {
	var dummy Dict

	tests := []struct {
		name     string
		db       *gorm.DB
		wantVars []any
		want     string
	}{
		{
			name: "empty field",
			db: newDb().Model(&Dict{}).
				Scopes(
					OrderExpr(),
				).
				Take(&dummy),
			wantVars: []any{1},
			want:     "SELECT * FROM `dict` LIMIT ?",
		},
		{
			name: "empty order",
			db: newDb().Model(&Dict{}).
				Scopes(
					OrderExpr(refDict.Score),
				).
				Take(&dummy),
			wantVars: []any{1},
			want:     "SELECT * FROM `dict` ORDER BY `dict`.`score` LIMIT ?",
		},
		{
			name: "desc",
			db: newDb().Model(&Dict{}).
				Scopes(
					OrderExpr(refDict.Score.Desc()),
				).
				Take(&dummy),
			wantVars: []any{1},
			want:     "SELECT * FROM `dict` ORDER BY `dict`.`score` DESC LIMIT ?",
		},
		{
			name: "multiple desc",
			db: newDb().Model(&Dict{}).
				Scopes(
					OrderExpr(refDict.Score.Desc(), refDict.Name),
				).
				Take(&dummy),
			wantVars: []any{1},
			want:     "SELECT * FROM `dict` ORDER BY `dict`.`score` DESC,`dict`.`name` LIMIT ?",
		},
		{
			name: "asc",
			db: newDb().Model(&Dict{}).
				Scopes(
					OrderExpr(refDict.Score.Asc()),
				).
				Take(&dummy),
			wantVars: []any{1},
			want:     "SELECT * FROM `dict` ORDER BY `dict`.`score` ASC LIMIT ?",
		},
		{
			name: "multiple asc and desc",
			db: newDb().Model(&Dict{}).
				Scopes(
					OrderExpr(refDict.Score.Desc(), refDict.Name.Asc()),
				).
				Take(&dummy),
			wantVars: []any{1},
			want:     "SELECT * FROM `dict` ORDER BY `dict`.`score` DESC,`dict`.`name` ASC LIMIT ?",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ReviewBuildDb(t, tt.db, tt.want, tt.wantVars)
		})
	}
}

func Test_Group(t *testing.T) {
	var dummy Dict

	tests := []struct {
		name     string
		db       *gorm.DB
		wantVars []any
		want     string
	}{
		{
			name: "",
			db: newDb().Model(&Dict{}).
				Scopes(
					GroupExpr(),
				).
				Take(&dummy),
			wantVars: []any{1},
			want:     "SELECT * FROM `dict` LIMIT ?",
		},
		{
			name: "",
			db: newDb().Model(&Dict{}).
				Scopes(
					GroupExpr(refDict.Name),
				).
				Take(&dummy),
			wantVars: []any{1},
			want:     "SELECT * FROM `dict` GROUP BY `dict`.`name` LIMIT ?",
		},
		{
			name: "",
			db: newDb().Model(&Dict{}).
				Scopes(
					SelectExpr(refDict.Score.Sum()),
					GroupExpr(refDict.Name),
				).
				Having(refDict.Score.Sum().Gt(100)).
				Take(&dummy),
			wantVars: []any{float64(100), 1},
			want:     "SELECT SUM(`dict`.`score`) FROM `dict` GROUP BY `dict`.`name` HAVING SUM(`dict`.`score`) > ? LIMIT ?",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ReviewBuildDb(t, tt.db, tt.want, tt.wantVars)
		})
	}
}

func Test_Locking(t *testing.T) {
	var dummy Dict

	tests := []struct {
		name     string
		db       *gorm.DB
		wantVars []any
		want     string
	}{
		{
			name: "",
			db: newDb().Model(&Dict{}).
				Scopes(
					GroupExpr(),
					LockingUpdate(),
				).
				Take(&dummy),
			wantVars: []any{1},
			want:     "SELECT * FROM `dict` LIMIT ? FOR UPDATE",
		},
		{
			name: "",
			db: newDb().Model(&Dict{}).
				Scopes(
					GroupExpr(),
					LockingShare(),
				).
				Take(&dummy),
			wantVars: []any{1},
			want:     "SELECT * FROM `dict` LIMIT ? FOR SHARE",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ReviewBuildDb(t, tt.db, tt.want, tt.wantVars)
		})
	}
}
