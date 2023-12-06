package mysqldb

import (
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func Open() (*gorm.DB, error) {
	dsn, _ := os.LookupEnv("MYSQL_DSN")

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		return nil, err
	}

	return db, nil
}

func Close(db *gorm.DB) error {
	sqlDB, err := db.DB()

	if err != nil {
		return err
	}

	return sqlDB.Close()
}
