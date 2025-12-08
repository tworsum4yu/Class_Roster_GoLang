package service

import (
	"myapp/dto"
)

func UpdateStudentRecordFix(old, updated *dto.Student) *dto.Student {
	return &dto.Student{
		ID:        old.ID,
		FirstName: pick(old.FirstName, updated.FirstName),
		LastName:  pick(old.LastName, updated.LastName),
	}
}

func UpdateTeacherRecordFix(old, updated *dto.Teacher) *dto.Teacher {
	return &dto.Teacher{
		ID:         old.ID,
		FirstName:  pick(old.FirstName, updated.FirstName),
		LastName:   pick(old.LastName, updated.LastName),
		Title:      pick(old.Title, updated.Title),
		Office:     pick(old.Office, updated.Office),
		Department: pick(old.Department, updated.Department),
	}
}

func UpdatCourseRecordFix(old, updated *dto.Course) *dto.Course {
	return &dto.Course{
		ID:   old.ID,
		Name: pick(old.Name, updated.Name),
	}
}

func pick(oldVal, newVal string) string {
	if newVal == "" {
		return oldVal
	}
	return newVal
}
