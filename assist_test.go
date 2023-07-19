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

var xDict = New_X_Dict("dict")

type X_DictImpl struct {
	// private fields
	xTableName string

	ALL       Asterisk
	Id        Int64
	Pid       Int64
	Score     Float64
	IsPin     Bool
	Sort      Uint16
	Name      String
	CreatedAt Time
}

func New_X_Dict(tableName string) X_DictImpl {
	return X_DictImpl{
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

func X_Dict() X_DictImpl {
	return xDict
}

func (*X_DictImpl) As(alias string) X_DictImpl {
	return New_X_Dict(alias)
}

func (*X_DictImpl) X_Model() *Dict {
	return &Dict{}
}

func (x *X_DictImpl) X_TableName() string {
	return x.xTableName
}
func (*X_DictImpl) X_Executor(db *gorm.DB) *Executor[Dict] {
	return NewExecutor[Dict](db)
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
			db: newDb().Model(xDict.X_Model()).
				Scopes(
					TableExpr(),
				).
				Take(&dummy),
			wantVars: nil,
			want:     "SELECT * FROM `dict` LIMIT 1",
		},
		{
			name: "single table",
			db: newDb().Model(xDict.X_Model()).
				Scopes(
					TableExpr(
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
					TableExpr(
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
		wantVars []any
		want     string
	}{
		{
			name: "select *",
			db: newDb().Model(xDict.X_Model()).
				Scopes(
					SelectExpr(),
				).
				Take(&dummy),
			wantVars: nil,
			want:     "SELECT * FROM `dict` LIMIT 1",
		},
		{
			name: "select field",
			db: newDb().Model(xDict.X_Model()).
				Scopes(
					SelectExpr(
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
					SelectExpr(xDict.Id, xDict.Score),
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
					SelectExpr(One),
				).
				Take(&dummy),
			wantVars: nil,
			want:     "SELECT 1 FROM `dict` LIMIT 1",
		},
		{
			name: "select COUNT(1)",
			db: newDb().Model(xDict.X_Model()).
				Scopes(
					SelectExpr(One.Count()),
				).
				Take(&dummy),
			wantVars: nil,
			want:     "SELECT COUNT(1) FROM `dict` LIMIT 1",
		},
		{
			name: "select COUNT(*)",
			db: newDb().Model(xDict.X_Model()).
				Scopes(
					SelectExpr(Star.Count()),
				).
				Take(&dummy),
			wantVars: nil,
			want:     "SELECT COUNT(*) FROM `dict` LIMIT 1",
		},
		{
			name: "select AVG(field)",
			db: newDb().Model(xDict.X_Model()).
				Scopes(
					SelectExpr(xDict.Score.Avg()),
				).
				Take(&dummy),
			wantVars: nil,
			want:     "SELECT AVG(`dict`.`score`) FROM `dict` LIMIT 1",
		},
		{
			name: "update with select field",
			db: newDb().Model(xDict.X_Model()).
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

func Test_Distinct(t *testing.T) {
	tests := []struct {
		name     string
		db       *gorm.DB
		wantVars []any
		want     string
	}{
		{
			name: "select * using distinct",
			db: newDb().Model(xDict.X_Model()).
				Scopes(
					DistinctExpr(),
					SelectExpr(xDict.Id),
				).
				Take(&Dict{}),
			wantVars: nil,
			want:     "SELECT DISTINCT `dict`.`id` FROM `dict` LIMIT 1",
		},
		{
			name: "distinct field",
			db: newDb().Model(xDict.X_Model()).
				Scopes(DistinctExpr(xDict.Id)).
				Take(&Dict{}),
			wantVars: nil,
			want:     "SELECT DISTINCT `dict`.`id` FROM `dict` LIMIT 1",
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
			db: newDb().Model(xDict.X_Model()).
				Scopes(
					OrderExpr(),
				).
				Take(&dummy),
			wantVars: nil,
			want:     "SELECT * FROM `dict` LIMIT 1",
		},
		{
			name: "",
			db: newDb().Model(xDict.X_Model()).
				Scopes(
					OrderExpr(xDict.Score),
				).
				Take(&dummy),
			wantVars: nil,
			want:     "SELECT * FROM `dict` ORDER BY `dict`.`score` LIMIT 1",
		},
		{
			name: "",
			db: newDb().Model(xDict.X_Model()).
				Scopes(
					OrderExpr(xDict.Score.Desc()),
				).
				Take(&dummy),
			wantVars: nil,
			want:     "SELECT * FROM `dict` ORDER BY `dict`.`score` DESC LIMIT 1",
		},
		{
			name: "",
			db: newDb().Model(xDict.X_Model()).
				Scopes(
					OrderExpr(xDict.Score.Desc(), xDict.Name),
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
		wantVars []any
		want     string
	}{
		{
			name: "",
			db: newDb().Model(xDict.X_Model()).
				Scopes(
					GroupExpr(),
				).
				Take(&dummy),
			wantVars: nil,
			want:     "SELECT * FROM `dict` LIMIT 1",
		},
		{
			name: "",
			db: newDb().Model(xDict.X_Model()).
				Scopes(
					GroupExpr(xDict.Name),
				).
				Take(&dummy),
			wantVars: nil,
			want:     "SELECT * FROM `dict` GROUP BY `dict`.`name` LIMIT 1",
		},
		{
			name: "",
			db: newDb().Model(xDict.X_Model()).
				Scopes(
					SelectExpr(xDict.Score.Sum()),
					GroupExpr(xDict.Name),
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
		wantVars []any
		want     string
	}{
		{
			name: "",
			db: newDb().Model(xDict.X_Model()).
				Scopes(
					GroupExpr(),
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
					GroupExpr(),
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

func CheckBuildExprSql(t *testing.T, db *gorm.DB, want string, vars []any) {
	stmt := db.Statement
	if got := stmt.SQL.String(); got != want {
		t.Errorf("SQL expects %v got %v", want, got)
	}
	if !reflect.DeepEqual(stmt.Vars, vars) {
		t.Errorf("Vars expects %+v got %+v", vars, stmt.Vars)
	}
}
