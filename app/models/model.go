package models

import (
	"fmt"
	"gotemplate/env"
	"sync"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	lock     *sync.Mutex = &sync.Mutex{}
	instance *gorm.DB

	// DSN (Data Source Name) for mysql
	dsn = fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		env.DatabaseUser,
		env.DatabasePassword,
		env.DatabaseHost,
		env.DatabasePort,
		env.DatabaseName,
	)

	// Gorm config
	config *gorm.Config = &gorm.Config{
		PrepareStmt: true,
	}

	// Mysql config
	mysqlConfig mysql.Config = mysql.Config{
		DSN:                       dsn,
		DefaultStringSize:         256,
		SkipInitializeWithVersion: false,
	}
)

// GetInstance returns a singleton instance of gorm.DB
func GetInstance() *gorm.DB {
	var err error

	if instance == nil {
		lock.Lock()
		defer lock.Unlock()

		if instance == nil {
			if env.DatabaseConnection == "mysql" {
				instance, err = gorm.Open(mysql.New(mysqlConfig), config)
			}

			if err != nil {
				panic(err)
			}
		}
	}

	return instance
}
