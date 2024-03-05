package test

import (
	"testing"

	rapier "github.com/thinkgos/gorm-rapier"
	"github.com/thinkgos/gorm-rapier/examples/model"
)

func Test_Create(t *testing.T) {
	// single record
	newDict := model.Dict{
		Key:    "key1",
		Name:   "name1",
		IsPin:  true,
		Remark: "remark1",
	}
	err := rapier.NewExecutor[model.Dict](db).Create(&newDict)
	_ = err // return error

	// multiple record
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

func Test_CreateInBatch(t *testing.T) {
	// multiple record
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
		{
			Key:    "key3",
			Name:   "name3",
			IsPin:  true,
			Remark: "remark3",
		},
	}
	err := rapier.NewExecutor[model.Dict](db).CreateInBatches(newDicts, 2)
	_ = err // return error
}
