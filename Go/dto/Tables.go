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
	StudentID int
	CourseID  int
	Semester  string
	Grade     string
}

type TeacherCourse struct {
	TeacherID int
	CourseID  int
	Semester  string
}
