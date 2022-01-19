package mysql

import (
	"fmt"
	"github.com/spf13/viper"
	"github.com/ulyyyyyy/tapd_notify/internal/logger"
	"github.com/ulyyyyyy/tapd_notify/internal/model/webhook_cfg"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"strings"
)

const (
	_cfgKeyDSN          = "db.mysql.dsn"
	_cfgKeyMaxIdleConns = "db.mysql.conns.max.idle"
	_cfgKeyMaxOpenConns = "db.mysql.conns.max.open"
)

var db *gorm.DB

func Initialize() (err error) {
	err = AutoMigrate(&webhook_cfg.WebhookCfg{})
	if err != nil {
		return
	}

	dsn := strings.TrimSpace(viper.GetString(_cfgKeyDSN))
	maxIdle := viper.GetInt(_cfgKeyMaxIdleConns)
	maxOpen := viper.GetInt(_cfgKeyMaxOpenConns)
	if len(dsn) == 0 {
		return fmt.Errorf("configs file maybe not loaded: %s", _cfgKeyDSN)
	}
	if maxIdle <= 0 {
		return fmt.Errorf("configs file maybe not as expected: %s", _cfgKeyMaxIdleConns)
	}
	if maxOpen <= 0 {
		return fmt.Errorf("configs file maybe not as expected: %s", _cfgKeyMaxOpenConns)
	}

	cfg := mysql.Config{
		DSN:                       dsn,   // DSN ( data source name )
		DefaultStringSize:         256,   // string 类型字段的默认长度
		DisableDatetimePrecision:  true,  // 禁用 datetime 精度，MySQL 5.6 之前的数据库不支持
		DontSupportRenameIndex:    true,  // 重命名索引时采用删除并新建的方式，MySQL 5.7 之前的数据库和 MariaDB 不支持重命名索引
		DontSupportRenameColumn:   true,  // 用 `change` 重命名列，MySQL 8 之前的数据库和 MariaDB 不支持重命名列
		SkipInitializeWithVersion: false, // 根据版本自动配置
	}

	if db, err = gorm.Open(mysql.New(cfg), &gorm.Config{
		Logger:                                   logger.GormLogger,
		DisableForeignKeyConstraintWhenMigrating: true,
	}); err != nil {
		return err
	}

	sqlDB, err := db.DB()
	if err != nil {
		return err
	}

	if err = sqlDB.Ping(); err != nil {
		return err
	}

	sqlDB.SetMaxIdleConns(maxIdle)
	sqlDB.SetMaxOpenConns(maxOpen)
	// sqlDB.SetConnMaxLifetime()

	return
}

func AutoMigrate(dst ...interface{}) error {
	return db.AutoMigrate(dst...)
}

func DB() *gorm.DB {
	return db
}

func DBDebug() *gorm.DB {
	return db.Debug()
}

func _db(tx ...*gorm.DB) *gorm.DB {
	_current := db
	if len(tx) != 0 && tx[0] != nil {
		_current = tx[0]
	}
	return _current
}

func Insert(v interface{}, tx ...*gorm.DB) error {
	return _db(tx...).Create(v).Error
}

func InsertMany(value interface{}, batchSize int) error {
	return db.CreateInBatches(value, batchSize).Error
}

func FindByID(id string, v interface{}) error {
	if err := db.Where("id = ?", id).Take(v).Error; err != nil {
		return err
	}
	return nil
}

func FindOneBy(filter map[string]interface{}, v interface{}, tx ...*gorm.DB) error {
	if err := _db(tx...).Where(filter).First(v).Error; err != nil {
		return err
	}
	return nil
}

func FindAll(v interface{}) error {
	if err := db.Find(v).Error; err != nil {
		return err
	}
	return nil
}

func FindAllBy(filter map[string]interface{}, v interface{}) error {
	if err := db.Where(filter).Find(v).Error; err != nil {
		return err
	}
	return nil
}

func ExistBy(filter map[string]interface{}, v interface{}, tx ...*gorm.DB) (existent bool, err error) {
	err = FindOneBy(filter, v, tx...)
	if err == gorm.ErrRecordNotFound {
		return false, nil
	}
	if err != nil {
		return false, err
	}
	return true, nil
}

func UpdateOneBy(filter, update map[string]interface{}, v interface{}, tx ...*gorm.DB) error {
	if err := FindOneBy(filter, v); err != nil {
		return err
	}
	if err := _db(tx...).Model(v).Updates(update).Error; err != nil {
		return err
	}
	return nil
}

func UpdateOneByModelID(v interface{}, update map[string]interface{}, tx ...*gorm.DB) error {
	if err := _db(tx...).Model(v).Updates(update).Error; err != nil {
		return err
	}
	return nil
}

func Save(v interface{}, tx ...*gorm.DB) error {
	if err := _db(tx...).Save(v).Error; err != nil {
		return err
	}
	return nil
}

func Delete(v interface{}, tx ...*gorm.DB) error {
	if err := _db(tx...).Delete(v).Error; err != nil {
		return err
	}
	return nil
}

func DeleteById(id string, v interface{}, tx ...*gorm.DB) error {
	return _db(tx...).Where("id", id).Delete(v).Error
}

func FindAllByDeleted(v interface{}) error {
	if err := db.Unscoped().Find(v).Error; err != nil {
		return err
	}
	return nil
}
