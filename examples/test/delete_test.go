package test

import (
	"testing"

	rapier "github.com/thinkgos/gorm-rapier"
	"github.com/thinkgos/gorm-rapier/examples/model"
)

func Test_Delete(t *testing.T) {
	refDict := model.Ref_Dict()
	rowsAffected, err := rapier.NewExecutor[model.Dict](db).
		Model().
		Where(refDict.Id.Eq(100)).
		Delete()
	_ = err          // return error
	_ = rowsAffected // return row affected
}
