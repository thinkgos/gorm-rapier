package assist

import (
	"reflect"
	"testing"
	"time"

	"gorm.io/gorm"
)

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

var x_Dict_Model_Type = Indirect(&Dict{})

type DictImpl struct {
	// private fields
	xTableName string
	xModelType reflect.Type

	ALL Asterisk

	Id        Int64
	Pid       Int64
	Score     Float64
	IsPin     Bool
	Sort      Uint16
	Name      String
	CreatedAt Time
}

func new_X_Dict(tableName string) DictImpl {
	return DictImpl{
		xTableName: tableName,
		xModelType: x_Dict_Model_Type,

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

func New_X_Dict() DictImpl {
	return new_X_Dict("dict")
}

func (x *DictImpl) As(alias string) DictImpl {
	return new_X_Dict(alias)
}

func (x *DictImpl) X_Model() any {
	return reflect.New(x.xModelType).Interface()
}

func (x *DictImpl) Xc_Model() Condition {
	return func(db *gorm.DB) *gorm.DB {
		return db.Model(reflect.New(x.xModelType).Interface())
	}
}

func (x *DictImpl) X_TableName() string {
	return x.xTableName
}

var xDict = New_X_Dict()

func newDb() *gorm.DB {
	return db.Session(&gorm.Session{DryRun: true})
}

func Test_Table(t *testing.T) {
	var dummy Dict

	tests := []struct {
		name     string
		db       *gorm.DB
		wantVars []interface{}
		want     string
	}{
		{
			name: "empty table",
			db: newDb().Model(xDict.X_Model()).
				Scopes(
					Table(),
				).
				Take(&dummy),
			wantVars: nil,
			want:     "SELECT * FROM `dict` LIMIT 1",
		},
		{
			name: "single table",
			db: newDb().Model(xDict.X_Model()).
				Scopes(
					Table(
						From{
							"a",
							newDb().
								Model(xDict.X_Model()),
						},
					),
				).
				Take(&dummy),
			wantVars: nil,
			want:     "SELECT * FROM (SELECT * FROM `dict`) AS `a` LIMIT 1",
		},
		{
			name: "multi table",
			db: newDb().Model(xDict.X_Model()).
				Scopes(
					Table(
						From{
							"a",
							newDb().
								Model(xDict.X_Model()),
						},
						From{
							"b",
							newDb().
								Model(xDict.X_Model()),
						},
					),
				).
				Take(&dummy),
			wantVars: nil,
			want:     "SELECT * FROM (SELECT * FROM `dict`) AS `a`, (SELECT * FROM `dict`) AS `b` LIMIT 1",
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
		wantVars []interface{}
		want     string
	}{
		{
			name: "select *",
			db: newDb().
				Scopes(
					xDict.Xc_Model(),
					Select(),
				).
				Take(&Dict{}),
			wantVars: nil,
			want:     "SELECT * FROM `dict` LIMIT 1",
		},
		{
			name: "select field",
			db: newDb().Model(xDict.X_Model()).
				Scopes(
					Select(
						xDict.Id,
						xDict.CreatedAt.UnixTimestamp().As("created_at"),
						xDict.CreatedAt.UnixTimestamp().IfNull(0).As("created_at1"),
					),
				).
				Take(&dummy),
			wantVars: []any{int64(0)},
			want:     "SELECT `dict`.`id`,UNIX_TIMESTAMP(`dict`.`created_at`) AS `created_at`,IFNULL(UNIX_TIMESTAMP(`dict`.`created_at`),?) AS `created_at1` FROM `dict` LIMIT 1",
		},
		{
			name: "select field where",
			db: newDb().Model(xDict.X_Model()).
				Scopes(
					Select(xDict.Id, xDict.Score),
				).
				Where(xDict.Name.Eq(""), xDict.IsPin.Is(true)).
				Take(&dummy),
			wantVars: []any{"", true},
			want:     "SELECT `dict`.`id`,`dict`.`score` FROM `dict` WHERE `dict`.`name` = ? AND `dict`.`is_pin` = ? LIMIT 1",
		},
		{
			name: "select 1",
			db: newDb().Model(xDict.X_Model()).
				Scopes(
					Select(One),
				).
				Take(&dummy),
			wantVars: nil,
			want:     "SELECT 1 FROM `dict` LIMIT 1",
		},
		{
			name: "select COUNT(1)",
			db: newDb().Model(xDict.X_Model()).
				Scopes(
					Select(One.Count()),
				).
				Take(&dummy),
			wantVars: nil,
			want:     "SELECT COUNT(1) FROM `dict` LIMIT 1",
		},
		{
			name: "select COUNT(*)",
			db: newDb().Model(xDict.X_Model()).
				Scopes(
					Select(Star.Count()),
				).
				Take(&dummy),
			wantVars: nil,
			want:     "SELECT COUNT(*) FROM `dict` LIMIT 1",
		},
		{
			name: "select AVG(field)",
			db: newDb().Model(xDict.X_Model()).
				Scopes(
					Select(xDict.Score.Avg()),
				).
				Take(&dummy),
			wantVars: nil,
			want:     "SELECT AVG(`dict`.`score`) FROM `dict` LIMIT 1",
		},
		{
			name: "update with select field",
			db: newDb().Model(xDict.X_Model()).
				Scopes(
					Select(
						xDict.Score,
						xDict.IsPin,
					),
				).
				Where(xDict.Id.Eq(100)).
				Updates(&Dict{
					Score: 100,
					IsPin: true,
				}),
			wantVars: []interface{}{float64(100), true, int64(100)},
			want:     "UPDATE `dict` SET `score`=?,`is_pin`=? WHERE `dict`.`id` = ?",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			CheckBuildExprSql(t, tt.db, tt.want, tt.wantVars)
		})
	}
}

func Test_Where(t *testing.T) {
	var dummy Dict

	tests := []struct {
		name     string
		db       *gorm.DB
		wantVars []interface{}
		want     string
	}{
		{
			name: "",
			db: newDb().Model(xDict.X_Model()).
				Scopes(
					Where(xDict.Id.Eq(100)),
				).
				Take(&dummy),
			wantVars: []any{int64(100)},
			want:     "SELECT * FROM `dict` WHERE `dict`.`id` = ? LIMIT 1",
		},
		{
			name: "",
			db: newDb().Model(xDict.X_Model()).
				Scopes(
					Where(),
				).
				Take(&dummy),
			wantVars: nil,
			want:     "SELECT * FROM `dict` LIMIT 1",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			CheckBuildExprSql(t, tt.db, tt.want, tt.wantVars)
		})
	}
}

func Test_Having(t *testing.T) {
	var dummy Dict

	tests := []struct {
		name     string
		db       *gorm.DB
		wantVars []interface{}
		want     string
	}{
		{
			name: "",
			db: newDb().Model(xDict.X_Model()).
				Scopes(
					Having(xDict.Id.Eq(100)),
					Group(xDict.Id),
				).
				Take(&dummy),
			wantVars: []any{int64(100)},
			want:     "SELECT * FROM `dict` GROUP BY `dict`.`id` HAVING `dict`.`id` = ? LIMIT 1",
		},
		{
			name: "",
			db: newDb().Model(xDict.X_Model()).
				Scopes(
					Having(),
				).
				Take(&dummy),
			wantVars: nil,
			want:     "SELECT * FROM `dict` LIMIT 1",
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
		wantVars []interface{}
		want     string
	}{
		{
			name: "",
			db: newDb().Model(xDict.X_Model()).
				Scopes(
					Order(),
				).
				Take(&dummy),
			wantVars: nil,
			want:     "SELECT * FROM `dict` LIMIT 1",
		},
		{
			name: "",
			db: newDb().Model(xDict.X_Model()).
				Scopes(
					Order(xDict.Score),
				).
				Take(&dummy),
			wantVars: nil,
			want:     "SELECT * FROM `dict` ORDER BY `dict`.`score` LIMIT 1",
		},
		{
			name: "",
			db: newDb().Model(xDict.X_Model()).
				Scopes(
					Order(xDict.Score.Desc()),
				).
				Take(&dummy),
			wantVars: nil,
			want:     "SELECT * FROM `dict` ORDER BY `dict`.`score` DESC LIMIT 1",
		},
		{
			name: "",
			db: newDb().Model(xDict.X_Model()).
				Scopes(
					Order(xDict.Score.Desc(), xDict.Name),
				).
				Take(&dummy),
			wantVars: nil,
			want:     "SELECT * FROM `dict` ORDER BY `dict`.`score` DESC,`dict`.`name` LIMIT 1",
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
		wantVars []interface{}
		want     string
	}{
		{
			name: "",
			db: newDb().Model(xDict.X_Model()).
				Scopes(
					Group(),
				).
				Take(&dummy),
			wantVars: nil,
			want:     "SELECT * FROM `dict` LIMIT 1",
		},
		{
			name: "",
			db: newDb().Model(xDict.X_Model()).
				Scopes(
					Group(xDict.Name),
				).
				Take(&dummy),
			wantVars: nil,
			want:     "SELECT * FROM `dict` GROUP BY `dict`.`name` LIMIT 1",
		},
		{
			name: "",
			db: newDb().Model(xDict.X_Model()).
				Scopes(
					Select(xDict.Score.Sum()),
					Group(xDict.Name),
				).
				Having(xDict.Score.Sum().Gt(100)).
				Take(&dummy),
			wantVars: []any{float64(100)},
			want:     "SELECT SUM(`dict`.`score`) FROM `dict` GROUP BY `dict`.`name` HAVING SUM(`dict`.`score`) > ? LIMIT 1",
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
		wantVars []interface{}
		want     string
	}{
		{
			name: "",
			db: newDb().Model(xDict.X_Model()).
				Scopes(
					Group(),
					LockingUpdate(),
				).
				Take(&dummy),
			wantVars: nil,
			want:     "SELECT * FROM `dict` LIMIT 1 FOR UPDATE",
		},
		{
			name: "",
			db: newDb().Model(xDict.X_Model()).
				Scopes(
					Group(),
					LockingShare(),
				).
				Take(&dummy),
			wantVars: nil,
			want:     "SELECT * FROM `dict` LIMIT 1 FOR SHARE",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			CheckBuildExprSql(t, tt.db, tt.want, tt.wantVars)
		})
	}
}

func CheckBuildExprSql(t *testing.T, db *gorm.DB, want string, vars []interface{}) {
	stmt := db.Statement
	if got := stmt.SQL.String(); got != want {
		t.Errorf("SQL expects %v got %v", want, got)
	}
	if !reflect.DeepEqual(stmt.Vars, vars) {
		t.Errorf("Vars expects %+v got %v", vars, stmt.Vars)
	}
}

func Test_Conditions(t *testing.T) {
	t.Run("empty", func(t *testing.T) {
		NewConditions().
			Table().
			Select().
			Order().
			Group().
			Where().
			Having().
			LockingUpdate().
			LockingShare().
			Pagination(1, 20).
			CrossJoin(xDict.X_TableName()).
			InnerJoin(xDict.X_TableName()).
			Join(xDict.X_TableName()).
			LeftJoin(xDict.X_TableName()).
			RightJoin(xDict.X_TableName()).
			CrossJoinX(xDict.X_TableName(), "x").
			InnerJoinX(xDict.X_TableName(), "x").
			JoinX(xDict.X_TableName(), "x").
			LeftJoinX(xDict.X_TableName(), "x").
			RightJoinX(xDict.X_TableName(), "x").
			Build()
	})
}
