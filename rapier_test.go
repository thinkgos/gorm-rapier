package rapier

import (
	"reflect"
	"testing"
	"time"

	"gorm.io/gorm"
)

var xDict = New_X_Dict("dict")
var xDictItem = New_X_DictItem("dict_item")

type Dict struct {
	Id        int64
	Pid       int64
	Name      string
	Score     float64
	IsPin     bool
	Sort      uint16
	CreatedAt time.Time
}

func (*Dict) TableName() string {
	return "dict"
}

type Dict_Active struct {
	// private fields
	xTableName string

	ALL       Asterisk
	Id        Int64
	Pid       Int64
	Name      String
	Score     Float64
	IsPin     Bool
	Sort      Uint16
	CreatedAt Time
}

func New_X_Dict(tableName string) Dict_Active {
	return Dict_Active{
		xTableName: tableName,

		ALL:       NewAsterisk(tableName),
		Id:        NewInt64(tableName, "id"),
		Pid:       NewInt64(tableName, "pid"),
		Name:      NewString(tableName, "name"),
		Score:     NewFloat64(tableName, "score"),
		IsPin:     NewBool(tableName, "is_pin"),
		Sort:      NewUint16(tableName, "sort"),
		CreatedAt: NewTime(tableName, "created_at"),
	}
}

func X_Dict() Dict_Active {
	return xDict
}

func (*Dict_Active) As(alias string) Dict_Active {
	return New_X_Dict(alias)
}

func (*Dict_Active) TableName() string {
	return "dict"
}

func (x *Dict_Active) X_Alias() string {
	return x.xTableName
}
func (*Dict_Active) New_Executor(db *gorm.DB) *Executor[Dict] {
	return NewExecutor[Dict](db)
}

type DictItem struct {
	Id        int64
	DictId    int64
	DictName  string
	Name      string
	Sort      uint32
	IsEnabled bool
}

type DictItem_Active struct {
	// private fields
	xTableName string

	ALL       Asterisk
	Id        Int64
	DictId    Int64
	DictName  String
	Name      String
	Sort      Uint32
	IsEnabled Bool
}

func New_X_DictItem(tableName string) DictItem_Active {
	return DictItem_Active{
		xTableName: tableName,

		ALL:       NewAsterisk(tableName),
		Id:        NewInt64(tableName, "id"),
		DictId:    NewInt64(tableName, "dict_id"),
		DictName:  NewString(tableName, "dict_name"),
		Name:      NewString(tableName, "name"),
		Sort:      NewUint32(tableName, "sort"),
		IsEnabled: NewBool(tableName, "is_enabled"),
	}
}

func X_DictItem() DictItem_Active {
	return xDictItem
}
func (*DictItem_Active) As(alias string) DictItem_Active {
	return New_X_DictItem(alias)
}
func (*DictItem_Active) TableName() string {
	return "dict_item"
}
func (x *DictItem_Active) X_Alias() string {
	return x.xTableName
}
func (*DictItem_Active) New_Executor(db *gorm.DB) *Executor[DictItem] {
	return NewExecutor[DictItem](db)
}

func newDb() *gorm.DB {
	return db.Session(&gorm.Session{DryRun: true})
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
			CheckBuildExprSql(t, tt.db, tt.want, tt.wantVars)
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
						xDict.Id,
						xDict.CreatedAt.UnixTimestamp().As("created_at"),
						xDict.CreatedAt.UnixTimestamp().IfNull(0).As("created_at1"),
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
					SelectExpr(xDict.Id, xDict.Score),
				).
				Where(xDict.Name.Eq(""), xDict.IsPin.Is(true)).
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
					SelectExpr(xDict.Score.Avg()),
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
						xDict.Score,
						xDict.IsPin,
					),
				).
				Where(xDict.Id.Eq(100)).
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
			CheckBuildExprSql(t, tt.db, tt.want, tt.wantVars)
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
						xDict.CreatedAt,
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
						xDict.Score,
						xDict.CreatedAt,
					),
				).
				Take(&dummy),
			wantVars: []any{1},
			want:     "SELECT `dict`.`id`,`dict`.`pid`,`dict`.`name`,`dict`.`is_pin`,`dict`.`sort` FROM `dict` LIMIT ?",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			CheckBuildExprSql(t, tt.db, tt.want, tt.wantVars)
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
					SelectExpr(xDict.Id),
				).
				Take(&Dict{}),
			wantVars: []any{1},
			want:     "SELECT DISTINCT `dict`.`id` FROM `dict` LIMIT ?",
		},
		{
			name: "distinct field",
			db: newDb().Model(&Dict{}).
				Scopes(DistinctExpr(xDict.Id)).
				Take(&Dict{}),
			wantVars: []any{1},
			want:     "SELECT DISTINCT `dict`.`id` FROM `dict` LIMIT ?",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			CheckBuildExprSql(t, tt.db, tt.want, tt.wantVars)
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
			name: "",
			db: newDb().Model(&Dict{}).
				Scopes(
					OrderExpr(),
				).
				Take(&dummy),
			wantVars: []any{1},
			want:     "SELECT * FROM `dict` LIMIT ?",
		},
		{
			name: "",
			db: newDb().Model(&Dict{}).
				Scopes(
					OrderExpr(xDict.Score),
				).
				Take(&dummy),
			wantVars: []any{1},
			want:     "SELECT * FROM `dict` ORDER BY `dict`.`score` LIMIT ?",
		},
		{
			name: "",
			db: newDb().Model(&Dict{}).
				Scopes(
					OrderExpr(xDict.Score.Desc()),
				).
				Take(&dummy),
			wantVars: []any{1},
			want:     "SELECT * FROM `dict` ORDER BY `dict`.`score` DESC LIMIT ?",
		},
		{
			name: "",
			db: newDb().Model(&Dict{}).
				Scopes(
					OrderExpr(xDict.Score.Desc(), xDict.Name),
				).
				Take(&dummy),
			wantVars: []any{1},
			want:     "SELECT * FROM `dict` ORDER BY `dict`.`score` DESC,`dict`.`name` LIMIT ?",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			CheckBuildExprSql(t, tt.db, tt.want, tt.wantVars)
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
					GroupExpr(xDict.Name),
				).
				Take(&dummy),
			wantVars: []any{1},
			want:     "SELECT * FROM `dict` GROUP BY `dict`.`name` LIMIT ?",
		},
		{
			name: "",
			db: newDb().Model(&Dict{}).
				Scopes(
					SelectExpr(xDict.Score.Sum()),
					GroupExpr(xDict.Name),
				).
				Having(xDict.Score.Sum().Gt(100)).
				Take(&dummy),
			wantVars: []any{float64(100), 1},
			want:     "SELECT SUM(`dict`.`score`) FROM `dict` GROUP BY `dict`.`name` HAVING SUM(`dict`.`score`) > ? LIMIT ?",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			CheckBuildExprSql(t, tt.db, tt.want, tt.wantVars)
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
			CheckBuildExprSql(t, tt.db, tt.want, tt.wantVars)
		})
	}
}

func CheckBuildExprSql(t *testing.T, db *gorm.DB, want string, vars []any) {
	stmt := db.Statement
	if got := stmt.SQL.String(); got != want {
		t.Errorf("SQL expects %v got %v", want, got)
	}
	if !reflect.DeepEqual(stmt.Vars, vars) {
		t.Errorf("Vars expects %+v got %+v", vars, stmt.Vars)
	}
}
