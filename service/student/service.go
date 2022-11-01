package students

import (
	"errors"
	"regexp"
	
	"github.com/student-management/store"
	"github.com/student-management/models"
)

type studentService struct {
	usrStoreHandler store.Repository
}

func New(store store.Repository) *studentService {
	return &studentService{
		usrStoreHandler: store,
	}
}

// validateId utility to validate the id
func validateId(id int) bool {
	// check if id is of type int and is a positive number
	return id > 0
}

// validateEmail utility to validate email ids
func validateEmail(email string) bool {
	emailRegex := regexp.MustCompile(`^[a-z0-9._%+\-]+@[a-z0-9.\-]+\.[a-z]{2,4}$`)
	return emailRegex.MatchString(email)
}

// validatePhone utility to validate phone number
func validatePhone(phone string) bool {
	phoneRegex := regexp.MustCompile(`^[+]*[(]{0,1}[0-9]{1,4}[)]{0,1}[-\s\./0-9]*$`)
	return phoneRegex.MatchString(phone) && len(phone) <= 10
}

func (us *studentService) GetById(id int) (*models.Student, error) {
	// check if id is valid
	if validateId(id) {
		resp, err := us.usrStoreHandler.GetById(id)
		if err != nil {
			return &models.Student{}, err
		}
		return resp, nil
	}
	return &models.Student{}, errors.New("error invalid id type")
}

func (us *studentService) GetAll() ([]*models.Student, error) {
	return us.usrStoreHandler.GetAll()
}

func (us *studentService) Create(student models.Student) error {
	// validate id, email and phone
	if !validateId(student.Id) {
		return errors.New("error invalid id")
	}
	if !validateEmail(student.Email) {
		return errors.New("error invalid email")
	}
	if !validatePhone(student.Phone) {
		return errors.New("error invalid phone")
	}
	return us.usrStoreHandler.Create(student)
}

func (us *studentService) Update(student models.Student) error {
	// Validate id, email and phone only if they are to be updated
	if !validateId(student.Id) {
		return errors.New("error invalid id")
	}
	if student.Email != "" && !validateEmail(student.Email) {
		return errors.New("error invalid email")
	}
	if student.Phone != "" && !validatePhone(student.Phone) {
		return errors.New("error invalid phone")
	}
	return us.usrStoreHandler.Update(student)
}

func (us *studentService) Delete(id int) error {
	// validation of id
	if !validateId(id) {
		return errors.New("error invalid id")
	}
	return us.usrStoreHandler.Delete(id)
}
