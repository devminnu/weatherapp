package app

import (
	"time"

	"github.com/devminnu/weatherapp/location/config"
	"github.com/jmoiron/sqlx"
	"go.uber.org/zap"
)

var (
	db     *sqlx.DB
	logger *zap.SugaredLogger
)

func Init() {
	InitLogger()

	err := initDB()
	if err != nil {
		panic(err)
	}
}

func InitLogger() {
	zapLogger, err := zap.NewProduction()
	if err != nil {
		panic(err)
	}

	logger = zapLogger.Sugar()
}

func initDB() (err error) {
	dbConfig := config.Database()

	db, err = sqlx.Open(dbConfig.Driver(), dbConfig.ConnectionURL())
	if err != nil {
		return
	}

	if err = db.Ping(); err != nil {
		return
	}

	db.SetMaxIdleConns(dbConfig.MaxPoolSize())
	db.SetMaxOpenConns(dbConfig.MaxOpenConns())
	db.SetConnMaxLifetime(time.Duration(dbConfig.MaxLifeTimeMins()) * time.Minute)

	return
}

func GetDB() *sqlx.DB {
	return db
}

func GetLogger() *zap.SugaredLogger {
	return logger
}

func Close() {
	logger.Sync()
	db.Close()
}
