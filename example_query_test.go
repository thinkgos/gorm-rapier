package rapier_test

import (
	"testing"
	"time"

	rapier "github.com/thinkgos/gorm-rapier"
	"github.com/thinkgos/gorm-rapier/testdata"
)

func Test_Example_Query_SingleObject(t *testing.T) {
	var dummy testdata.Dict

	// Get the first record ordered by primary key
	record1, err := rapier.NewExecutor[testdata.Dict](db).FirstOne()
	_ = err     // return error
	_ = record1 // return record
	// Get one record, no specified order
	record1, err = rapier.NewExecutor[testdata.Dict](db).TakeOne()
	_ = err     // return error
	_ = record1 // return record
	// Get one record, no specified order
	record1, err = rapier.NewExecutor[testdata.Dict](db).LastOne()
	_ = err     // return error
	_ = record1 // return record

	// Get the first record ordered by primary key  with original gorm api
	err = rapier.NewExecutor[testdata.Dict](db).First(&dummy)
	_ = err     // return error
	_ = record1 // return record
	// Get one record, no specified order with original gorm api
	err = rapier.NewExecutor[testdata.Dict](db).Take(&dummy)
	_ = err     // return error
	_ = record1 // return record
	// Get one record, no specified order with original gorm api
	err = rapier.NewExecutor[testdata.Dict](db).Last(&dummy)
	_ = err     // return error
	_ = record1 // return record
}

func Test_Example_Query_SingleObject_SingleFiled(t *testing.T) {
	var err error

	refDict := testdata.Ref_Dict()
	// Get the first record ordered returned single field.
	_, err = rapier.NewExecutor[testdata.Dict](db).SelectExpr(refDict.Key).FirstBool()
	_ = err // return error
	_, err = rapier.NewExecutor[testdata.Dict](db).SelectExpr(refDict.Key).FirstString()
	_ = err // return error
	_, err = rapier.NewExecutor[testdata.Dict](db).SelectExpr(refDict.Key).FirstFloat32()
	_ = err // return error
	_, err = rapier.NewExecutor[testdata.Dict](db).SelectExpr(refDict.Key).FirstFloat64()
	_ = err // return error
	_, err = rapier.NewExecutor[testdata.Dict](db).SelectExpr(refDict.Key).FirstInt()
	_ = err // return error
	_, err = rapier.NewExecutor[testdata.Dict](db).SelectExpr(refDict.Key).FirstInt8()
	_ = err // return error
	_, err = rapier.NewExecutor[testdata.Dict](db).SelectExpr(refDict.Key).FirstInt16()
	_ = err // return error
	_, err = rapier.NewExecutor[testdata.Dict](db).SelectExpr(refDict.Key).FirstInt32()
	_ = err // return error
	_, err = rapier.NewExecutor[testdata.Dict](db).SelectExpr(refDict.Key).FirstInt64()
	_ = err // return error
	_, err = rapier.NewExecutor[testdata.Dict](db).SelectExpr(refDict.Key).FirstUint()
	_ = err // return error
	_, err = rapier.NewExecutor[testdata.Dict](db).SelectExpr(refDict.Key).FirstUint8()
	_ = err // return error
	_, err = rapier.NewExecutor[testdata.Dict](db).SelectExpr(refDict.Key).FirstUint16()
	_ = err // return error
	_, err = rapier.NewExecutor[testdata.Dict](db).SelectExpr(refDict.Key).FirstUint32()
	_ = err // return error
	_, err = rapier.NewExecutor[testdata.Dict](db).SelectExpr(refDict.Key).FirstUint64()
	_ = err // return error

	// Get one record, no specified order returned single field.
	_, err = rapier.NewExecutor[testdata.Dict](db).SelectExpr(refDict.Key).TakeBool()
	_ = err // return error
	_, err = rapier.NewExecutor[testdata.Dict](db).SelectExpr(refDict.Key).TakeString()
	_ = err // return error
	_, err = rapier.NewExecutor[testdata.Dict](db).SelectExpr(refDict.Key).TakeFloat32()
	_ = err // return error
	_, err = rapier.NewExecutor[testdata.Dict](db).SelectExpr(refDict.Key).TakeFloat64()
	_ = err // return error
	_, err = rapier.NewExecutor[testdata.Dict](db).SelectExpr(refDict.Key).TakeInt()
	_ = err // return error
	_, err = rapier.NewExecutor[testdata.Dict](db).SelectExpr(refDict.Key).TakeInt8()
	_ = err // return error
	_, err = rapier.NewExecutor[testdata.Dict](db).SelectExpr(refDict.Key).TakeInt16()
	_ = err // return error
	_, err = rapier.NewExecutor[testdata.Dict](db).SelectExpr(refDict.Key).TakeInt32()
	_ = err // return error
	_, err = rapier.NewExecutor[testdata.Dict](db).SelectExpr(refDict.Key).TakeInt64()
	_ = err // return error
	_, err = rapier.NewExecutor[testdata.Dict](db).SelectExpr(refDict.Key).TakeUint()
	_ = err // return error
	_, err = rapier.NewExecutor[testdata.Dict](db).SelectExpr(refDict.Key).TakeUint8()
	_ = err // return error
	_, err = rapier.NewExecutor[testdata.Dict](db).SelectExpr(refDict.Key).TakeUint16()
	_ = err // return error
	_, err = rapier.NewExecutor[testdata.Dict](db).SelectExpr(refDict.Key).TakeUint32()
	_ = err // return error
	_, err = rapier.NewExecutor[testdata.Dict](db).SelectExpr(refDict.Key).TakeUint64()
	_ = err // return error

	_ = err // return error
	// Get one record, no specified order
	_, err = rapier.NewExecutor[testdata.Dict](db).LastOne()
	_ = err // return error
}

func Test_Example_Query_MultipleObject(t *testing.T) {
	// Get the multiple record.
	records1, err := rapier.NewExecutor[testdata.Dict](db).
		FindAll()
	_ = err      // return error
	_ = records1 // return records

	var records2 []*testdata.Dict
	// Get the multiple record.
	err = rapier.NewExecutor[testdata.Dict](db).
		SelectExpr(rapier.All).
		Find(&records2)
	_ = err      // return error
	_ = records1 // return records
}
func Test_Example_Query_Condition(t *testing.T) {
	refDict := testdata.Ref_Dict()

	// =
	_, _ = rapier.NewExecutor[testdata.Dict](db).Where(refDict.Key.Eq("key1")).TakeOne()
	// <>
	_, _ = rapier.NewExecutor[testdata.Dict](db).Where(refDict.Key.Neq("key1")).TakeOne()
	// IN
	_, _ = rapier.NewExecutor[testdata.Dict](db).Where(refDict.Key.In("key1", "key2")).TakeOne()
	// NOT IN
	_, _ = rapier.NewExecutor[testdata.Dict](db).Where(refDict.Key.NotIn("key1", "key2")).TakeOne()
	// Fuzzy LIKE
	_, _ = rapier.NewExecutor[testdata.Dict](db).Where(refDict.Key.FuzzyLike("key1")).TakeOne()
	// Left LIKE
	_, _ = rapier.NewExecutor[testdata.Dict](db).Where(refDict.Key.LeftLike("key1")).TakeOne()
	// LIKE
	_, _ = rapier.NewExecutor[testdata.Dict](db).Where(refDict.Key.Like("%key1%")).TakeOne()
	// NOT LIKE
	_, _ = rapier.NewExecutor[testdata.Dict](db).Where(refDict.Key.NotLike("%key1%")).TakeOne()
	// AND
	_, _ = rapier.NewExecutor[testdata.Dict](db).Where(refDict.Key.Eq("key1"), refDict.IsPin.Eq(true)).TakeOne()
	// >
	_, _ = rapier.NewExecutor[testdata.Dict](db).Where(refDict.CreatedAt.Gt(time.Now())).TakeOne()
	// >=
	_, _ = rapier.NewExecutor[testdata.Dict](db).Where(refDict.CreatedAt.Gte(time.Now())).TakeOne()
	// <
	_, _ = rapier.NewExecutor[testdata.Dict](db).Where(refDict.CreatedAt.Lt(time.Now())).TakeOne()
	// <=
	_, _ = rapier.NewExecutor[testdata.Dict](db).Where(refDict.CreatedAt.Lte(time.Now())).TakeOne()
	// BETWEEN
	_, _ = rapier.NewExecutor[testdata.Dict](db).Where(refDict.CreatedAt.Between(time.Now().Add(time.Hour), time.Now())).TakeOne()
	// NOT BETWEEN
	_, _ = rapier.NewExecutor[testdata.Dict](db).Where(refDict.CreatedAt.NotBetween(time.Now().Add(time.Hour), time.Now())).TakeOne()

	// not condition
	_, _ = rapier.NewExecutor[testdata.Dict](db).Not(refDict.Key.Eq("key1")).TakeOne()
	_, _ = rapier.NewExecutor[testdata.Dict](db).Where(rapier.Not(refDict.Key.Eq("key1"))).TakeOne()
	_, _ = rapier.NewExecutor[testdata.Dict](db).Not(refDict.Key.In("key1", "key2")).TakeOne()
	_, _ = rapier.NewExecutor[testdata.Dict](db).Where(rapier.Not(refDict.Key.In("key1", "key2"))).TakeOne()
	// Or condition
	_, _ = rapier.NewExecutor[testdata.Dict](db).Where(refDict.Key.Eq("key1")).Or(refDict.Key.Eq("key2")).TakeOne()
	_, _ = rapier.NewExecutor[testdata.Dict](db).Where(rapier.Or(refDict.Key.Eq("key1"), refDict.Key.Eq("key2"))).TakeOne()
	_, _ = rapier.NewExecutor[testdata.Dict](db).Where(refDict.Key.Eq("key1")).Or(refDict.Key.Eq("key2"), refDict.IsPin.Eq(true)).TakeOne()
	_, _ = rapier.NewExecutor[testdata.Dict](db).Where(rapier.Or(refDict.Key.Eq("key1"), rapier.And(refDict.Key.Eq("key2"), refDict.IsPin.Eq(true)))).TakeOne()
}

func Test_Example_Query_SelectSpecificFields(t *testing.T) {
	var records []*struct {
		Key   string
		IsPin bool
	}
	refDict := testdata.Ref_Dict()

	// with expr
	_ = rapier.NewExecutor[testdata.Dict](db).SelectExpr(refDict.Key, refDict.IsPin).Find(&records)
	_ = rapier.NewExecutor[testdata.Dict](db).SelectExpr(refDict.Key.Trim("1").As(refDict.Key.ColumnName()), refDict.IsPin).Find(&records)

	// with original gorm api
	_ = rapier.NewExecutor[testdata.Dict](db).Select("key", "is_pin").Find(&records)
}

func Test_Example_Query_Order(t *testing.T) {
	refDict := testdata.Ref_Dict()

	// with expr
	_, _ = rapier.NewExecutor[testdata.Dict](db).OrderExpr(refDict.Key.Desc(), refDict.Name).FindAll()
	_, _ = rapier.NewExecutor[testdata.Dict](db).OrderExpr(refDict.Key.Desc()).OrderExpr(refDict.Name).FindAll()

	// with original gorm api
	_, _ = rapier.NewExecutor[testdata.Dict](db).Order("`key` DESC,name").FindAll()
	_, _ = rapier.NewExecutor[testdata.Dict](db).Order("`key` DESC").Order("name").FindAll()
}

func Test_Example_Query_Scan(t *testing.T) {
	var dummy testdata.Dict

	// Get one record, no specified order
	record1, err := rapier.NewExecutor[testdata.Dict](db).ScanOne()
	_ = err     // return error
	_ = record1 // return record
	// Get one record, no specified order with original gorm api
	err = rapier.NewExecutor[testdata.Dict](db).Scan(&dummy)
	_ = err     // return error
	_ = record1 // return record
}
