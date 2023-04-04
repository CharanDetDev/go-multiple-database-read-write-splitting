package database

import (
	"context"
	"fmt"
	"time"

	"github.com/CharanDetDev/go-multiple-database-read-write-splitting/util/config"
	"github.com/CharanDetDev/go-multiple-database-read-write-splitting/util/logg"
	mysqlDSN "github.com/go-sql-driver/mysql"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/plugin/dbresolver"
)

var Conn *gorm.DB

type SqlLogger struct {
	logger.Interface
}

func InitDatabase() bool {

	gorm_db, err := initDatabaseGORM()
	if err != nil {
		logg.Printlogger("\t\t Database connection FAILED :: Connection GORM ", err)
		return false
	}
	Conn = gorm_db

	return true

}

func initDatabaseGORM() (db *gorm.DB, err error) {

	dsn_1 := mysqlDSN.Config{
		User:      config.Env.MYSQL_USERNAME,
		Passwd:    config.Env.MYSQL_PASSWORD,
		Net:       "tcp",
		Addr:      config.Env.MYSQL_HOSTNAME_1,
		DBName:    config.Env.MYSQL_DB_NAME,
		Loc:       time.Local,
		ParseTime: true,
	}
	dsn_2 := mysqlDSN.Config{
		User:      config.Env.MYSQL_USERNAME,
		Passwd:    config.Env.MYSQL_PASSWORD,
		Net:       "tcp",
		Addr:      config.Env.MYSQL_HOSTNAME_2,
		DBName:    config.Env.MYSQL_DB_NAME,
		Loc:       time.Local,
		ParseTime: true,
	}
	logg.Printlogger("\t\t ***** dsn_1 ***** :: | ", dsn_1.FormatDSN())
	logg.Printlogger("\t\t ***** dsn_2 ***** :: | ", dsn_2.FormatDSN())

	gorm_db, err := gorm.Open(
		mysql.Open(
			dsn_1.FormatDSN(),
		),
		&gorm.Config{
			Logger: &SqlLogger{},
			// Logger: logger.Default.LogMode(logger.Info),
			// DryRun: true,
		},
	)
	if err != nil {
		logg.Printlogger("\t\t ***** Connect failed (dsn_1) ***** :: | ", err.Error())
		return nil, err
	}

	err = gorm_db.Use(
		dbresolver.Register(
			dbresolver.Config{
				Replicas: []gorm.Dialector{mysql.Open(dsn_2.FormatDSN())},
			},
		).
			SetConnMaxIdleTime(time.Hour).
			SetConnMaxLifetime(24 * time.Hour).
			SetMaxIdleConns(100).
			SetMaxOpenConns(200),
	)
	if err != nil {
		logg.Printlogger("\t\t ***** Connect failed (dsn_2) ***** :: | ", err.Error())
		return nil, err
	}

	return gorm_db, nil
}

func ConnectionClose() {
	dbInstance, _ := Conn.DB()
	_ = dbInstance.Close()
}

// Colors
const (
	Reset       = "\033[0m"
	Red         = "\033[31m"
	Green       = "\033[32m"
	Yellow      = "\033[33m"
	Blue        = "\033[34m"
	Magenta     = "\033[35m"
	Cyan        = "\033[36m"
	White       = "\033[37m"
	BlueBold    = "\033[34;1m"
	MagentaBold = "\033[35;1m"
	RedBold     = "\033[31;1m"
	YellowBold  = "\033[33;1m"
)

func (sqlLog *SqlLogger) Trace(ctx context.Context, begin time.Time, fc func() (sql string, rowsAffected int64), err error) {
	sqlStatement, rowsAffected := fc()
	logg.Printlogger(fmt.Sprintf("%v***** SQL Statement ***** | ", Green), fmt.Sprintf("%v[ Row Affected : %v%v%v ] -> %v%v \n%v", BlueBold, YellowBold, rowsAffected, BlueBold, YellowBold, sqlStatement, Reset))
}
