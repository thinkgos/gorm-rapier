package test

import (
	"testing"

	rapier "github.com/thinkgos/gorm-rapier"
	"github.com/thinkgos/gorm-rapier/examples/model"
)

func Test_Create(t *testing.T) {
	newDict := model.Dict{
		Key:    "key1",
		Name:   "name1",
		IsPin:  true,
		Remark: "remark1",
	}
	err := rapier.NewExecutor[model.Dict](db).Create(&newDict)
	_ = err // return error

	newDicts := []*model.Dict{
		{
			Key:    "key1",
			Name:   "name1",
			IsPin:  true,
			Remark: "remark1",
		},
		{
			Key:    "key2",
			Name:   "name2",
			IsPin:  true,
			Remark: "remark2",
		},
	}
	err = rapier.NewExecutor[model.Dict](db).Create(newDicts...)
	_ = err // return error
}
