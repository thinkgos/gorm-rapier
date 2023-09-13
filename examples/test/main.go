package main

import (
	"log"
	"os"
	"time"

	"github.com/thinkgos/gorm-rapier/examples/model"

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

	_, err := xDict.New_Executor(db).
		SelectExpr(
			xDict.Id,
			xDict.New_Executor(db).SelectExpr(xDict.Key).Where(xDict.Id.Eq(1)).IntoSubQueryExpr().As("aaa"),
		).
		Where(
			xDict.Id.Eq(100),
			xDict.Key.EqSubQuery(xDict.New_Executor(db).Where(xDict.Id.Eq(1)).IntoDB()),
		).
		FindAll()
	checkError(err)

	_, err = xDict.New_Executor(db).
		Where(
			xDict.Id.Eq(100),
		).
		UpdatesExpr(
			xDict.Key.Value("1000"),
			xDict.IsPin.Value(false),
		)
	checkError(err)
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}
