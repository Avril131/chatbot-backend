package bootstrap

import (
	"chatbot-backend/global"
	"chatbot-backend/app/models"
	"fmt"
	"io"
	"log"
	"os"
	"time"

	"go.uber.org/zap"
	"gopkg.in/natefinch/lumberjack.v2"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func InitializeDB() *gorm.DB {
	switch global.App.Config.Database.Driver {
	case "portgres":
		return initPostgreSqlGorm()
	default:
		return initPostgreSqlGorm()
	}
}

func initPostgreSqlGorm() *gorm.DB {

	dbConfig := global.App.Config.Database
	if dbConfig.Database == "" {
		fmt.Errorf("do not config db\n")
		return nil
	}

	dsn := fmt.Sprintf("postgres://%s:%s@%s:%d/%s?sslmode=disable",
        dbConfig.UserName,
        dbConfig.Password,
        dbConfig.Host,
        dbConfig.Port,
        dbConfig.Database,
    )

	postgresConfig := postgres.Config{
		DSN: dsn, // DSN data source name
	}

	if db, err := gorm.Open(postgres.New(postgresConfig), &gorm.Config{
		Logger: getGormLogger(),
	}); err != nil {
		global.App.Log.Error("sql connect failed, err:", zap.Any("err", err))
		return nil
	} else {
		sqlDB, _ := db.DB()
		sqlDB.SetMaxIdleConns(dbConfig.MaxIdleConns)
		sqlDB.SetMaxOpenConns(dbConfig.MaxOpenConns)
		initTables(db)
		return db
	}
}

// init tables in db
func initTables(db *gorm.DB) {
	err := db.AutoMigrate(
		models.User{},
	)
	if err != nil {
		global.App.Log.Error("migrate table failed", zap.Any("err", err))
		os.Exit(0)
	}
}

func getGormLogger() logger.Interface {
	var logMode logger.LogLevel

	switch global.App.Config.Database.LogMode {
	case "silent":
		logMode = logger.Silent
	case "error":
		logMode = logger.Error
	case "warn":
		logMode = logger.Warn
	case "info":
		logMode = logger.Info
	default:
		logMode = logger.Info
	}

	return logger.New(getGormLogWriter(), logger.Config{
		SlowThreshold:             200 * time.Millisecond,
		LogLevel:                  logMode,
		IgnoreRecordNotFoundError: false,
		Colorful:                  !global.App.Config.Database.EnableFileLogWriter,
	})
}

// define gorm Writer
func getGormLogWriter() logger.Writer {
	var writer io.Writer

	// config db log
	if global.App.Config.Database.EnableFileLogWriter {
		writer = &lumberjack.Logger{
			Filename:   global.App.Config.Log.RootDir + "/" + global.App.Config.Database.LogFilename,
			MaxSize:    global.App.Config.Log.MaxSize,
			MaxBackups: global.App.Config.Log.MaxBackups,
			MaxAge:     global.App.Config.Log.MaxAge,
			Compress:   global.App.Config.Log.Compress,
		}
	} else {
		writer = os.Stdout
	}
	return log.New(writer, "\r\n", log.LstdFlags)
}
