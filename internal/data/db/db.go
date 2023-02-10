package db

import (
	"github.com/google/wire"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	gormLogger "gorm.io/gorm/logger"
)

var ProvideSet = wire.NewSet(New)

func New(opts []Option) (*gorm.DB, error) {
	op := options{
		Username: "root",
		Password: "root",
		Host:     "localhost",
		Port:     3306,
		Database: "test",
	}
	for _, o := range opts {
		o.apply(&op)
	}
	db, err := gorm.Open(mysql.Open(op.ConnectionString()), &gorm.Config{
		Logger:               gormLogger.Default.LogMode(gormLogger.Info),
		DisableAutomaticPing: false,
	})
	if err != nil {
		return nil, err
	}
	return db, err
}
