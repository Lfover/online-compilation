package data

import (
	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
	"gorm.io/gorm"
	"online-compilation/internal/conf"
)

// ProviderSet is data providers.
var ProviderSet = wire.NewSet(NewData, NewGreeterRepo, NewNoteRepo) //注入位置

// Data .
type Data struct {
	db  *gorm.DB
	log *log.Helper
}

// NewData .
// 这里就是建立数据库连接，你可以用本地测试一下，把想这些注释放开就行
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
