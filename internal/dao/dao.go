package dao

import (
	"errors"
	"fmt"
	"maya/configs"
	"maya/internal/model"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
)

var (
	// DBEngine *gorm.DB
	Db *gorm.DB
)

// func NewDBEngine(config *configs.Config) (*gorm.DB, error) {
func NewDBEngine(config *configs.Config) error {
	dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=%s&parseTime=%t&loc=Local",
		config.Database.UserName,
		config.Database.Password,
		config.Database.Host,
		config.Database.DbName,
		config.Database.Charset,
		config.Database.ParseTime,
	)

	if config.Database.DbType != "mysql" {
		// return nil, errors.New("暂时只支持mysql")
		return errors.New("暂时只支持mysql")
	}

	database, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			TablePrefix:   config.Database.TablePrefix,
			SingularTable: true,
		},
	})
	if err != nil {
		// return nil, err
		return err
	}

	if config.Server.RunMode == "debug" {
		database.Logger.LogMode(logger.Info)
	}

	sqlDB, err := database.DB()
	if err != nil {
		return err
	}

	// SetMaxIdleConns sets the maximum number of connections in the idle connection pool.
	sqlDB.SetMaxIdleConns(config.Database.MaxIdleConns)
	sqlDB.SetMaxOpenConns(config.Database.MaxOpenConns)

	// SetConnMaxLifetime sets the maximum amount of time a connection may be reused.
	sqlDB.SetConnMaxLifetime(time.Hour)

	// DBEngine = database
	Db = database

	Db.AutoMigrate(&model.User{})
	fmt.Println("Database Migrated")

	return nil
	// return db, nil
}
