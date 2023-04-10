package usecase

import (
	"github.com/dwlpra/todolist/domain/entity"
	"github.com/dwlpra/todolist/domain/repository"
)

type ActivityService struct {
	repo repository.ActivityRepository
}

func NewActvityService(repo repository.ActivityRepository) *ActivityService {
	return &ActivityService{repo}
}

func (s *ActivityService) GetAllActivities() (*[]entity.Activity, error) {
	return s.repo.FindAll()
}

func (s *ActivityService) GetActivityByID(id int) (*entity.Activity, error) {
	return s.repo.FindByID(id)
}

func (s *ActivityService) CreateActivity(activity *entity.Activity) error {
	return s.repo.Save(activity)
}

func (s *ActivityService) UpdateActivity(id int, activity *entity.Activity) (*entity.Activity, error) {
	s.repo.Update(id, activity)
	return s.repo.FindByID(id)
}

func (s *ActivityService) DeleteActivity(id int) error {
	return s.repo.Delete(id)
}
