package grades

import (
	"fmt"
	"sync"
)

type Student struct {
	ID        int
	FirstName string
	LastName  string
	Grades    []Grade
}

func (s Student) Average() float32 {
	var result float32
	for _, grade := range s.Grades {
		result += grade.Score
	}
	return result / float32(len(s.Grades))
}

type Students []Student

func (s Students) GetById(id int) (*Student, error) {
	for i := range s {
		if students[i].ID == id {
			return &students[i], nil
		}
	}
	return nil, fmt.Errorf("Student with id %d not found", id)
}

var (
	students     Students
	studentMutex sync.Mutex
)

type Grade struct {
	Title string
	Type  GradeType
	Score float32
}

type GradeType string

const (
	GradeTest     = GradeType("Test")
	GradeHomework = GradeType("Homework")
	GradeQuiz     = GradeType("Quiz")
)
