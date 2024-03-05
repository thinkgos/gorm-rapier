package test

import (
	"testing"

	rapier "github.com/thinkgos/gorm-rapier"
	"github.com/thinkgos/gorm-rapier/examples/model"
)

func Test_Update(t *testing.T) {
	refDict := model.Ref_Dict()
	rowsAffected, err := rapier.NewExecutor[model.Dict](db).
		Model().
		Where(refDict.Id.Eq(100)).
		UpdatesExpr(
			refDict.Key.Value("k1"),
		)
	_ = err          // return error
	_ = rowsAffected // return row affected
}
