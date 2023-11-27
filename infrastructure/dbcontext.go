package infrastructure

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type DbContext struct {
	_db              *gorm.DB
	ConnectionString string
}

func NewDbContext(connectionString string) *DbContext {
	if db, err := gorm.Open(mysql.Open(connectionString), &gorm.Config{}); err != nil {
		panic(err)
	} else {
		return &DbContext{
			_db:              db,
			ConnectionString: connectionString,
		}
	}

}
