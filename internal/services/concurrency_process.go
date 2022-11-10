package services

import (
	"fmt"

	"github.com/GabrielEdwinSP/go-csv/internal/domain"
)

func ConcurrentProcessing(students []*domain.Students) {
	studentsCh := make(chan []*domain.Students)
	unvisitedStudents := make(chan *domain.Students)
	go func() {
		studentsCh <- students
	}()
	InitizializeConcurrency(unvisitedStudents, studentsCh, students)
	ProcessStudents(unvisitedStudents, studentsCh, len(students))
}

func InitizializeConcurrency(unvisitedStudents <-chan *domain.Students, studentsCh chan []*domain.Students, students []*domain.Students) {
	for i := 0; i < MAX_GOROUTINE; i++ {
		go func() {
			for student := range unvisitedStudents {
				SendSmsNotification(student)
				go func(student *domain.Students) {
					studentIds := student.StudentsIds
					students := []*domain.Students{}
					for _, studentId := range studentIds {
						student, err := FindUserById(studentId, students)
						if err != nil {
							fmt.Printf("Error %v\n", err)
							continue
						}
						students = append(students, student)
					}
					_, ok := <-studentsCh
					if ok {
						studentsCh <- students
					}
				}(student)
			}
		}()
	}
}

func ProcessStudents(unvisitedStudents chan<- *domain.Students, studentsCh chan []*domain.Students, size int) {
	visitedStudents := make(map[string]bool)
	count := 0
	for students := range studentsCh {
		for _, student := range students {
			if !visitedStudents[student.Id] {
				visitedStudents[student.Id] = true
				count++
				if count >= size {
					close(studentsCh)
				}
				unvisitedStudents <- student
			}
		}
	}
}
