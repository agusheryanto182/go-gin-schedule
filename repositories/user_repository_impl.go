package repositories

import (
	"github.com/agusheryanto182/go-schedule/models/domain"
	"gorm.io/gorm"
)

type UserRepositoryImpl struct {
	db *gorm.DB
}

func (self *UserRepositoryImpl) Save(user domain.User) (domain.User, error) {
	err := self.db.Create(&user).Error
	if err != nil {
		return user, err
	}
	return user, nil
}

func (self *UserRepositoryImpl) FindById(userId int) (domain.User, error) {
	var user domain.User
	err := self.db.Where("user_id = ?", userId).Find(&user).Error
	if err != nil {
		return user, err
	}
	return user, nil
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &UserRepositoryImpl{db: db}
}
