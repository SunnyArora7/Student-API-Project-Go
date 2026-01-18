package storage

import (
	typesFile "studentPackage/internal/type"
)

type Storage interface {
	CreateStudent(name string, email string, age int) (int64, error)
	GetStudent(id int) (typesFile.Student, error)
	GetStudents() ([]typesFile.Student, error)
}
