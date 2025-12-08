package dto

type Student struct {
	ID        int
	FirstName string
	LastName  string
}

type Teacher struct {
	ID         int
	FirstName  string
	LastName   string
	Title      string
	Office     string
	Department string
}

type Course struct {
	ID   int
	Name string
}

type StudentCourse struct {
	Student  Student
	Course   Course
	Semester string
	Grade    string
}

type TeacherCourse struct {
	Teacher  Teacher
	Course   Course
	Semester string
}
