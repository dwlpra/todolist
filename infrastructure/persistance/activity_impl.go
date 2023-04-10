package persistance

import (
	"errors"

	"github.com/dwlpra/todolist/domain/entity"
	"github.com/dwlpra/todolist/domain/repository"
	"gorm.io/gorm"
)

type activityRepositoryImpl struct {
	dbMaster *gorm.DB
}

func NewActivity(dbMaster *gorm.DB) repository.ActivityRepository {
	return &activityRepositoryImpl{dbMaster: dbMaster}
}

func (r *activityRepositoryImpl) FindAll() (*[]entity.Activity, error) {
	var activities = new([]entity.Activity)
	err := r.dbMaster.Find(activities).Error
	return activities, err

}

func (r *activityRepositoryImpl) FindByID(id int) (*entity.Activity, error) {
	var activity = new(entity.Activity)
	var activities = new([]entity.Activity)
	err := r.dbMaster.Where("id = ?", id).First(activities).Error
	return activity, err

}

func (r activityRepositoryImpl) Save(activity *entity.Activity) error {
	return r.dbMaster.Save(activity).Error
}

func (r *activityRepositoryImpl) Update(id int, activity *entity.Activity) error {
	return r.dbMaster.Table("activities").Where("id = ?", id).Updates(activity).Error
}

func (r *activityRepositoryImpl) Delete(id int) error {
	result := r.dbMaster.Delete(&entity.Activity{}, id)
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return errors.New("not found")
	}
	return nil
}
