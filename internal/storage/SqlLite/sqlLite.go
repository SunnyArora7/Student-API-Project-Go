package sqllite

import (
	"database/sql"
	"fmt"
	"studentPackage/internal/config"
	typesFile "studentPackage/internal/type"

	_ "github.com/mattn/go-sqlite3"
)

type Sqllite struct {
	Db *sql.DB
}

func New(cf *config.Config) (*Sqllite, error) {
	db, err := sql.Open("sqlite3", cf.StoragePath)
	if err != nil {
		return nil, err
	}
	_, err = db.Exec(`CREATE TABLE IF NOT EXISTS student (
	id INTEGER PRIMARY KEY AUTOINCREMENT,
	name TEXT,
	email TEXT,
	age INTEGER
)
 `)
	if err != nil {
		return nil, err
	}

	return &Sqllite{
		Db: db,
	}, nil

}

func (s *Sqllite) CreateStudent(name string, email string, age int) (int64, error) {
	stm, err := s.Db.Prepare("INSERT INTO student (name, email, age) VALUES (?, ?, ?)")
	if err != nil {
		return 0, err
	}
	defer stm.Close()
	res, err := stm.Exec(name, email, age)
	if err != nil {
		return 0, err
	}
	lastInsertId, err := res.LastInsertId()
	if err != nil {
		return 0, err
	}
	return lastInsertId, nil

}

func (s *Sqllite) GetStudent(id int) (typesFile.Student, error) {
	stmt, err := s.Db.Prepare("SELECT id, name, email, age FROM student WHERE id = ? LIMIT 1")
	if err != nil {
		return typesFile.Student{}, err
	}
	defer stmt.Close()
	var student typesFile.Student
	err = stmt.QueryRow(id).Scan(&student.Id, &student.Name, &student.Email, &student.Age)
	if err != nil {
		if err == sql.ErrNoRows {
			return typesFile.Student{}, fmt.Errorf("Student not found with id %d", id)
		}
		return typesFile.Student{}, err
	}

	return student, nil
}

func (s *Sqllite) GetStudents() ([]typesFile.Student, error) {
	stmt, err := s.Db.Prepare("SELECT id, name, email, age FROM student")
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	rows, err := stmt.Query()
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var students []typesFile.Student
	for rows.Next() {
		var student typesFile.Student
		if err := rows.Scan(&student.Id, &student.Name, &student.Email, &student.Age); err != nil {
			return nil, err
		}
		students = append(students, student)
	}

	return students, nil
}
