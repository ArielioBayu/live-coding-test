package repository

import (
	"a21hc3NpZ25tZW50/model"
)

type StudentRepository interface {
	FetchAll() ([]model.Student, error)
	Store(student *model.Student) error
	ResetStudentRepo()
}

type studentRepository struct {
	students []model.Student
}

func NewStudentRepo() *studentRepository {
	return &studentRepository{}
}

func (s *studentRepository) FetchAll() ([]model.Student, error) {
	return s.students, nil
}

func (s *studentRepository) Store(student *model.Student) error {
	s.students = append(s.students, *student)
	return nil
}

func (s *studentRepository) ResetStudentRepo() {
	s.students = []model.Student{}
}
