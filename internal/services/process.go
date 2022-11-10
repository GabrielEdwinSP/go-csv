package services

import (
	"fmt"
	"time"

	"github.com/GabrielEdwinSP/go-csv/internal/domain"
)

func SendSmsNotification(student *domain.Students) {
	time.Sleep(10 * time.Millisecond)
	fmt.Printf("Sending sms notification to %v\n", student.Phone)
}

func FindUserById(studentId string, students []*domain.Students) (*domain.Students, error) {
	for _, student := range students {
		if student.Id == studentId {
			return student, nil
		}
	}

	return nil, fmt.Errorf("User not found with id %v", studentId)
}
