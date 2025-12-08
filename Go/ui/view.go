package ui

import (
	"fmt"
	"myapp/dto"
)

func GetMenuOption() int {

	var userInput int

	fmt.Println("Main Menu")
	fmt.Println("1. List Records")
	fmt.Println("2. Create New Record")
	fmt.Println("3. View a Record")
	fmt.Println("4. Remove a Record")
	fmt.Println("5. Exit")

	fmt.Scanln(&userInput)

	return userInput
}

func GetListOption() int {

	var userInput int

	fmt.Println("List Options")
	fmt.Println("1. Students")
	fmt.Println("2. Teachers")
	fmt.Println("3. Courses")

	fmt.Scanln(&userInput)

	return userInput

}

func GetStudentInfo() (string, string) {

	var fName string
	var lName string

	fmt.Println("Students first name: ")
	fmt.Scanln(&fName)

	fmt.Println("Students last name: ")
	fmt.Scanln(&lName)

	return fName, lName
}

func GetTeacherInfo() (string, string, string, string, string) {

	var fName string
	var lName string
	var title string
	var office string
	var department string

	fmt.Println("Teachers first name: ")
	fmt.Scanln(&fName)

	fmt.Println("Teachers last name: ")
	fmt.Scanln(&lName)

	fmt.Println("Teachers title: ")
	fmt.Scanln(&title)

	fmt.Println("Teachers office: ")
	fmt.Scanln(&office)

	fmt.Println("Teachers department: ")
	fmt.Scanln(&department)

	return fName, lName, title, office, department
}

func GetCourseInfo() string {

	var name string

	fmt.Println("Course name: ")
	fmt.Scanln(&name)

	return name
}

func PrintStudent(student dto.Student) {
	fmt.Printf("Student:\n  ID: %d\n  First Name: %s\n  Last Name: %s\n\n",
		student.ID, student.FirstName, student.LastName)
}

func PrintTeacher(teacher dto.Teacher) {
	fmt.Printf("Teacher:\n  ID: %d\n  First Name: %s\n  Last Name: %s\n  Title: %s\n  Office: %s\n  Department: %s\n\n",
		teacher.ID, teacher.FirstName, teacher.LastName, teacher.Title, teacher.Office, teacher.Department)
}

func PrintCourse(course dto.Course) {
	fmt.Printf("Course:\n  ID: %d\n  Name: %s\n\n",
		course.ID, course.Name)
}

func PrintInvalidInput() {
	fmt.Println("Invalid input")
}
