package main

import (
	"log"
	"os"
	"time"

	assist "github.com/things-go/gorm-assist"
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

	var rows []model.Dict

	xDict := model.X_Dict()
	db.Model(&model.Dict{}).
		Scopes(
			model.Xc_SelectDict("aaa"),
		).
		Where(xDict.Id.Eq(100)).
		Find(&rows)

	db.Model(&model.Dict{}).
		Scopes(
			assist.Select(
				xDict.Key,
				xDict.Name,
				xDict.IsPin,
				xDict.Sort,
			),
		).
		Where(xDict.Id.Eq(100)).
		Updates(&model.Dict{
			Id:    0,
			Key:   "",
			Name:  "",
			IsPin: false,
			Sort:  0,
		})
}
