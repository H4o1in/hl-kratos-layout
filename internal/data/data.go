package data

import (
	"gin-gorm-kratos-layout/internal/conf"

	"github.com/go-kratos/kratos/v3/log"
	"github.com/google/wire"

	"github.com/redis/go-redis/v9"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// ProviderSet is data providers.
var ProviderSet = wire.NewSet(NewData)

// Data holds the long-lived storage clients shared by repos.
type Data struct {
	db  *gorm.DB
	rdb *redis.Client
}

// NewData opens the database client and returns it with a cleanup function.
func NewData(c *conf.Data) (*Data, func(), error) {
	dc := c.GetDatabase()

	db, err := gorm.Open(mysql.New(mysql.Config{DSN: dc.Source}), &gorm.Config{})
	if err != nil {
		return nil, nil, err
	}
	if dc.GetDebug() {
		db = db.Debug()
	}

	rdb := redis.NewClient(&redis.Options{Addr: c.Redis.Addr})

	cleanup := func() {
		log.Info("closing the data resources")
		if err := rdb.Close(); err != nil {
			log.Error("closing the redis error: " + err.Error())
		}
	}
	return &Data{db: db, rdb: rdb}, cleanup, nil
}
