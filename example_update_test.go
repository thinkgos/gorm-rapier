package rapier_test

import (
	"testing"
	"time"

	rapier "github.com/thinkgos/gorm-rapier"
	"github.com/thinkgos/gorm-rapier/testdata"
)

func Test_Example_Save(t *testing.T) {
	rowsAffected, err := rapier.NewExecutor[testdata.Dict](db).
		Save(&testdata.Dict{
			Id:     100,
			Key:    "k1",
			Remark: "remark1",
		})
	_ = err          // return error
	_ = rowsAffected // return row affected
}

func Test_Example_UpdateColumn(t *testing.T) {
	refDict := testdata.Ref_Dict()
	// update with expr
	rowsAffected, err := rapier.NewExecutor[testdata.Dict](db).
		Model().
		Where(refDict.Id.Eq(100)).
		UpdateExpr(refDict.Key, "k1")
	_ = err          // return error
	_ = rowsAffected // return row affected

	// update SetExpr with expr
	rowsAffected, err = rapier.NewExecutor[testdata.Dict](db).
		Model().
		Where(refDict.Id.Eq(100)).
		UpdateExpr(refDict.UpdatedAt, refDict.CreatedAt.Add(time.Second))
	_ = err          // return error
	_ = rowsAffected // return row affected

	// update with original gorm api
	rowsAffected, err = rapier.NewExecutor[testdata.Dict](db).
		Model().
		Where(refDict.Id.Eq(100)).
		Update("key", "k1")
	_ = err          // return error
	_ = rowsAffected // return row affected
}

func Test_Example_UpdateMultipleColumns(t *testing.T) {
	refDict := testdata.Ref_Dict()
	// update with expr
	rowsAffected, err := rapier.NewExecutor[testdata.Dict](db).
		Model().
		Where(refDict.Id.Eq(100)).
		UpdatesExpr(
			refDict.Key.Value("k1"),
			refDict.Remark.Value(""),
		)
	_ = err          // return error
	_ = rowsAffected // return row affected

	// update use `struct` with original gorm api
	rowsAffected, err = rapier.NewExecutor[testdata.Dict](db).
		Model().
		Where(refDict.Id.Eq(100)).
		Updates(&testdata.Dict{
			Key:    "k1",
			Remark: "remark1",
		})
	_ = err          // return error
	_ = rowsAffected // return row affected

	// update use map with original gorm api
	rowsAffected, err = rapier.NewExecutor[testdata.Dict](db).
		Model().
		Where(refDict.Id.Eq(100)).
		UpdatesMap(map[string]any{
			"key":    "k1",
			"remark": "remark1",
		})
	_ = err          // return error
	_ = rowsAffected // return row affected
}

func Test_Example_UpdateFromSubQuery(t *testing.T) {
	refDict := testdata.Ref_Dict()
	// update with expr
	rowsAffected, err := rapier.NewExecutor[testdata.Dict](db).
		Model().
		Where(refDict.Id.Eq(100)).
		UpdateExpr(
			refDict.Key,
			rapier.NewExecutor[testdata.Dict](db).Model().
				SelectExpr(refDict.Key).
				Where(refDict.Id.Eq(101)).
				IntoDB(),
		)
	_ = err          // return error
	_ = rowsAffected // return row affected

	// update with exprs
	rowsAffected, err = rapier.NewExecutor[testdata.Dict](db).
		Model().
		Where(refDict.Id.Eq(100)).
		UpdatesExpr(
			refDict.Key.ValueSubQuery(
				rapier.NewExecutor[testdata.Dict](db).Model().
					SelectExpr(refDict.Key).
					Where(refDict.Id.Eq(101)).
					IntoDB(),
			),
		)
	_ = err          // return error
	_ = rowsAffected // return row affected

	// update use map with original gorm api
	rowsAffected, err = rapier.NewExecutor[testdata.Dict](db).
		Model().
		Where(refDict.Id.Eq(100)).
		UpdatesMap(map[string]any{
			"key": rapier.NewExecutor[testdata.Dict](db).Model().
				SelectExpr(refDict.Key).
				Where(refDict.Id.Eq(101)).
				IntoDB(),
		})
	_ = err          // return error
	_ = rowsAffected // return row affected
}

func Test_Example_UpdateWithoutTimeTracking(t *testing.T) {
	refDict := testdata.Ref_Dict()
	// update with expr
	rowsAffected, err := rapier.NewExecutor[testdata.Dict](db).
		Model().
		Where(refDict.Id.Eq(100)).
		UpdateColumnsExpr(
			refDict.Key.Value("k1"),
			refDict.Remark.Value(""),
		)
	_ = err          // return error
	_ = rowsAffected // return row affected

	// update with expr
	rowsAffected, err = rapier.NewExecutor[testdata.Dict](db).
		Model().
		Where(refDict.Id.Eq(100)).
		UpdateColumnExpr(refDict.Key, "k1")
	_ = err          // return error
	_ = rowsAffected // return row affected

	// update with original gorm api
	rowsAffected, err = rapier.NewExecutor[testdata.Dict](db).
		Model().
		Where(refDict.Id.Eq(100)).
		UpdateColumn("key", "k1")
	_ = err          // return error
	_ = rowsAffected // return row affected

	// update with original gorm api
	rowsAffected, err = rapier.NewExecutor[testdata.Dict](db).
		Model().
		Where(refDict.Id.Eq(100)).
		UpdateColumns(&testdata.Dict{
			Key: "k1",
		})
	_ = err          // return error
	_ = rowsAffected // return row affected

	// update with original gorm api
	rowsAffected, err = rapier.NewExecutor[testdata.Dict](db).
		Model().
		Where(refDict.Id.Eq(100)).
		UpdateColumnsMap(map[string]any{
			"key": "k1",
		})
	_ = err          // return error
	_ = rowsAffected // return row affected
}
