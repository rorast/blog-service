package model

import (
	"fmt"
	otgorm "github.com/eddycjy/opentracing-gorm"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/rorast/blog-service/global"
	"github.com/rorast/blog-service/pkg/setting"
	"time"
)

const (
	STATE_OPEN  = 1
	STATE_CLOSE = 0
)

type Model struct {
	ID         uint32 `gorm:"primary_key" json:"id"`
	CreatedBy  string `json:"created_by"`
	ModifiedBy string `json:"modified_by"`
	CreatedOn  uint32 `json:"created_on"`
	ModifiedOn uint32 `json:"modified_on"`
	DeletedOn  uint32 `json:"deleted_on"`
	IsDel      uint8  `json:"is_del"`
}

func NewDBEngine(databaseSetting *setting.DatabaseSettingS) (*gorm.DB, error) {
	s := "%s:%s@tcp(%s)/%s?charset=%s&parseTime=%t&loc=Local"
	db, err := gorm.Open(databaseSetting.DBType, fmt.Sprintf(s,
		databaseSetting.UserName,
		databaseSetting.Password,
		databaseSetting.Host,
		databaseSetting.DBName,
		databaseSetting.Charset,
		databaseSetting.ParseTime,
	))
	if err != nil {
		return nil, err
	}

	if global.ServerSetting.RunMode == "debug" {
		db.LogMode(true)
	}
	db.SingularTable(true)
	db.Callback().Create().Replace("gorm:update_time_stamp", updateTimeStampForCreateCallback)
	db.Callback().Update().Replace("gorm:update_time_stamp", updateTimeStampForUpdateCallback)
	db.Callback().Delete().Replace("gorm:delete", deleteCallback)
	db.DB().SetMaxIdleConns(databaseSetting.MaxIdleConns)
	db.DB().SetMaxOpenConns(databaseSetting.MaxOpenConns)
	otgorm.AddGormCallbacks(db)
	return db, nil
}

func updateTimeStampForCreateCallback(scope *gorm.Scope) {
	if !scope.HasError() {
		nowTime := time.Now().Unix()
		if createTimeField, ok := scope.FieldByName("CreatedOn"); ok {
			if createTimeField.IsBlank {
				_ = createTimeField.Set(nowTime)
			}
		}

		if modifyTimeField, ok := scope.FieldByName("ModifiedOn"); ok {
			if modifyTimeField.IsBlank {
				_ = modifyTimeField.Set(nowTime)
			}
		}
	}
}

func updateTimeStampForUpdateCallback(scope *gorm.Scope) {
	if _, ok := scope.Get("gorm:update_column"); !ok {
		_ = scope.SetColumn("ModifiedOn", time.Now().Unix())
	}
}

func deleteCallback(scope *gorm.Scope) {
	if !scope.HasError() {
		var extraOption string
		if str, ok := scope.Get("gorm:delete_option"); ok {
			extraOption = fmt.Sprint(str)
		}

		deletedOnField, hasDeletedOnField := scope.FieldByName("DeletedOn")
		isDelField, hasIsDelField := scope.FieldByName("IsDel")
		if !scope.Search.Unscoped && hasDeletedOnField && hasIsDelField {
			now := time.Now().Unix()
			scope.Raw(fmt.Sprintf(
				"UPDATE %v SET %v=%v,%v=%v%v%v",
				scope.QuotedTableName(),
				scope.Quote(deletedOnField.DBName),
				scope.AddToVars(now),
				scope.Quote(isDelField.DBName),
				scope.AddToVars(1),
				addExtraSpaceIfExist(scope.CombinedConditionSql()),
				addExtraSpaceIfExist(extraOption),
			)).Exec()
		} else {
			scope.Raw(fmt.Sprintf(
				"DELETE FROM %v%v%v",
				scope.QuotedTableName(),
				addExtraSpaceIfExist(scope.CombinedConditionSql()),
				addExtraSpaceIfExist(extraOption),
			)).Exec()
		}
	}
}

func addExtraSpaceIfExist(str string) string {
	if str != "" {
		return " " + str
	}
	return ""
}

// NewDBEngine 建立 GORM v2 資料庫連線
//func NewDBEngine(databaseSetting *setting.DatabaseSettingS) (*gorm.DB, error) {
//	// DSN (Data Source Name) 格式
//	dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=%s&parseTime=%t&loc=Local",
//		databaseSetting.UserName,
//		databaseSetting.Password,
//		databaseSetting.Host,
//		databaseSetting.DBName,
//		databaseSetting.Charset,
//		databaseSetting.ParseTime,
//	)
//
//	// 設定 Logger
//	var gormLogger logger.Interface
//	if global.ServerSetting.RunMode == "debug" {
//		gormLogger = logger.Default.LogMode(logger.Info) // 設定 Debug 日誌
//	} else {
//		gormLogger = logger.Default.LogMode(logger.Silent)
//	}
//
//	// 連線到資料庫
//	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
//		Logger: gormLogger,
//	})
//	if err != nil {
//		return nil, err
//	}
//
//	// 取得底層 sql.DB 物件
//	sqlDB, err := db.DB()
//	if err != nil {
//		return nil, err
//	}
//
//	// 設定連線池
//	sqlDB.SetMaxIdleConns(databaseSetting.MaxIdleConns)
//	sqlDB.SetMaxOpenConns(databaseSetting.MaxOpenConns)
//	sqlDB.SetConnMaxLifetime(time.Hour) // 可根據需求設定
//
//	// 註冊回調函數（GORM v2 新用法）
//	db.Callback().Create().Before("gorm:create").Register("update_time_stamp", updateTimeStampForCreateCallback)
//	db.Callback().Update().Before("gorm:update").Register("update_time_stamp", updateTimeStampForUpdateCallback)
//	db.Callback().Delete().Before("gorm:delete").Register("delete_callback", deleteCallback)
//
//	// otgorm 的支援（如果仍然需要）
//	otelgorm.AddGormCallbacks(db)
//
//	return db, nil
//}
