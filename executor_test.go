package rapier

import (
	"context"
	"testing"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

func Test_Executor_Stand(t *testing.T) {
	t.Run("executor", func(t *testing.T) {
		err := refDict.New_Executor(newDb()).
			Debug().
			WithContext(context.Background()).
			Unscoped().
			Session(&gorm.Session{}).
			Clauses(clause.Locking{Strength: "UPDATE"}).
			Table("a").
			Select([]string{"id", "pid", "name"}).
			Distinct().
			Omit("sort").
			OmitExpr(refDict.Sort).
			Where("id = ?", 1).
			Scopes(func(d *gorm.DB) *gorm.DB {
				return d.Where("score > ?", 10)
			}).
			Or("pid = ?", 0).
			Not("is_ping = ?", false).
			Order("created_at").
			Group("name").
			Having("").
			InnerJoins(refDict.Ref_Alias()).
			Joins(refDict.Ref_Alias()).
			Limit(10).
			Offset(2).
			Find(&[]Dict{})
		if err != nil {
			t.Error(err)
		}
	})
}

func Test_Executor_Expr(t *testing.T) {
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
			name: "Expr: table",
			db: refDict.New_Executor(newDb()).
				TableExpr(
					From{
						"a",
						refDict.New_Executor(newDb()).IntoDB(),
					},
				).
				IntoDB().
				Take(&dummy),
			wantVars: []any{1},
			want:     "SELECT * FROM (SELECT * FROM `dict`) AS `a` LIMIT ?",
		},
		{
			name: "Expr: select *",
			db: refDict.New_Executor(newDb()).
				SelectExpr().
				IntoDB().
				Take(&dummy),
			wantVars: []any{1},
			want:     "SELECT * FROM `dict` LIMIT ?",
		},
		{
			name: "Expr: select field",
			db: refDict.New_Executor(newDb()).
				SelectExpr(
					refDict.Id,
					refDict.CreatedAt.UnixTimestamp().As("created_at"),
					refDict.CreatedAt.UnixTimestamp().IfNull(0).As("created_at1"),
				).
				IntoDB().
				Take(&dummy),
			wantVars: []any{int64(0), 1},
			want:     "SELECT `dict`.`id`,UNIX_TIMESTAMP(`dict`.`created_at`) AS `created_at`,IFNULL(UNIX_TIMESTAMP(`dict`.`created_at`),?) AS `created_at1` FROM `dict` LIMIT ?",
		},
		{
			name: "Expr: select * using distinct",
			db: refDict.New_Executor(newDb()).
				DistinctExpr(refDict.Id).
				IntoDB().
				Take(&dummy),
			wantVars: []any{1},
			want:     "SELECT DISTINCT `dict`.`id` FROM `dict` LIMIT ?",
		},
		{
			name: "Expr: order",
			db: refDict.New_Executor(newDb()).
				OrderExpr(refDict.Score).
				IntoDB().
				Take(&dummy),
			wantVars: []any{1},
			want:     "SELECT * FROM `dict` ORDER BY `dict`.`score` LIMIT ?",
		},
		{
			name: "Expr: group",
			db: refDict.New_Executor(newDb()).
				GroupExpr(refDict.Name).
				IntoDB().
				Take(&dummy),
			wantVars: []any{1},
			want:     "SELECT * FROM `dict` GROUP BY `dict`.`name` LIMIT ?",
		},
		{
			name: "Expr: cross join",
			db: refDict.New_Executor(newDb()).
				CrossJoinsExpr(
					&xDt,
					xDt.DictId.EqCol(refDict.Id),
				).
				IntoDB().
				Take(&dummy),
			wantVars: []any{1},
			want:     "SELECT `dict`.`id`,`dict`.`pid`,`dict`.`name`,`dict`.`score`,`dict`.`is_pin`,`dict`.`sort`,`dict`.`created_at` FROM `dict` CROSS JOIN `dict_item` ON `dict_item`.`dict_id` = `dict`.`id` LIMIT ?",
		},
		{
			name: "Expr: cross join X",
			db: refDict.New_Executor(newDb()).
				CrossJoinsXExpr(
					&xDd,
					xDd.Ref_Alias(),
					xDd.Id.EqCol(refDict.Pid),
					xDd.IsPin.Eq(true),
				).
				IntoDB().
				Take(&dummy),
			wantVars: []any{true, 1},
			want:     "SELECT `dict`.`id`,`dict`.`pid`,`dict`.`name`,`dict`.`score`,`dict`.`is_pin`,`dict`.`sort`,`dict`.`created_at` FROM `dict` CROSS JOIN `dict` `dd` ON `dd`.`id` = `dict`.`pid` AND `dd`.`is_pin` = ? LIMIT ?",
		},
		{
			name: "Expr: inner join",
			db: refDict.New_Executor(newDb()).
				InnerJoinsExpr(&xDt, xDt.DictId.EqCol(refDict.Id)).
				IntoDB().
				Take(&dummy),
			wantVars: []any{1},
			want:     "SELECT `dict`.`id`,`dict`.`pid`,`dict`.`name`,`dict`.`score`,`dict`.`is_pin`,`dict`.`sort`,`dict`.`created_at` FROM `dict` INNER JOIN `dict_item` ON `dict_item`.`dict_id` = `dict`.`id` LIMIT ?",
		},
		{
			name: "Expr: inner join X",
			db: refDict.New_Executor(newDb()).
				InnerJoinsXExpr(&xDd, xDd.Ref_Alias(), xDd.Id.EqCol(refDict.Pid), xDd.IsPin.Eq(true)).
				IntoDB().
				Take(&dummy),
			wantVars: []any{true, 1},
			want:     "SELECT `dict`.`id`,`dict`.`pid`,`dict`.`name`,`dict`.`score`,`dict`.`is_pin`,`dict`.`sort`,`dict`.`created_at` FROM `dict` INNER JOIN `dict` `dd` ON `dd`.`id` = `dict`.`pid` AND `dd`.`is_pin` = ? LIMIT ?",
		},
		{
			name: "Expr: left join",
			db: refDict.New_Executor(newDb()).
				LeftJoinsExpr(&xDt, xDt.DictId.EqCol(refDict.Id)).
				IntoDB().
				Take(&dummy),
			wantVars: []any{1},
			want:     "SELECT `dict`.`id`,`dict`.`pid`,`dict`.`name`,`dict`.`score`,`dict`.`is_pin`,`dict`.`sort`,`dict`.`created_at` FROM `dict` LEFT JOIN `dict_item` ON `dict_item`.`dict_id` = `dict`.`id` LIMIT ?",
		},
		{
			name: "Expr: left join X",
			db: refDict.New_Executor(newDb()).
				LeftJoinsXExpr(&xDd, xDd.Ref_Alias(), xDd.Id.EqCol(refDict.Pid), xDd.IsPin.Eq(true)).
				IntoDB().
				Take(&dummy),
			wantVars: []any{true, 1},
			want:     "SELECT `dict`.`id`,`dict`.`pid`,`dict`.`name`,`dict`.`score`,`dict`.`is_pin`,`dict`.`sort`,`dict`.`created_at` FROM `dict` LEFT JOIN `dict` `dd` ON `dd`.`id` = `dict`.`pid` AND `dd`.`is_pin` = ? LIMIT ?",
		},
		{
			name: "Expr: right join",
			db: refDict.New_Executor(newDb()).
				RightJoinsExpr(&xDt, xDt.DictId.EqCol(refDict.Id)).
				IntoDB().
				Take(&dummy),
			wantVars: []any{1},
			want:     "SELECT `dict`.`id`,`dict`.`pid`,`dict`.`name`,`dict`.`score`,`dict`.`is_pin`,`dict`.`sort`,`dict`.`created_at` FROM `dict` RIGHT JOIN `dict_item` ON `dict_item`.`dict_id` = `dict`.`id` LIMIT ?",
		},
		{
			name: "Expr: right join X",
			db: refDict.New_Executor(newDb()).
				RightJoinsXExpr(&xDd, xDd.Ref_Alias(), xDd.Id.EqCol(refDict.Pid), xDd.IsPin.Eq(true)).
				IntoDB().
				Take(&dummy),
			wantVars: []any{true, 1},
			want:     "SELECT `dict`.`id`,`dict`.`pid`,`dict`.`name`,`dict`.`score`,`dict`.`is_pin`,`dict`.`sort`,`dict`.`created_at` FROM `dict` RIGHT JOIN `dict` `dd` ON `dd`.`id` = `dict`.`pid` AND `dd`.`is_pin` = ? LIMIT ?",
		},
		{
			name: "clause: for update",
			db: refDict.New_Executor(newDb()).
				LockingUpdate().
				IntoDB().
				Take(&dummy),
			wantVars: []any{1},
			want:     "SELECT * FROM `dict` LIMIT ? FOR UPDATE",
		},
		{
			name: "clause: for share",
			db: refDict.New_Executor(newDb()).
				LockingShare().
				IntoDB().
				Take(&dummy),
			wantVars: []any{1},
			want:     "SELECT * FROM `dict` LIMIT ? FOR SHARE",
		},
		{
			name: "clause: pagination",
			db: refDict.New_Executor(newDb()).
				Pagination(2, 10).
				IntoDB().
				Find(&dummy),
			wantVars: []any{10, 10},
			want:     "SELECT * FROM `dict` LIMIT ? OFFSET ?",
		},
		{
			name: "clause: Returning",
			db: refDict.New_Executor(newDb()).
				Returning("id", "pid").
				updatesExpr(
					refDict.IsPin.Value(false),
				),
			wantVars: []any{false},
			want:     "UPDATE `dict` SET `is_pin`=? RETURNING `id`,`pid`",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ReviewBuildDb(t, tt.db, tt.want, tt.wantVars)
		})
	}
}

func Test_Executor_SubQuery(t *testing.T) {
	var dummy Dict

	tests := []struct {
		name     string
		db       *gorm.DB
		wantVars []any
		want     string
	}{
		{
			name: "sub query: IntoSubQueryExpr",
			db: refDict.New_Executor(newDb()).
				Where(
					refDict.Id.EqCol(
						refDict.New_Executor(db).
							SelectExpr(refDict.Id).
							Where(refDict.Pid.Eq(100)).
							IntoSubQueryExpr(),
					),
				).
				IntoDB().
				Take(&dummy),
			wantVars: []any{int64(100), 1},
			want:     "SELECT * FROM `dict` WHERE `dict`.`id` = (SELECT `dict`.`id` FROM `dict` WHERE `dict`.`pid` = ?) LIMIT ?",
		},
		{
			name: "sub query: IntoExistExpr",
			db: newDb().Model(&Dict{}).
				Where(
					refDict.New_Executor(newDb()).
						SelectExpr(refDict.Id.Min()).
						IntoExistExpr(),
				).
				Find(&dummy),
			wantVars: nil,
			want:     "SELECT * FROM `dict` WHERE EXISTS(SELECT MIN(`dict`.`id`) FROM `dict`)",
		},
		{
			name: "sub query: IntoNotExistExpr",
			db: newDb().Model(&Dict{}).
				Where(
					refDict.New_Executor(newDb()).
						SelectExpr(refDict.Id.Min()).
						IntoNotExistExpr(),
				).
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

func Test_Executor_Update_SetExpr(t *testing.T) {
	var nullString *string

	tests := []struct {
		name     string
		db       *gorm.DB
		wantVars []any
		want     string
	}{
		{
			name: "updateExpr: value",
			db: refDict.New_Executor(newDb()).
				Where(refDict.Id.Eq(1)).
				updateExpr(refDict.Sort, int(100)),
			wantVars: []any{int(100), int64(1)},
			want:     "UPDATE `dict` SET `sort`=? WHERE `dict`.`id` = ?",
		},
		{
			name: "updateExpr: value gorm.Expr",
			db: refDict.New_Executor(newDb()).
				Where(refDict.Id.Eq(1)).
				updateExpr(refDict.Sort, gorm.Expr("`sort`+?", 100)),
			wantVars: []any{int(100), int64(1)},
			want:     "UPDATE `dict` SET `sort`=`sort`+? WHERE `dict`.`id` = ?",
		},
		{
			name: "updateExpr: value SetExpr",
			db: refDict.New_Executor(newDb()).
				Where(refDict.Id.Eq(1)).
				updateExpr(refDict.Sort, refDict.Score.Add(100)),
			wantVars: []any{float64(100), int64(1)},
			want:     "UPDATE `dict` SET `sort`=`dict`.`score`+? WHERE `dict`.`id` = ?",
		},
		{
			name: "updatesExpr: value SetExpr",
			db: refDict.New_Executor(newDb()).
				Where(refDict.Id.Eq(1)).
				updatesExpr(
					refDict.Name.Value("abc"),
					refDict.Score.Add(10),
					refDict.Sort.ValueAny(gorm.Expr("`sort`+?", 100)),
					refDict.CreatedAt.ValueNull(),
				),
			wantVars: []any{"abc", float64(10), int(100), nil, int64(1)},
			want:     "UPDATE `dict` SET `name`=?,`score`=`dict`.`score`+?,`sort`=`sort`+?,`created_at`=? WHERE `dict`.`id` = ?",
		},
		{
			name: "updateColumnExpr: value",
			db: refDict.New_Executor(newDb()).
				Where(refDict.Id.Eq(1)).
				updateColumnExpr(refDict.Sort, int(100)),
			wantVars: []any{int(100), int64(1)},
			want:     "UPDATE `dict` SET `sort`=? WHERE `dict`.`id` = ?",
		},
		{
			name: "updateColumnExpr: value SetExpr",
			db: refDict.New_Executor(newDb()).
				Where(refDict.Id.Eq(1)).
				updateColumnExpr(refDict.Sort, refDict.Sort.Add(100)),
			wantVars: []any{uint16(100), int64(1)},
			want:     "UPDATE `dict` SET `sort`=`dict`.`sort`+? WHERE `dict`.`id` = ?",
		},
		{
			name: "updateColumnsExpr: value SetExpr",
			db: refDict.New_Executor(newDb()).
				Where(refDict.Id.Eq(1)).
				updateColumnsExpr(
					refDict.Sort.Value(100),
					refDict.Score.Add(10),
					refDict.Name.ValueAny(nullString),
					refDict.CreatedAt.ValueAny(nil),
				),
			wantVars: []any{uint16(100), float64(10), nullString, nil, int64(1)},
			want:     "UPDATE `dict` SET `sort`=?,`score`=`dict`.`score`+?,`name`=?,`created_at`=? WHERE `dict`.`id` = ?",
		},
		{
			name: "updatesExpr: SubQuery",
			db: refDict.New_Executor(newDb()).
				Where(refDict.Id.Eq(1)).
				updatesExpr(
					refDict.Score.SetSubQuery(refDict.New_Executor(newDb()).SelectExpr(refDict.Score).Where(refDict.Id.Eq(2)).IntoDB()),
				),
			wantVars: []any{int64(2), int64(1)},
			want:     "UPDATE `dict` SET `score`=(SELECT `dict`.`score` FROM `dict` WHERE `dict`.`id` = ?) WHERE `dict`.`id` = ?",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ReviewBuildDb(t, tt.db, tt.want, tt.wantVars)
		})
	}
}

func Test_Executor_Attrs(t *testing.T) {
	t.Run("attr", func(t *testing.T) {
		wantName := "aaaa"
		wantSort := uint16(1111)
		got1, err := refDict.New_Executor(newDb()).
			Where(refDict.Id.Eq(1)).
			Attrs(&Dict{
				Name: wantName,
				Sort: wantSort,
			}).
			FirstOrCreate()
		if err != nil {
			t.Fatal(err)
		}
		if got1.Name != wantName || got1.Sort != wantSort {
			t.Errorf("name want: %v, got: %v,  sort want: %v got: %v", wantName, got1.Name, wantSort, got1.Sort)
		}

		got2, err := refDict.New_Executor(newDb()).
			Where(refDict.Id.Eq(1)).
			Attrs(&Dict{
				Name: wantName,
				Sort: wantSort,
			}).
			FirstOrInit()
		if err != nil {
			t.Fatal(err)
		}
		if got2.Name != wantName || got2.Sort != wantSort {
			t.Errorf("name want: %v, got: %v,  sort want: %v got: %v", wantName, got2.Name, wantSort, got2.Sort)
		}
	})
	t.Run("attr expr", func(t *testing.T) {
		wantName := "bbbb"
		wantSort := uint16(2222)

		got1, err := refDict.New_Executor(newDb()).
			Where(refDict.Id.Eq(1)).
			AttrsExpr(
				refDict.Name.Value(wantName),
				refDict.Sort.Value(wantSort),
			).
			FirstOrCreate()
		if err != nil {
			t.Fatal(err)
		}
		if got1.Name != wantName || got1.Sort != wantSort {
			t.Errorf("name want: %v, got: %v,  sort want: %v got: %v", wantName, got1.Name, wantSort, got1.Sort)
		}

		got2, err := refDict.New_Executor(newDb()).
			Where(refDict.Id.Eq(1)).
			AttrsExpr(
				refDict.Name.Value(wantName),
				refDict.Sort.Value(wantSort),
			).
			FirstOrInit()
		if err != nil {
			t.Fatal(err)
		}
		if got2.Name != wantName || got2.Sort != wantSort {
			t.Errorf("name want: %v, got: %v,  sort want: %v got: %v", wantName, got2.Name, wantSort, got2.Sort)
		}
	})
}

func Test_Executor_Assign(t *testing.T) {
	t.Run("assign", func(t *testing.T) {
		wantName := "aaaa"
		wantSort := uint16(1111)
		got1, err := refDict.New_Executor(newDb()).
			Where(refDict.Id.Eq(1)).
			Assign(&Dict{
				Name: wantName,
				Sort: wantSort,
			}).
			FirstOrCreate()
		if err != nil {
			t.Fatal(err)
		}
		if got1.Name != wantName || got1.Sort != wantSort {
			t.Errorf("name want: %v, got: %v,  sort want: %v got: %v", wantName, got1.Name, wantSort, got1.Sort)
		}

		got2, err := refDict.New_Executor(newDb()).
			Where(refDict.Id.Eq(1)).
			Assign(&Dict{
				Name: wantName,
				Sort: wantSort,
			}).
			FirstOrInit()
		if err != nil {
			t.Fatal(err)
		}
		if got2.Name != wantName || got2.Sort != wantSort {
			t.Errorf("name want: %v, got: %v,  sort want: %v got: %v", wantName, got1.Name, wantSort, got1.Sort)
		}
	})
	t.Run("assign expr", func(t *testing.T) {
		wantName := "bbbb"
		wantSort := uint16(2222)
		got1, err := refDict.New_Executor(newDb()).
			Where(refDict.Id.Eq(1)).
			AssignExpr(
				refDict.Name.Value(wantName),
				refDict.Sort.Value(wantSort),
			).
			FirstOrCreate()
		if err != nil {
			t.Fatal(err)
		}
		if got1.Name != wantName || got1.Sort != wantSort {
			t.Errorf("name want: %v, got: %v,  sort want: %v got: %v", wantName, got1.Name, wantSort, got1.Sort)
		}
		got2, err := refDict.New_Executor(newDb()).
			Where(refDict.Id.Eq(1)).
			AssignExpr(
				refDict.Name.Value(wantName),
				refDict.Sort.valueEq(wantSort),
			).
			FirstOrInit()
		if err != nil {
			t.Fatal(err)
		}
		if got2.Name != wantName || got2.Sort != wantSort {
			t.Errorf("name want: %v, got: %v,  sort want: %v got: %v", wantName, got1.Name, wantSort, got1.Sort)
		}
	})
}

func Test_Executor_Preload(t *testing.T) {
	var td TestDict

	_ = refDict.New_Executor(newDbWithLog()).
		Preload("DictItem").Take(&td)
}
