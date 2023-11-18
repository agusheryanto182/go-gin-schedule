package repositories

import (
	"github.com/agusheryanto182/go-todo/models/domain"
	"gorm.io/gorm"
)

type ActivityRepositoryImpl struct {
	db *gorm.DB
}

func (r *ActivityRepositoryImpl) Save(activity domain.Activity) (domain.Activity, error) {
	err := r.db.Create(&activity).Error
	if err != nil {
		return activity, err
	}
	return activity, nil
}

func (r *ActivityRepositoryImpl) Update(activity domain.Activity) (domain.Activity, error) {
	err := r.db.Save(&activity).Error
	if err != nil {
		return activity, err
	}
	return activity, nil
}

func (r *ActivityRepositoryImpl) Delete(Id int) error {
	err := r.db.Delete(&domain.Activity{}, "activity_id = ?", Id).Error
	if err != nil {
		return err
	}
	return nil
}

func (r *ActivityRepositoryImpl) FindById(Id int) (domain.Activity, error) {
	var activity domain.Activity
	err := r.db.Where("activity_id  = ?", Id).Find(&activity).Error
	if err != nil {
		return activity, err
	}
	return activity, nil
}

func (r *ActivityRepositoryImpl) FindAll() ([]domain.Activity, error) {
	var activity []domain.Activity
	err := r.db.Find(&activity).Error
	if err != nil {
		return activity, err
	}
	return activity, nil
}

func NewActivityRepository(db *gorm.DB) ActivityRepository {
	return &ActivityRepositoryImpl{db: db}
}
