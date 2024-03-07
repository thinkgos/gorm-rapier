package rapier_test

import (
	"testing"

	rapier "github.com/thinkgos/gorm-rapier"
	"github.com/thinkgos/gorm-rapier/testdata"
	"gorm.io/gorm/clause"
)

func Test_Example_AdvanceQuery_Locking(t *testing.T) {
	// Basic FOR UPDATE lock
	_, _ = rapier.NewExecutor[testdata.Dict](db).
		LockingUpdate().
		TakeOne()
	// Basic FOR UPDATE lock with Clauses api
	_, _ = rapier.NewExecutor[testdata.Dict](db).
		Clauses(clause.Locking{Strength: "UPDATE"}).
		TakeOne()

	// Basic FOR SHARE lock
	_, _ = rapier.NewExecutor[testdata.Dict](db).
		LockingShare().
		TakeOne()
	// Basic FOR SHARE lock with Clauses api
	_, _ = rapier.NewExecutor[testdata.Dict](db).
		Clauses(clause.Locking{Strength: "SHARE"}).
		TakeOne()

	// Basic FOR UPDATE NOWAIT lock with Clauses api
	_, _ = rapier.NewExecutor[testdata.Dict](db).
		Clauses(clause.Locking{Strength: "UPDATE", Options: "NOWAIT"}).
		TakeOne()
}

func Test_Example_AdvanceQuery_SubQuery(t *testing.T) {
	refDict := testdata.Ref_Dict()
	_, _ = rapier.NewExecutor[testdata.Dict](db).
		Where(refDict.Key.EqSubQuery(
			rapier.NewExecutor[testdata.Dict](db).
				SelectExpr(refDict.Key).
				Where(refDict.Id.Eq(1001)).
				IntoDB(),
		)).
		FindAll()
}

func Test_Example_AdvanceQuery_FromSubQuery(t *testing.T) {
	refDict := testdata.Ref_Dict()
	_, _ = rapier.NewExecutor[testdata.Dict](db).
		TableExpr(
			rapier.From{
				Alias: "u",
				SubQuery: rapier.NewExecutor[testdata.Dict](db).
					SelectExpr(refDict.Key).
					IntoDB(),
			},
			rapier.From{
				Alias: "p",
				SubQuery: rapier.NewExecutor[testdata.Dict](db).
					SelectExpr(refDict.Key).
					IntoDB(),
			},
		).
		FindAll()
}

func Test_Example_AdvanceQuery_Pluck(t *testing.T) {
	var ids []int64

	refDict := testdata.Ref_Dict()
	// with expr api
	_, _ = rapier.NewExecutor[testdata.Dict](db).PluckExprString(refDict.Name)
	_, _ = rapier.NewExecutor[testdata.Dict](db).PluckExprBool(refDict.IsPin)
	_ = rapier.NewExecutor[testdata.Dict](db).Pluck("id", &ids)
	_, _ = rapier.NewExecutor[testdata.Dict](db).PluckExprInt(refDict.Id)
	_, _ = rapier.NewExecutor[testdata.Dict](db).PluckExprInt8(refDict.Id)
	_, _ = rapier.NewExecutor[testdata.Dict](db).PluckExprInt16(refDict.Id)
	_, _ = rapier.NewExecutor[testdata.Dict](db).PluckExprInt32(refDict.Id)
	_, _ = rapier.NewExecutor[testdata.Dict](db).PluckExprInt64(refDict.Id)
	_, _ = rapier.NewExecutor[testdata.Dict](db).PluckExprUint(refDict.Id)
	_, _ = rapier.NewExecutor[testdata.Dict](db).PluckExprUint8(refDict.Id)
	_, _ = rapier.NewExecutor[testdata.Dict](db).PluckExprUint16(refDict.Id)
	_, _ = rapier.NewExecutor[testdata.Dict](db).PluckExprUint32(refDict.Id)
	_, _ = rapier.NewExecutor[testdata.Dict](db).PluckExprUint64(refDict.Id)

	// with original gorm api
	_, _ = rapier.NewExecutor[testdata.Dict](db).PluckString("name")
	_, _ = rapier.NewExecutor[testdata.Dict](db).PluckBool("is_pin")
	_ = rapier.NewExecutor[testdata.Dict](db).Pluck("id", &ids)
	_, _ = rapier.NewExecutor[testdata.Dict](db).PluckInt("id")
	_, _ = rapier.NewExecutor[testdata.Dict](db).PluckInt8("id")
	_, _ = rapier.NewExecutor[testdata.Dict](db).PluckInt16("id")
	_, _ = rapier.NewExecutor[testdata.Dict](db).PluckInt32("id")
	_, _ = rapier.NewExecutor[testdata.Dict](db).PluckInt64("id")
	_, _ = rapier.NewExecutor[testdata.Dict](db).PluckUint("id")
	_, _ = rapier.NewExecutor[testdata.Dict](db).PluckUint8("id")
	_, _ = rapier.NewExecutor[testdata.Dict](db).PluckUint16("id")
	_, _ = rapier.NewExecutor[testdata.Dict](db).PluckUint32("id")
	_, _ = rapier.NewExecutor[testdata.Dict](db).PluckUint64("id")
}
func Test_Example_AdvanceQuery_Count(t *testing.T) {
	total, err := rapier.NewExecutor[testdata.Dict](db).Count()
	_ = err
	_ = total
}

func Test_Example_AdvanceQuery_Exist(t *testing.T) {
	refDict := testdata.Ref_Dict()
	b, err := rapier.NewExecutor[testdata.Dict](db).Where(refDict.Id.Eq(100)).Exist()
	_ = err
	_ = b
}
