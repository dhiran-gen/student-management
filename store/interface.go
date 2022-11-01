package store

import (
	"github.com/student-management/models"
)

type Repository interface {
	GetById(id int) (*models.Student, error)
	GetAll() ([]*models.Student, error)
	Create( models.Student) error
	Update( models.Student) error
	Delete(id int) error
}
