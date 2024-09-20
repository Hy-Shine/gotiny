package gormx

// docs: https://gorm.io
// github repo: https://github.com/go-gorm/gorm

import "gorm.io/gorm"

type user struct {
	gorm.Model
	Name    string `gorm:"type:varchar(50);index:idx_name;default:''" json:"name"`
	Age     int    `gorm:"type:int;default:0" json:"age"`
	Email   string `gorm:"type:varchar(255);not null;index:idx_email" json:"email"`
	Address string `gorm:"type:varchar(255);default:''" json:"address"`
}

const usersTableName = "users"

var ormDb *gorm.DB

type ormFunc func(*gorm.DB)

func (u *user) TableName() string {
	return usersTableName
}

func (u *user) createUser() error {
	return ormDb.Create(u).Error
}

func (u *user) createOrSelectUser() error {
	return ormDb.FirstOrCreate(u).Error
}

func (u *user) userInfoWithID(id uint) error {
	query := ormDb.Model(&user{}).Where("id = ?", id)
	err := query.First(u).Error
	return err
}

func (u *user) userInfoWithCondition() error {
	query := ormDb.Table(u.TableName())
	if u.ID != 0 {
		query.Where("id = ?", u.ID)
	}
	if u.Email != "" {
		query.Where("email = ?", u.Email)
	}

	return query.First(u).Error
}

func queryColumns(column string) ormFunc {
	return func(db *gorm.DB) {
		db.Select(column)
	}
}

func (u *user) userInfoWithFunc(fl ...ormFunc) {
	query := ormDb.Model(user{})
	for _, f := range fl {
		f(query)
	}
}

// deleteUserWithID soft delete
func (u *user) deleteUserWithID() error {
	return ormDb.Delete(u).Error
}

// deleteUserWithCondition delete user with condition
func (u *user) deleteUserWithCondition() error {
	ormDelete := ormDb.Table(u.TableName())
	if u.ID != 0 {
		ormDelete.Where("id = ?", u.ID)
	}
	if u.Email != "" {
		ormDelete.Where("email = ?", u.Email)
	}
	return ormDelete.Delete(user{}).Error
}

// hard delete: with gorm.Db Unscoped func
func (u *user) deleteUser() error {
	return ormDb.Unscoped().Delete(u).Error
}
