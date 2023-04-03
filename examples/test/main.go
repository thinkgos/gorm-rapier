package main

import (
	"log"
	"os"
	"time"

	gen "github.com/things-go/gorm-assist"
	"github.com/things-go/gorm-assist/examples/dict"
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

	dictmf := dict.NewDict()
	db.Model(&model.Dict{}).
		Scopes(gen.Select(
			dictmf.Id,
			dictmf.Key,
			dictmf.Name,
			dictmf.IsPin,
			dictmf.Sort,
			dictmf.CreatedAt.UnixTimestamp().IfNull(0).As("created_at"),
		)).
		Where(dictmf.Id.Eq(100)).
		Find(&rows)

	db.Model(&model.Dict{}).
		Scopes(gen.Select(
			dictmf.Key,
			dictmf.Name,
			dictmf.IsPin,
			dictmf.Sort,
		)).
		Where(dictmf.Id.Eq(100)).
		Updates(&model.Dict{
			Id:    0,
			Key:   "",
			Name:  "",
			IsPin: false,
			Sort:  0,
		})

}
