package students

import (
	"database/sql"
	"errors"
	"log"

	"github.com/student-management/models"
)

type store struct {
	db *sql.DB
}

func New(db *sql.DB) *store {
	return &store{
		db: db,
	}
}

// GetById Retrieve a single student from database
func (s *store) GetById(id int) (*models.Student, error) {
	var student models.Student

	//define the query
	query := "select id, name, email, phone, age from student where id = ?"

	// Execute the query and return the student struct and error if any
	err := s.db.QueryRow(query, id).Scan(&student.Id, &student.Name, &student.Email, &student.Phone, &student.Age)
	if err != nil {
		return &models.Student{}, errors.New("error fetching from database, id not found")
	}
	return &student, err
}

// GetAll Retrieve all the students form the database
func (s *store) GetAll() ([]*models.Student, error) {
	var student []*models.Student

	// define the query
	query := "select id,name,email,phone,age from student"
	rows, err := s.db.Query(query)
	if err != nil {
		return []*models.Student{}, errors.New("error fetching data from database")
	}

	// iterate through the result and add it to the student slice
	for rows.Next() {
		var u models.Student
		_ = rows.Scan(&u.Id, &u.Name, &u.Email, &u.Phone, &u.Age)
		student = append(student, &u)
	}
	return student, nil
}

// Create method Create a new student entry
func (s *store) Create(u models.Student) error {
	// define the query
	query := "insert into student values (?,?,?,?,?)"

	_, err := s.db.Exec(query, u.Id, u.Name, u.Email, u.Phone, u.Age)
	if err != nil {
		log.Print(err)
		return err
	}
	return nil
}

// Update Update the fields based on the input
func (s *store) Update(u models.Student) error {
	fields, values := formQuery(u)
	// form the query with the fields
	query := "update student set" + fields + " where id = ?"
	_, err := s.db.Exec(query, values...)
	if err != nil {
		return errors.New("error, no id provided, cannot update")
	}
	return nil
}

// Delete Delete record from database based on input id
func (s *store) Delete(id int) error {
	// define the delete query
	query := "delete from student where id = ?"

	_, err := s.db.Exec(query, id)
	if err != nil {
		return errors.New("error, not able to delete data")
	}
	return nil
}


func formQuery(u models.Student) (string, []interface{}) {
	// declare a variable to hold query to be updated
	var query string
	var values []interface{}

	if u.Id < 0 {
		return "", nil
	}
	if u.Name != "" {
		query += " name = ?,"
		values = append(values, u.Name)
	}
	if u.Email != "" {
		query += " email = ?,"
		values = append(values, u.Email)
	}
	if u.Phone != "" {
		query += " phone = ?,"
		values = append(values, u.Phone)
	}
	if u.Age != 0 {
		query += " age = ?,"
		values = append(values, u.Age)
	}
	query = query[:len(query)-1]
	values = append(values, u.Id)
	return query, values
}