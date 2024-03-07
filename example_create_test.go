package rapier_test

import (
	"testing"

	rapier "github.com/thinkgos/gorm-rapier"
	"github.com/thinkgos/gorm-rapier/testdata"
)

func Test_Example_EmptyRecord_Create(t *testing.T) {
	// empty record
	err := rapier.NewExecutor[testdata.Dict](db).Create()
	_ = err // return error
}
func Test_Example_SingleRecord_Create(t *testing.T) {
	// single record
	newDict := testdata.Dict{
		Key:    "key1",
		Name:   "name1",
		IsPin:  true,
		Remark: "remark1",
	}
	err := rapier.NewExecutor[testdata.Dict](db).Create(&newDict)
	_ = err // return error
}
func Test_Example_Multiple_Create(t *testing.T) {
	// multiple record
	newDicts := []*testdata.Dict{
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
	err := rapier.NewExecutor[testdata.Dict](db).Create(newDicts...)
	_ = err // return error
}

func Test_Example_Batch_CreateInBatch(t *testing.T) {
	// batch insert multiple record
	newDicts := []*testdata.Dict{
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
	err := rapier.NewExecutor[testdata.Dict](db).CreateInBatches(newDicts, 2)
	_ = err // return error
}
