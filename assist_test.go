package assist

import (
	"reflect"
	"testing"
	"time"

	"gorm.io/gorm"
)

func newDb() *gorm.DB {
	return db.Session(&gorm.Session{DryRun: true})
}

type Dict struct {
	Id        int64
	Name      string
	Score     float64
	IsPin     bool
	Sort      uint16
	CreatedAt time.Time
}

func (*Dict) TableName() string {
	return "dict"
}

type DictImpl struct {
	// private fields
	tableName string
	modelType reflect.Type

	ALL Asterisk

	Id        Int64
	Score     Float64
	IsPin     Bool
	Sort      Uint16
	Name      String
	CreatedAt Time
}

func NewDict() DictImpl {
	d := &Dict{}
	tableName := d.TableName()
	return DictImpl{
		tableName: tableName,
		modelType: Indirect(d),

		ALL:       NewAsterisk(tableName),
		Id:        NewInt64(tableName, "id"),
		Name:      NewString(tableName, "name"),
		Score:     NewFloat64(tableName, "score"),
		IsPin:     NewBool(tableName, "is_pin"),
		Sort:      NewUint16(tableName, "sort"),
		CreatedAt: NewTime(tableName, "created_at"),
	}
}

func (d DictImpl) As(alias string) *DictImpl {
	return &DictImpl{
		tableName: alias,
		modelType: d.modelType,
		Id:        NewInt64(alias, "id"),
		Score:     NewFloat64(alias, "score"),
		Name:      NewString(alias, "name"),
		IsPin:     NewBool(alias, "is_pin"),
		Sort:      NewUint16(alias, "sort"),
		CreatedAt: NewTime(alias, "created_at"),
	}
}

func (d DictImpl) Assist_Model() any {
	return reflect.New(d.modelType).Interface()
}

var dictpm = NewDict()

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
			db: newDb().Model(dictpm.Assist_Model()).
				Scopes(
					Select(),
				).
				Take(&Dict{}),
			wantVars: nil,
			want:     "SELECT * FROM `dict` LIMIT 1",
		},
		{
			name: "select field",
			db: newDb().Model(dictpm.Assist_Model()).
				Scopes(
					Select(
						dictpm.Id,
						dictpm.CreatedAt.UnixTimestamp().As("created_at"),
						dictpm.CreatedAt.UnixTimestamp().IfNull(0).As("created_at1"),
					),
				).
				Take(&dummy),
			wantVars: []any{int64(0)},
			want:     "SELECT `dict`.`id`,UNIX_TIMESTAMP(`dict`.`created_at`) AS `created_at`,IFNULL(UNIX_TIMESTAMP(`dict`.`created_at`),?) AS `created_at1` FROM `dict` LIMIT 1",
		},
		{
			name: "select field where",
			db: newDb().Model(dictpm.Assist_Model()).
				Scopes(
					Select(dictpm.Id, dictpm.Score),
				).
				Where(dictpm.Name.Eq("")).
				Where(dictpm.IsPin.Is(true)).
				Take(&dummy),
			wantVars: []any{"", true},
			want:     "SELECT `dict`.`id`,`dict`.`score` FROM `dict` WHERE `dict`.`name` = ? AND `dict`.`is_pin` = ? LIMIT 1",
		},
		{
			name: "select 1",
			db: newDb().Model(dictpm.Assist_Model()).
				Scopes(
					Select(One),
				).
				Take(&dummy),
			wantVars: nil,
			want:     "SELECT 1 FROM `dict` LIMIT 1",
		},
		{
			name: "select COUNT(1)",
			db: newDb().Model(dictpm.Assist_Model()).
				Scopes(
					Select(One.Count()),
				).
				Take(&dummy),
			wantVars: nil,
			want:     "SELECT COUNT(1) FROM `dict` LIMIT 1",
		},
		{
			name: "select COUNT(*)",
			db: newDb().Model(dictpm.Assist_Model()).
				Scopes(
					Select(Star.Count()),
				).
				Take(&dummy),
			wantVars: nil,
			want:     "SELECT COUNT(*) FROM `dict` LIMIT 1",
		},
		{
			name: "select AVG(field)",
			db: newDb().Model(dictpm.Assist_Model()).
				Scopes(
					Select(dictpm.Score.Avg()),
				).
				Take(&dummy),
			wantVars: nil,
			want:     "SELECT AVG(`dict`.`score`) FROM `dict` LIMIT 1",
		},
		{
			name: "update with select field",
			db: newDb().Model(dictpm.Assist_Model()).
				Scopes(
					Select(
						dictpm.Score,
						dictpm.IsPin,
					),
				).
				Where(dictpm.Id.Eq(100)).
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
			db: newDb().Model(dictpm.Assist_Model()).
				Scopes(
					Order(),
				).
				Take(&dummy),
			wantVars: nil,
			want:     "SELECT * FROM `dict` LIMIT 1",
		},
		{
			name: "",
			db: newDb().Model(dictpm.Assist_Model()).
				Scopes(
					Order(dictpm.Score),
				).
				Take(&dummy),
			wantVars: nil,
			want:     "SELECT * FROM `dict` ORDER BY `dict`.`score` LIMIT 1",
		},
		{
			name: "",
			db: newDb().Model(dictpm.Assist_Model()).
				Scopes(
					Order(dictpm.Score.Desc()),
				).
				Take(&dummy),
			wantVars: nil,
			want:     "SELECT * FROM `dict` ORDER BY `dict`.`score` DESC LIMIT 1",
		},
		{
			name: "",
			db: newDb().Model(dictpm.Assist_Model()).
				Scopes(
					Order(dictpm.Score.Desc(), dictpm.Name),
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
			db: newDb().Model(dictpm.Assist_Model()).
				Scopes(
					Group(),
				).
				Take(&dummy),
			wantVars: nil,
			want:     "SELECT * FROM `dict` LIMIT 1",
		},
		{
			name: "",
			db: newDb().Model(dictpm.Assist_Model()).
				Scopes(
					Group(dictpm.Name),
				).
				Take(&dummy),
			wantVars: nil,
			want:     "SELECT * FROM `dict` GROUP BY `dict`.`name` LIMIT 1",
		},
		{
			name: "",
			db: newDb().Model(dictpm.Assist_Model()).
				Scopes(
					Select(dictpm.Score.Sum()),
					Group(dictpm.Name),
				).
				Having(dictpm.Score.Sum().Gt(100)).
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

func CheckBuildExprSql(t *testing.T, db *gorm.DB, want string, vars []interface{}) {
	stmt := db.Statement
	if got := stmt.SQL.String(); got != want {
		t.Errorf("SQL expects %v got %v", want, got)
	}
	if !reflect.DeepEqual(stmt.Vars, vars) {
		t.Errorf("Vars expects %+v got %v", vars, stmt.Vars)
	}
}
