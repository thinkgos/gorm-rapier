package main

import (
	"log"
	"os"
	"time"

	"github.com/things-go/gorm-assist/examples/model"

	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/utils/tests"
)

var db, _ = gorm.Open(tests.DummyDialector{}, nil)

func main() {
	db.Logger = logger.New(log.New(os.Stdout, "\r\n", log.LstdFlags), logger.Config{
		SlowThreshold:             200 * time.Millisecond,
		LogLevel:                  logger.Info,
		IgnoreRecordNotFoundError: false,
		Colorful:                  true,
	})
	db.DryRun = true

	xDict := model.X_Dict()
	xEntityDict := model.New_Dict(db)

	_, err := xEntityDict.Executor().
		Where(xDict.Id.Eq(100)).
		FindAll()
	checkError(err)

	_, err = xEntityDict.Executor().
		SelectExpr(
			xDict.Key,
			xDict.Name,
			xDict.IsPin,
			xDict.Remark,
		).
		Where(xDict.Id.Eq(100)).
		Updates(&model.Dict{
			Id:     0,
			Key:    "123",
			Name:   "456",
			IsPin:  false,
			Remark: "remark",
		})
	checkError(err)
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}
