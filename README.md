# The Go Language Programming : Multiple databases, read-write splitting for GORM

> ## **Package**
>
> _Fiber, GORM, dbresolver, MySQL, Viper_

- github.com/gofiber/fiber/v2
- gorm.io/plugin/dbresolver
- gorm.io/gorm
- gorm.io/driver/mysql
- gorm.io/gorm/logger
- github.com/go-sql-driver/mysql
- github.com/spf13/viper
- database สามารถ git clone เพื่อลองใช้งานได้ที่ https://github.com/CharanDetDev/docker-mysql-database-1 (Read/Write), https://github.com/CharanDetDev/docker-mysql-database-2 docker-mysql-database-2(Read-only)

> ## **Example**

```golang

    dsn_1 := mysqlDSN.Config{
		User:      "username",
		Passwd:    "password",
		Net:       "tcp",
		Addr:      "hostname:port",
		DBName:    "databaseName",
		Loc:       time.Local,
		ParseTime: true,
	}
    dsn_2 := mysqlDSN.Config{
		User:      "username",
		Passwd:    "password",
		Net:       "tcp",
		Addr:      "hostname:port",
		DBName:    "databaseName",
		Loc:       time.Local,
		ParseTime: true,
	}

    gorm_db, _ := gorm.Open(
		mysql.Open(
			dsn_1.FormatDSN(), // Data source name database Read/Write
		),
		&gorm.Config{
			Logger: &SqlLogger{},
		},
	)

    gorm_db.Use(
		dbresolver.Register(
			dbresolver.Config{
				Replicas: []gorm.Dialector{mysql.Open(dsn_2.FormatDSN())}, // Data source name database Read Only
			},
		),
	)
    
```
