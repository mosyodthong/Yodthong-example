package config

import (
	"context"
	"fmt"
	"os"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type SqlLogger struct {
	logger.Interface
}

func InitDatabase() *gorm.DB {
	dsn := fmt.Sprintf("host=%v user=%v password=%v dbname=%v port=5432 sslmode=disable TimeZone=Asia/Bangkok",
		os.Getenv("PG_HOST"),
		os.Getenv("PG_USER"),
		os.Getenv("PG_PASSWORD"),
		os.Getenv("PG_NAME"),
	)
	var err error
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		//Logger: &SqlLogger{},
		// DryRun: true,
	})
	if err != nil {
		panic(err)
	}
	return db
}

func (l SqlLogger) Trace(ctx context.Context, begin time.Time, fc func() (sql string, rowsAffected int64), err error) {
	sql, _ := fc()
	fmt.Printf("\n ------------------------------------------------------------------  \n%v\n  ------------------------------------------------------------------ \n", sql)
}
