package persistance

import (
	"errors"

	"github.com/dwlpra/todolist/domain/entity"
	"github.com/dwlpra/todolist/domain/repository"
	"gorm.io/gorm"
)

type todoRepositoryImpl struct {
	dbMaster *gorm.DB
}

func NewTodo(dbMaster *gorm.DB) repository.TodoRepository {
	return &todoRepositoryImpl{dbMaster: dbMaster}
}

func (r *todoRepositoryImpl) FindAll(activityGroupID *string) (*[]entity.Todo, error) {
	var todo = new([]entity.Todo)
	if activityGroupID == nil {
		err := r.dbMaster.Find(todo).Error
		return todo, err
	} else {
		err := r.dbMaster.Where("activity_group_id = ?", activityGroupID).Find(&todo).Error
		return todo, err
	}
}

func (r *todoRepositoryImpl) FindByID(id int) (*entity.Todo, error) {
	var todo entity.Todo
	err := r.dbMaster.Where("id = ? ", id).First(&todo).Error
	return &todo, err
}

func (r todoRepositoryImpl) Save(todo *entity.Todo) error {
	return r.dbMaster.Save(todo).Error
}

func (r *todoRepositoryImpl) Update(id int, todo *entity.Todo) error {
	return r.dbMaster.Table("todos").Where("id = ?", id).Updates(todo).Error
}

func (r *todoRepositoryImpl) Delete(id int) error {
	result := r.dbMaster.Delete(&entity.Todo{}, id)
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return errors.New("not found")
	}
	return nil
}
