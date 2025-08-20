package di

import (
	"assessment/data/postgres/clinet"
	"fmt"

	"assessment/pkg"
	"github.com/sirupsen/logrus"
	gormPostgres "gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var (
	db           *gorm.DB
	dbDatasource *clinet.Postgres
)

func DB() *gorm.DB {
	if db != nil {
		return db
	}

	connection := pkg.Config.Postgres
	dsn := fmt.Sprintf("host=%v user=%v password=%v dbname=%v port=%v sslmode=disable TimeZone=%v",
		connection.Host, connection.User, connection.Password, connection.Db, connection.Port, connection.Timezone)
	var err error
	db, err = gorm.Open(gormPostgres.Open(dsn), &gorm.Config{Logger: logger.Default.LogMode(logger.Silent)})
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"err": err,
		}).Panic("failed to connect Postgres!")
	}

	d, err := db.DB()
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"err": err,
		}).Panic("failed to get DB instance")
	}
	if err = d.Ping(); err != nil {
		logrus.WithFields(logrus.Fields{
			"err": err,
		}).Panic("database has no ping")
	}

	return db
}

func DbDatasource() *clinet.Postgres {
	if dbDatasource == nil {
		dbDatasource = clinet.NewPostgres(DB())
	}

	return dbDatasource
}
