package orm

import (
	"github.com/watermint/toolbox/essentials/database/orm_logger"
	"github.com/watermint/toolbox/essentials/log/esl"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func NewOrm(l esl.Logger, path string) (*gorm.DB, error) {
	return newOrmWithConfig(sqlite.Open(path),
		&gorm.Config{
			Logger: orm_logger.NewGormLogger(l),
		},
	)
}

func NewOrmOnMemory(l esl.Logger) (*gorm.DB, error) {
	return newOrmWithConfig(sqlite.Open("file::memory:?cache=shared"),
		&gorm.Config{
			Logger: orm_logger.NewGormLogger(l),
		},
	)
}

func newOrmWithConfig(db gorm.Dialector, config *gorm.Config) (*gorm.DB, error) {
	return gorm.Open(db, config)
}
