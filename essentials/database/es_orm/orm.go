package es_orm

import (
	"github.com/watermint/toolbox/essentials/database/es_orm_logger"
	"github.com/watermint/toolbox/essentials/log/esl"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func NewOrm(l esl.Logger, path string) (*gorm.DB, error) {
	return newOrmWithConfig(sqlite.Open(path),
		&gorm.Config{
			Logger: es_orm_logger.NewGormLogger(l),
		},
	)
}

func NewOrmOnMemory(l esl.Logger) (*gorm.DB, error) {
	return newOrmWithConfig(sqlite.Open("file::memory:"),
		&gorm.Config{
			Logger: es_orm_logger.NewGormLogger(l),
		},
	)
}

func newOrmWithConfig(db gorm.Dialector, config *gorm.Config) (*gorm.DB, error) {
	d, err := gorm.Open(db, config)
	if err != nil {
		return nil, err
	}
	//ddb, err := d.DB()
	//if err != nil {
	//	return nil, err
	//}
	//ddb.SetMaxOpenConns(1)
	return d, nil
}
