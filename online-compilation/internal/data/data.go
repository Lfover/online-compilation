package data

import (
	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
	"gorm.io/gorm"
	"online-compilation/internal/conf"
)

// ProviderSet is data providers.
var ProviderSet = wire.NewSet(NewData, NewGreeterRepo, NewNoteRepo)

// Data .
type Data struct {
	db  *gorm.DB
	log *log.Helper
}

// NewData .
func NewData(c *conf.Data, logger log.Logger) (*Data, func(), error) {
	//db, err := gorm.Open(mysql.Open(c.Database.Source), &gorm.Config{})
	//if err != nil {
	//	return nil, nil, err
	//}
	//if err := db.Use(otelgorm.NewPlugin()); err != nil {
	//	return nil, nil, errors.Wrap(err, "data: db.Use error")
	//}

	d := &Data{
		//db:  db,
		log: log.NewHelper(logger),
	}
	return d, func() {
		_db, err := d.db.DB()
		if err != nil {
			log.NewHelper(logger).Errorf("database close err:%+v", err)
		}
		_ = _db.Close()
		log.NewHelper(logger).Info("closing the mysql")
	}, nil
}
