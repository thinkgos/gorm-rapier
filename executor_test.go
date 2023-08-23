package assist

import (
	"context"
	"testing"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

func Test_Executor_Stand(t *testing.T) {
	t.Run("executor", func(t *testing.T) {
		err := xDict.New_Executor(newDb()).
			Debug().
			WithContext(context.Background()).
			Unscoped().
			Session(&gorm.Session{}).
			Clauses(clause.Locking{Strength: "UPDATE"}).
			Table("a").
			Select([]string{"id", "pid", "name"}).
			Distinct().
			Omit("sort").
			OmitExpr(xDict.Sort).
			Where("id = ?", 1).
			Scopes(func(d *gorm.DB) *gorm.DB {
				return d.Where("score > ?", 10)
			}).
			Or("pid = ?", 0).
			Not("is_ping = ?", false).
			Order("created_at").
			Group("name").
			Having("").
			InnerJoins(xDict.X_Alias()).
			Joins(xDict.X_Alias()).
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

	xDd := xDict.As("dd")
	xDt := xDictItem

	tests := []struct {
		name     string
		db       *gorm.DB
		wantVars []any
		want     string
	}{
		{
			name: "Expr: table",
			db: xDict.New_Executor(newDb()).
				TableExpr(
					From{
						"a",
						xDict.New_Executor(newDb()).IntoDB(),
					},
				).
				IntoDB().
				Take(&dummy),
			wantVars: nil,
			want:     "SELECT * FROM (SELECT * FROM `dict`) AS `a` LIMIT 1",
		},
		{
			name: "Expr: select *",
			db: xDict.New_Executor(newDb()).
				SelectExpr().
				IntoDB().
				Take(&dummy),
			wantVars: nil,
			want:     "SELECT * FROM `dict` LIMIT 1",
		},
		{
			name: "Expr: select field",
			db: xDict.New_Executor(newDb()).
				SelectExpr(
					xDict.Id,
					xDict.CreatedAt.UnixTimestamp().As("created_at"),
					xDict.CreatedAt.UnixTimestamp().IfNull(0).As("created_at1"),
				).
				IntoDB().
				Take(&dummy),
			wantVars: []any{int64(0)},
			want:     "SELECT `dict`.`id`,UNIX_TIMESTAMP(`dict`.`created_at`) AS `created_at`,IFNULL(UNIX_TIMESTAMP(`dict`.`created_at`),?) AS `created_at1` FROM `dict` LIMIT 1",
		},
		{
			name: "Expr: select * using distinct",
			db: xDict.New_Executor(newDb()).
				DistinctExpr(xDict.Id).
				IntoDB().
				Take(&dummy),
			wantVars: nil,
			want:     "SELECT DISTINCT `dict`.`id` FROM `dict` LIMIT 1",
		},
		{
			name: "Expr: order",
			db: xDict.New_Executor(newDb()).
				OrderExpr(xDict.Score).
				IntoDB().
				Take(&dummy),
			wantVars: nil,
			want:     "SELECT * FROM `dict` ORDER BY `dict`.`score` LIMIT 1",
		},
		{
			name: "Expr: group",
			db: xDict.New_Executor(newDb()).
				GroupExpr(xDict.Name).
				IntoDB().
				Take(&dummy),
			wantVars: nil,
			want:     "SELECT * FROM `dict` GROUP BY `dict`.`name` LIMIT 1",
		},
		{
			name: "Expr: cross join",
			db: xDict.New_Executor(newDb()).
				CrossJoinsExpr(
					&xDt,
					xDt.DictId.EqCol(xDict.Id),
				).
				IntoDB().
				Take(&dummy),
			wantVars: nil,
			want:     "SELECT `dict`.`id`,`dict`.`pid`,`dict`.`name`,`dict`.`score`,`dict`.`is_pin`,`dict`.`sort`,`dict`.`created_at` FROM `dict` CROSS JOIN `dict_item` ON `dict_item`.`dict_id` = `dict`.`id` LIMIT 1",
		},
		{
			name: "Expr: cross join X",
			db: xDict.New_Executor(newDb()).
				CrossJoinsXExpr(
					&xDd,
					xDd.X_Alias(),
					xDd.Id.EqCol(xDict.Pid),
					xDd.IsPin.Eq(true),
				).
				IntoDB().
				Take(&dummy),
			wantVars: []any{true},
			want:     "SELECT `dict`.`id`,`dict`.`pid`,`dict`.`name`,`dict`.`score`,`dict`.`is_pin`,`dict`.`sort`,`dict`.`created_at` FROM `dict` CROSS JOIN `dict` `dd` ON `dd`.`id` = `dict`.`pid` AND `dd`.`is_pin` = ? LIMIT 1",
		},
		{
			name: "Expr: inner join",
			db: xDict.New_Executor(newDb()).
				InnerJoinsExpr(&xDt, xDt.DictId.EqCol(xDict.Id)).
				IntoDB().
				Take(&dummy),
			wantVars: nil,
			want:     "SELECT `dict`.`id`,`dict`.`pid`,`dict`.`name`,`dict`.`score`,`dict`.`is_pin`,`dict`.`sort`,`dict`.`created_at` FROM `dict` INNER JOIN `dict_item` ON `dict_item`.`dict_id` = `dict`.`id` LIMIT 1",
		},
		{
			name: "Expr: inner join X",
			db: xDict.New_Executor(newDb()).
				InnerJoinsXExpr(&xDd, xDd.X_Alias(), xDd.Id.EqCol(xDict.Pid), xDd.IsPin.Eq(true)).
				IntoDB().
				Take(&dummy),
			wantVars: []any{true},
			want:     "SELECT `dict`.`id`,`dict`.`pid`,`dict`.`name`,`dict`.`score`,`dict`.`is_pin`,`dict`.`sort`,`dict`.`created_at` FROM `dict` INNER JOIN `dict` `dd` ON `dd`.`id` = `dict`.`pid` AND `dd`.`is_pin` = ? LIMIT 1",
		},
		{
			name: "Expr: left join",
			db: xDict.New_Executor(newDb()).
				LeftJoinsExpr(&xDt, xDt.DictId.EqCol(xDict.Id)).
				IntoDB().
				Take(&dummy),
			wantVars: nil,
			want:     "SELECT `dict`.`id`,`dict`.`pid`,`dict`.`name`,`dict`.`score`,`dict`.`is_pin`,`dict`.`sort`,`dict`.`created_at` FROM `dict` LEFT JOIN `dict_item` ON `dict_item`.`dict_id` = `dict`.`id` LIMIT 1",
		},
		{
			name: "Expr: left join X",
			db: xDict.New_Executor(newDb()).
				LeftJoinsXExpr(&xDd, xDd.X_Alias(), xDd.Id.EqCol(xDict.Pid), xDd.IsPin.Eq(true)).
				IntoDB().
				Take(&dummy),
			wantVars: []any{true},
			want:     "SELECT `dict`.`id`,`dict`.`pid`,`dict`.`name`,`dict`.`score`,`dict`.`is_pin`,`dict`.`sort`,`dict`.`created_at` FROM `dict` LEFT JOIN `dict` `dd` ON `dd`.`id` = `dict`.`pid` AND `dd`.`is_pin` = ? LIMIT 1",
		},
		{
			name: "Expr: right join",
			db: xDict.New_Executor(newDb()).
				RightJoinsExpr(&xDt, xDt.DictId.EqCol(xDict.Id)).
				IntoDB().
				Take(&dummy),
			wantVars: nil,
			want:     "SELECT `dict`.`id`,`dict`.`pid`,`dict`.`name`,`dict`.`score`,`dict`.`is_pin`,`dict`.`sort`,`dict`.`created_at` FROM `dict` RIGHT JOIN `dict_item` ON `dict_item`.`dict_id` = `dict`.`id` LIMIT 1",
		},
		{
			name: "Expr: right join X",
			db: xDict.New_Executor(newDb()).
				RightJoinsXExpr(&xDd, xDd.X_Alias(), xDd.Id.EqCol(xDict.Pid), xDd.IsPin.Eq(true)).
				IntoDB().
				Take(&dummy),
			wantVars: []any{true},
			want:     "SELECT `dict`.`id`,`dict`.`pid`,`dict`.`name`,`dict`.`score`,`dict`.`is_pin`,`dict`.`sort`,`dict`.`created_at` FROM `dict` RIGHT JOIN `dict` `dd` ON `dd`.`id` = `dict`.`pid` AND `dd`.`is_pin` = ? LIMIT 1",
		},
		{
			name: "clause: for update",
			db: xDict.New_Executor(newDb()).
				LockingUpdate().
				IntoDB().
				Take(&dummy),
			wantVars: nil,
			want:     "SELECT * FROM `dict` LIMIT 1 FOR UPDATE",
		},
		{
			name: "clause: for share",
			db: xDict.New_Executor(newDb()).
				LockingShare().
				IntoDB().
				Take(&dummy),
			wantVars: nil,
			want:     "SELECT * FROM `dict` LIMIT 1 FOR SHARE",
		},
		{
			name: "clause: pagination",
			db: xDict.New_Executor(newDb()).
				Pagination(2, 10).
				IntoDB().
				Find(&dummy),
			wantVars: nil,
			want:     "SELECT * FROM `dict` LIMIT 10 OFFSET 10",
		},
		{
			name: "clause: Returning",
			db: xDict.New_Executor(newDb()).
				Returning("id", "pid").
				updatesExpr(
					xDict.IsPin.Value(false),
				),
			wantVars: []any{false},
			want:     "UPDATE `dict` SET `is_pin`=? RETURNING `id`,`pid`",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			CheckBuildExprSql(t, tt.db, tt.want, tt.wantVars)
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
			db: xDict.New_Executor(newDb()).
				Where(
					xDict.Id.EqCol(
						xDict.New_Executor(db).
							SelectExpr(xDict.Id).
							Where(xDict.Pid.Eq(100)).
							IntoSubQueryExpr(),
					),
				).
				IntoDB().
				Take(&dummy),
			wantVars: []any{int64(100)},
			want:     "SELECT * FROM `dict` WHERE `dict`.`id` = (SELECT `dict`.`id` FROM `dict` WHERE `dict`.`pid` = ?) LIMIT 1",
		},
		{
			name: "sub query: IntoExistExpr",
			db: newDb().Model(&Dict{}).
				Where(
					xDict.New_Executor(newDb()).
						SelectExpr(xDict.Id.Min()).
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
					xDict.New_Executor(newDb()).
						SelectExpr(xDict.Id.Min()).
						IntoNotExistExpr(),
				).
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
			db: xDict.New_Executor(newDb()).
				Where(xDict.Id.Eq(1)).
				updateExpr(xDict.Sort, int(100)),
			wantVars: []any{int(100), int64(1)},
			want:     "UPDATE `dict` SET `sort`=? WHERE `dict`.`id` = ?",
		},
		{
			name: "updateExpr: value gorm.Expr",
			db: xDict.New_Executor(newDb()).
				Where(xDict.Id.Eq(1)).
				updateExpr(xDict.Sort, gorm.Expr("`sort`+?", 100)),
			wantVars: []any{int(100), int64(1)},
			want:     "UPDATE `dict` SET `sort`=`sort`+? WHERE `dict`.`id` = ?",
		},
		{
			name: "updateExpr: value SetExpr",
			db: xDict.New_Executor(newDb()).
				Where(xDict.Id.Eq(1)).
				updateExpr(xDict.Sort, xDict.Score.Add(100)),
			wantVars: []any{float64(100), int64(1)},
			want:     "UPDATE `dict` SET `sort`=`dict`.`score`+? WHERE `dict`.`id` = ?",
		},
		{
			name: "updatesExpr: value SetExpr",
			db: xDict.New_Executor(newDb()).
				Where(xDict.Id.Eq(1)).
				updatesExpr(
					xDict.Name.Value("abc"),
					xDict.Score.Add(10),
					xDict.Sort.ValueAny(gorm.Expr("`sort`+?", 100)),
					xDict.CreatedAt.ValueNull(),
				),
			wantVars: []any{"abc", float64(10), int(100), nil, int64(1)},
			want:     "UPDATE `dict` SET `name`=?,`score`=`dict`.`score`+?,`sort`=`sort`+?,`created_at`=? WHERE `dict`.`id` = ?",
		},
		{
			name: "updateColumnExpr: value",
			db: xDict.New_Executor(newDb()).
				Where(xDict.Id.Eq(1)).
				updateColumnExpr(xDict.Sort, int(100)),
			wantVars: []any{int(100), int64(1)},
			want:     "UPDATE `dict` SET `sort`=? WHERE `dict`.`id` = ?",
		},
		{
			name: "updateColumnExpr: value SetExpr",
			db: xDict.New_Executor(newDb()).
				Where(xDict.Id.Eq(1)).
				updateColumnExpr(xDict.Sort, xDict.Sort.Add(100)),
			wantVars: []any{uint16(100), int64(1)},
			want:     "UPDATE `dict` SET `sort`=`dict`.`sort`+? WHERE `dict`.`id` = ?",
		},
		{
			name: "updateColumnsExpr: value SetExpr",
			db: xDict.New_Executor(newDb()).
				Where(xDict.Id.Eq(1)).
				updateColumnsExpr(
					xDict.Sort.Value(100),
					xDict.Score.Add(10),
					xDict.Name.ValueAny(nullString),
					xDict.CreatedAt.ValueAny(nil),
				),
			wantVars: []any{uint16(100), float64(10), nullString, nil, int64(1)},
			want:     "UPDATE `dict` SET `sort`=?,`score`=`dict`.`score`+?,`name`=?,`created_at`=? WHERE `dict`.`id` = ?",
		},
		{
			name: "updatesExpr: SubQuery",
			db: xDict.New_Executor(newDb()).
				Where(xDict.Id.Eq(1)).
				updatesExpr(
					xDict.Score.SetSubQuery(xDict.New_Executor(newDb()).SelectExpr(xDict.Score).Where(xDict.Id.Eq(2)).IntoDB()),
				),
			wantVars: []any{int64(2), int64(1)},
			want:     "UPDATE `dict` SET `score`=(SELECT `dict`.`score` FROM `dict` WHERE `dict`.`id` = ?) WHERE `dict`.`id` = ?",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			CheckBuildExprSql(t, tt.db, tt.want, tt.wantVars)
		})
	}
}

func Test_Executor_Attrs(t *testing.T) {
	t.Run("attr", func(t *testing.T) {
		wantName := "aaaa"
		wantSort := uint16(1111)
		got1, err := xDict.New_Executor(newDb()).
			Where(xDict.Id.Eq(1)).
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

		got2, err := xDict.New_Executor(newDb()).
			Where(xDict.Id.Eq(1)).
			Attrs(&Dict{
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
	t.Run("attr expr", func(t *testing.T) {
		wantName := "bbbb"
		wantSort := uint16(2222)

		got1, err := xDict.New_Executor(newDb()).
			Where(xDict.Id.Eq(1)).
			AttrsExpr(
				xDict.Name.Value(wantName),
				xDict.Sort.Value(wantSort),
			).
			FirstOrCreate()
		if err != nil {
			t.Fatal(err)
		}
		if got1.Name != wantName || got1.Sort != wantSort {
			t.Errorf("name want: %v, got: %v,  sort want: %v got: %v", wantName, got1.Name, wantSort, got1.Sort)
		}

		got2, err := xDict.New_Executor(newDb()).
			Where(xDict.Id.Eq(1)).
			AttrsExpr(
				xDict.Name.Value(wantName),
				xDict.Sort.Value(wantSort),
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

func Test_Executor_Assign(t *testing.T) {
	t.Run("assign", func(t *testing.T) {
		wantName := "aaaa"
		wantSort := uint16(1111)
		got1, err := xDict.New_Executor(newDb()).
			Where(xDict.Id.Eq(1)).
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

		got2, err := xDict.New_Executor(newDb()).
			Where(xDict.Id.Eq(1)).
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
		got1, err := xDict.New_Executor(newDb()).
			Where(xDict.Id.Eq(1)).
			AssignExpr(
				xDict.Name.Value(wantName),
				xDict.Sort.Value(wantSort),
			).
			FirstOrCreate()
		if err != nil {
			t.Fatal(err)
		}
		if got1.Name != wantName || got1.Sort != wantSort {
			t.Errorf("name want: %v, got: %v,  sort want: %v got: %v", wantName, got1.Name, wantSort, got1.Sort)
		}
		got2, err := xDict.New_Executor(newDb()).
			Where(xDict.Id.Eq(1)).
			AssignExpr(
				xDict.Name.Value(wantName),
				xDict.Sort.valueEq(wantSort),
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
