package service

import (
	"errors"
	"unit-test/entity"
	"unit-test/repository"
)

type CategoryService struct {
	Repository repository.CategoryRepository
}

func (service CategoryService) Get(id string) (*entity.Category, error) {
	category := service.Repository.FindById(id)
	if category == nil {
		return category, errors.New("Category not found")
	} else {
		return category, nil
	}
}