package dao

import (
	"context"
	"github.com/HaleyLeoZhang/email_server/common/model/po"
	"github.com/jinzhu/gorm"
	"github.com/pkg/errors"
)

func (d *Dao) EmailInsert(ctx context.Context, tx *gorm.DB, row *po.Email) (err error) {
	if tx == nil {
		tx = d.db.Begin()
		defer func() {
			if r := recover(); r != nil {
				tx.Rollback()
				return
			}
			if err != nil {
				tx.Rollback()
				return
			}
			err = errors.WithStack(tx.Commit().Error)
		}()
	}

	if tx.NewRecord(row) {
		err = errors.WithStack(tx.Create(row).Error)
	}

	return
}

func (d *Dao) EmailUpdate(ctx context.Context, tx *gorm.DB, where map[string]interface{}, update map[string]interface{}) (affected int64, err error) {
	if tx == nil {
		tx = d.db.Begin()
		defer func() {
			if r := recover(); r != nil {
				tx.Rollback()
				return
			}
			if err != nil {
				tx.Rollback()
				return
			}
			err = errors.WithStack(tx.Commit().Error)
		}()
	}
	email := &po.Email{}
	db := tx.Table(email.TableName()).Where(where).Updates(update)
	if db.Error != nil {
		err = errors.WithStack(db.Error)
		return
	}

	affected = db.RowsAffected

	return
}
