package rapier_test

import (
	"testing"

	rapier "github.com/thinkgos/gorm-rapier"
	"github.com/thinkgos/gorm-rapier/testdata"
)

func Test_Delete(t *testing.T) {
	refDict := testdata.Ref_Dict()
	rowsAffected, err := rapier.NewExecutor[testdata.Dict](db).
		Model().
		Where(refDict.Id.Eq(100)).
		Delete()
	_ = err          // return error
	_ = rowsAffected // return row affected
}
