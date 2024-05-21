package rapier

import (
	"log"
	"os"
	"reflect"
	"strings"
	"sync"
	"testing"
	"time"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
	"gorm.io/gorm/utils/tests"
)

/******************* test type ***********************************************/

type TestInteger int32
type TestFloat float64
type TestString string
type TestBytes []byte
type TestTime time.Time
type TestDict struct {
	Id       int64
	Pid      int64
	Name     string
	DictItem []*TestDictItem `gorm:"foreignKey:DictId"`
}
type TestDictItem struct {
	Id     int64
	DictId int64
	Name   string
}

/******************* test function *******************************************/

var db, _ = gorm.Open(tests.DummyDialector{}, nil)

func newDb() *gorm.DB {
	return db.Session(&gorm.Session{DryRun: true})
}

func newDbWithLog() *gorm.DB {
	newDB := db.Session(&gorm.Session{DryRun: true})
	newDB.Logger = logger.New(log.New(os.Stdout, "\r\n", log.LstdFlags), logger.Config{
		SlowThreshold:             200 * time.Millisecond,
		Colorful:                  true,
		IgnoreRecordNotFoundError: false,
		LogLevel:                  logger.Info,
	})
	return newDB
}

func NewStatement() *gorm.Statement {
	dd, _ := schema.Parse(&Dict{}, &sync.Map{}, db.NamingStrategy)
	return &gorm.Statement{
		DB:      db,
		Table:   dd.Table,
		Schema:  dd,
		Clauses: map[string]clause.Clause{},
	}
}

// BuildExpr return sql and vars
func BuildExpr(e Expr) (string, []any) {
	stmt := NewStatement()
	sql, vars := e.BuildWithArgs(stmt)
	return sql, vars
}

func ReviewBuildExpr(t *testing.T, e Expr, wantSQL string, wantVars []any) {
	gotSQL, gotVars := BuildExpr(e)
	if got := strings.TrimSpace(gotSQL); got != wantSQL {
		t.Errorf("\nSQL:\n\twant: %v\n\tgot %v", wantSQL, gotSQL)
	}
	if !reflect.DeepEqual(gotVars, wantVars) {
		t.Errorf("\nVars:\n\twant %+v\n\tgot %v", wantVars, gotVars)
	}
}

func ReviewBuildDb(t *testing.T, db *gorm.DB, wantSQL string, wantVars []any) {
	stmt := db.Statement
	if gotSQL := stmt.SQL.String(); gotSQL != wantSQL {
		t.Errorf("\nSQL:\n\twant: %v\n\tgot: %v", wantSQL, gotSQL)
	}
	if !reflect.DeepEqual(stmt.Vars, wantVars) {
		t.Errorf("\nVars:\n\twant: %+v\n\tgot: %+v", wantVars, stmt.Vars)
	}
}

/******************* test model **********************************************/

type Dict struct {
	Id        int64 `gorm:"autoIncrement:true;not null;primaryKey"`
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

var refDict = New_Dict("dict")

type Dict_Active struct {
	refAlias     string
	refTableName string
	ALL          Asterisk
	Id           Int64
	Pid          Int64
	Name         String
	Score        Float64
	Sort         Uint16
	IsPin        Bool
	CreatedAt    Time
}

func new_Dict(tableName, alias string) Dict_Active {
	return Dict_Active{
		refAlias:     alias,
		refTableName: tableName,
		ALL:          NewAsterisk(alias),
		Id:           NewInt64(alias, "id"),
		Pid:          NewInt64(alias, "pid"),
		Name:         NewString(alias, "name"),
		Score:        NewFloat64(alias, "score"),
		Sort:         NewUint16(alias, "sort"),
		IsPin:        NewBool(alias, "is_pin"),
		CreatedAt:    NewTime(alias, "created_at"),
	}
}
func Ref_Dict() Dict_Active { return refDict }

func New_Dict(tableName string) Dict_Active {
	return new_Dict(tableName, tableName)
}

func (x *Dict_Active) As(alias string) Dict_Active { return new_Dict(x.refTableName, alias) }
func (x *Dict_Active) TableName() string           { return x.refTableName }
func (x *Dict_Active) Alias() string               { return x.refAlias }
func (*Dict_Active) New_Executor(db *gorm.DB) *Executor[Dict] {
	return NewExecutor[Dict](db)
}

var refDictItem = New_DictItem("dict_item")

type DictItem struct {
	Id        int64 `gorm:"autoIncrement:true;not null;primaryKey"`
	DictId    int64
	Name      string
	Sort      uint32
	IsEnabled bool
}

type DictItem_Active struct {
	refAlias     string
	refTableName string
	ALL          Asterisk
	Id           Int64
	DictId       Int64
	Name         String
	Sort         Uint32
	IsEnabled    Bool
}

func new_DictItem(tableName, alias string) DictItem_Active {
	return DictItem_Active{
		refAlias:     alias,
		refTableName: tableName,
		ALL:          NewAsterisk(alias),
		Id:           NewInt64(alias, "id"),
		DictId:       NewInt64(alias, "dict_id"),
		Name:         NewString(alias, "name"),
		Sort:         NewUint32(alias, "sort"),
		IsEnabled:    NewBool(alias, "is_enabled"),
	}
}

func Ref_DictItem() DictItem_Active { return refDictItem }

func New_DictItem(tableName string) DictItem_Active {
	return new_DictItem(tableName, tableName)
}

func (x *DictItem_Active) As(alias string) DictItem_Active {
	return new_DictItem(x.refTableName, alias)
}
func (x *DictItem_Active) TableName() string { return x.refTableName }
func (x *DictItem_Active) Alias() string     { return x.refAlias }
func (*DictItem_Active) New_Executor(db *gorm.DB) *Executor[DictItem] {
	return NewExecutor[DictItem](db)
}
