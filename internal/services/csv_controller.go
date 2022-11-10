package services

import (
	"bufio"
	"log"
	"os"
	"strings"

	"github.com/GabrielEdwinSP/go-csv/internal/domain"
)

const (
	FILE_NAME     = "internal/csv/students.csv"
	MAX_GOROUTINE = 10
)

func ProcessFile() {
	f, err := os.Open(FILE_NAME)
	if err != nil {
		log.Fatal(err)
	}

	students := ScanFile(f)

	ConcurrentProcessing(students)
}

func ScanFile(f *os.File) []*domain.Students {
	s := bufio.NewScanner(f)
	students := []*domain.Students{}
	for s.Scan() {
		line := strings.Trim(s.Text(), " ")
		linearray := strings.Split(line, ",")
		ids := strings.Split(linearray[5], " ")
		ids = ids[1 : len(ids)-1]
		student := &domain.Students{
			Id:          linearray[0],
			Name:        linearray[1],
			LastName:    linearray[2],
			Email:       linearray[3],
			Phone:       linearray[4],
			StudentsIds: ids,
		}
		students = append(students, student)
	}
	return students
}
