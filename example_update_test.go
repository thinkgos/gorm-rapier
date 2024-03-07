package rapier_test

import (
	"testing"

	rapier "github.com/thinkgos/gorm-rapier"
	testdata "github.com/thinkgos/gorm-rapier/testdata"
)

func Test_Update(t *testing.T) {
	refDict := testdata.Ref_Dict()
	rowsAffected, err := rapier.NewExecutor[testdata.Dict](db).
		Model().
		Where(refDict.Id.Eq(100)).
		UpdatesExpr(
			refDict.Key.Value("k1"),
		)
	_ = err          // return error
	_ = rowsAffected // return row affected
}
