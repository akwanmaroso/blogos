package databases

import (
	"github.com/akwanmaroso/blogos/config"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func Connect() (*gorm.DB, error) {
	db, err := gorm.Open(config.DbDriver, config.DbUrl)
	if err != nil {
		return nil, err
	}
	return db, nil
}
