package controller

import (
	"database/sql"
	"fmt"
	"myapp/dao"
	"myapp/dto"
	"myapp/ui"

	_ "github.com/go-sql-driver/mysql"
)

func Run(db *sql.DB) {

	for {

		menuOption := ui.GetMenuOption()

		switch menuOption {
		case 1:
			listRecords(db)
		case 2:
			createRecords(db)
		case 3:
			viewRecords(db)
		case 4:
			removeRecords(db)
		case 5:
			return
		default:
			ui.PrintInvalidInput()
		}

	}

}

func listRecords(db *sql.DB) {

	option := ui.GetListOption()

	switch option {
	case 1:
		students, err := dao.GetStudentList(db)

		if err != nil {
			fmt.Println(err)
			return
		}

		for _, student := range students {
			ui.PrintStudent(student)
		}

	case 2:
		teachers, err := dao.GetTeacherList(db)

		if err != nil {
			fmt.Println(err)
			return
		}

		for _, teacher := range teachers {
			ui.PrintTeacher(teacher)
		}
	case 3:
		courses, err := dao.GetCourseList(db)

		if err != nil {
			fmt.Println(err)
			return
		}

		for _, course := range courses {
			ui.PrintCourse(course)
		}
	default:
		ui.PrintInvalidInput()
	}

	// fmt.Println("List")
}

func createRecords(db *sql.DB) {
	option := ui.GetListOption()

	switch option {
	case 1:

		student := &dto.Student{}

		student.FirstName, student.LastName = ui.GetStudentInfo()

		if err := dao.CreateStudentRecord(db, student); err != nil {
			fmt.Println(err)
			return
		}

	case 2:
		teacher := &dto.Teacher{}

		teacher.FirstName, teacher.LastName, teacher.Title, teacher.Office, teacher.Department = ui.GetTeacherInfo()

		if err := dao.CreateTeacherRecord(db, teacher); err != nil {
			fmt.Println(err)
			return
		}

	case 3:
		course := &dto.Course{}

		course.Name = ui.GetCourseInfo()

		if err := dao.CreateCourseRecord(db, course); err != nil {
			fmt.Println(err)
			return
		}

	default:
		ui.PrintInvalidInput()
	}
}

func viewRecords(db *sql.DB) {
	fmt.Println("View")
}

func removeRecords(db *sql.DB) {
	option := ui.GetListOption()

	switch option {
	case 1:

		id := ui.GetRemoveID()

		if err := dao.DeleteStudentRecord(db, id); err != nil {
			fmt.Println(err)
			return
		}

	case 2:

		id := ui.GetRemoveID()

		if err := dao.DeleteTeacherRecord(db, id); err != nil {
			fmt.Println(err)
			return
		}

	case 3:

		id := ui.GetRemoveID()

		if err := dao.DeleteCourseRecord(db, id); err != nil {
			fmt.Println(err)
			return
		}

	default:
		ui.PrintInvalidInput()
	}
}
