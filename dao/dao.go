package dao

import (
	"github.com/HaleyLeoZhang/go-component/driver/db"
	"github.com/HaleyLeoZhang/go-component/driver/xlog"
	"github.com/jinzhu/gorm"
)

type Dao struct {
	db *gorm.DB
}

func New(cfg *db.Config) *Dao {
	var err error

	d := &Dao{}
	if d.db, err = db.New(cfg); err != nil {
		panic(err)
	}
	return d
}

func (d *Dao) Close() {
	err := d.db.Close()
	if err != nil {
		xlog.Errorf("curl_avatar.db.Err(%+v)", err)
	}
}
