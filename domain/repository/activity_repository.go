package repository

import "github.com/dwlpra/todolist/domain/entity"

type ActivityRepository interface {
	FindAll() (*[]entity.Activity, error)
	FindByID(int) (*entity.Activity, error)
	Save(settlementConfiguration *entity.Activity) error
	Update(int, *entity.Activity) error
	Delete(int) error
}
